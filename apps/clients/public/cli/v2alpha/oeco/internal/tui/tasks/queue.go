package tasks

import (
	"fmt"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// TaskExecutor defines an interface for executing tasks within a given ProgramContext and returning a message of type tea.Msg.
type TaskExecutor interface {
	Execute(pctx *context.ProgramContext, err error) (tea.Msg, error)
}

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
	Task   Task
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
	pwg              sync.WaitGroup
	once             sync.Once
	fifoTaskChan     = make(chan string, 10)
	parallelTaskChan = make(chan string, 10)
)

// Task represents a unit of work with its associated metadata and lifecycle states.
type Task struct {
	Ctx          *context.ProgramContext
	ID           string
	Parallel     bool
	StartText    string
	FinishedText string
	State        State
	Error        error
	StartTime    time.Time
	FinishedTime *time.Time
	Msg          tea.Msg
	TaskExecutor TaskExecutor
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
	if task.Ctx == nil {
		panic("Please add a Ctx prior to adding a task")
	}

	if task.TaskExecutor == nil {
		panic("Please add a Task Executor prior to adding a task")
	}

	if task.ID == "" {
		log.Debug("Silently dropping tasks as ID is missing")
		return &TaskMap
	}
	TaskMap.Store(task.ID, task)

	if task.Parallel {
		pwg.Add(1)
		parallelTaskChan <- task.ID
		log.Debugf("Parallel Task added: %s", task.ID)
	} else {
		wg.Add(1)
		fifoTaskChan <- task.ID
		log.Debugf("FIFO Task added: %s", task.ID)
	}

	return &TaskMap
}

// fifoWorker Worker for FIFO tasks (single worker, strict order)
func fifoWorker() {
	for taskID := range fifoTaskChan {
		if task, ok := TaskMap.Load(taskID); ok {
			handleWork(taskID, task)
		}
		wg.Done()
	}
}

func parallelWorker() {
	for taskID := range parallelTaskChan {
		if task, ok := TaskMap.Load(taskID); ok {
			handleWork(taskID, task)
		}
		pwg.Done()
	}
}

func handleWork(taskID string, task any) {
	t := task.(Task)

	TaskMap.Store(taskID, t)
	msg, err := t.TaskExecutor.Execute(t.Ctx, t.Error)

	now := time.Now()
	t.Msg = msg
	t.Done = true

	t.FinishedTime = &now
	TaskMap.Delete(taskID)

	if err != nil {
		CompletedTaskCmdsChan <- TaskFinishedMsg{
			Task:  t,
			State: TaskError,
		}
	} else {
		CompletedTaskCmdsChan <- TaskFinishedMsg{
			Task:  t,
			State: TaskFinished,
		}
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
	pwg.Wait()

	close(fifoTaskChan)
	close(parallelTaskChan)
}
