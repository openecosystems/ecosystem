package tasks

import (
	"fmt"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// TaskFinishedMsg represents a message indicating the completion of a task, including task details and an optional error.
type TaskFinishedMsg struct {
	Task        Task
	SectionID   int
	SectionType string
	State       State
}

// ClearTaskMsg represents a message used to clear or remove a specific task identified by its TaskID.
type ClearTaskMsg struct {
	TaskID string
}

// State is an alias for the int type, commonly used to represent various states or stages in a process.
type State = int

// TaskStart represents the initial state of a task.
// TaskFinished represents the state when a task is successfully completed.
// TaskError represents the state when a task encounters an error.
const (
	TaskStart State = iota
	TaskFinished
	TaskError
)

const (
	workerCount = 3
)

// TaskMap is a concurrent map used to store and manage tasks by their unique IDs.
// TaskQueue is a slice that maintains the order of task IDs to ensure FIFO processing of tasks.
// QueueLock is a mutex used to ensure thread-safe access to the task queue for maintaining FIFO order of tasks.
var (
	CompletedTaskCmdsChan = make(chan tea.Msg)
	TaskMap               sync.Map

	wg               sync.WaitGroup
	once             sync.Once
	fifoTaskChan     = make(chan string, 10)
	parallelTaskChan = make(chan string, 10)
)

// Task represents a unit of work with its associated metadata and lifecycle states.
type Task struct {
	ID           string
	Parallel     bool
	StartText    string
	FinishedText string
	State        State
	Error        error
	StartTime    time.Time
	FinishedTime *time.Time
	Msg          func(pctx *context.ProgramContext, err error) tea.Msg
	Done         bool
}

// CmdTask extends Task, adding a tea.Cmd value to represent additional commands within the task's lifecycle.
type CmdTask struct {
	Task
	Cmd tea.Cmd
}

// AddTask adds a new task to the program context and updates the task map and task queue in a FIFO manner.
// If the task has an empty ID, it will not be added. Returns the updated ProgramContext.
func AddTask(task Task) *sync.Map {
	if task.ID == "" {
		log.Debug("Silently dropping tasks as ID is missing")
		return &TaskMap
	}
	TaskMap.Store(task.ID, task)

	wg.Add(1)
	if task.Parallel {
		parallelTaskChan <- task.ID
		log.Debugf("Parallel Task added: %s", task.ID)
	} else {
		fifoTaskChan <- task.ID
		log.Debugf("FIFO Task added: %s", task.ID)
	}

	return &TaskMap
}

// fifoWorker Worker for FIFO tasks (single worker, strict order)
func fifoWorker() {
	for taskID := range fifoTaskChan {
		if task, ok := TaskMap.Load(taskID); ok {
			t := task.(Task)

			// RUN INTERFACE METHOD HERE OR FUNCTION
			// fmt.Println("Processing FIFO:", taskID)
			time.Sleep(1 * time.Second)

			t.Done = true
			TaskMap.Store(taskID, t)
			now := time.Now()
			t.FinishedTime = &now
			log.Debug("Completed FIFO:", taskID)

			TaskMap.Delete(taskID)

			CompletedTaskCmdsChan <- TaskFinishedMsg{
				Task:  t,
				State: TaskFinished,
			}
		}
		wg.Done()
	}
}

func parallelWorker() {
	for taskID := range parallelTaskChan {
		if task, ok := TaskMap.Load(taskID); ok {
			t := task.(Task)

			// Simulate processing
			// fmt.Println("Processing Parallel:", taskID)
			time.Sleep(1 * time.Second)

			// if Error return State Error

			t.Done = true
			TaskMap.Store(taskID, t)
			now := time.Now()
			t.FinishedTime = &now
			log.Debug("Completed Parallel:", taskID)

			TaskMap.Delete(taskID)

			CompletedTaskCmdsChan <- TaskFinishedMsg{
				Task:  t,
				State: TaskFinished,
			}
		}
		wg.Done()
	}
}

// ProcessTasks Runs in the background
func ProcessTasks() {
	once.Do(func() {
		go fifoWorker()

		for i := 0; i < workerCount; i++ {
			go parallelWorker()
		}
	})
}

// Close gracefully shutdown all channels by waiting for in flight tasks
func Close() {
	fmt.Println("Waiting for tasks to complete")
	wg.Wait()

	close(fifoTaskChan)
	close(parallelTaskChan)
}
