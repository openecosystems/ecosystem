package zaploggerv1

import (
	apexlog "github.com/apex/log"
	"go.uber.org/zap"
)

type ZapHandler struct {
	z *zap.Logger
}

func (h *ZapHandler) HandleLog(e *apexlog.Entry) error {
	msg := e.Message
	fields := make([]zap.Field, 0, len(e.Fields))

	for k, v := range e.Fields {
		fields = append(fields, zap.Any(k, v))
	}

	switch e.Level {
	case apexlog.DebugLevel:
		h.z.Debug(msg, fields...)
	case apexlog.InfoLevel:
		h.z.Info(msg, fields...)
	case apexlog.WarnLevel:
		h.z.Warn(msg, fields...)
	case apexlog.ErrorLevel:
		h.z.Error(msg, fields...)
	case apexlog.FatalLevel:
		h.z.Fatal(msg, fields...)
	default:
		h.z.Info(msg, fields...)
	}

	return nil
}
