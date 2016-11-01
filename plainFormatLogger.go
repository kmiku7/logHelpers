package logHelpers

import (
	"io"
	"log"
)

type plainFormatLogger struct {
	logger *log.Logger
}

func NewPlainFormatLogger(output io.Writer) Logger {
	return &plainFormatLogger{
		logger: log.New(output, "", log.LstdFlags|log.Lshortfile),
	}
}

func (s *plainFormatLogger) Debug(args ...interface{}) {
	s.logger.Print(args)
}

func (s *plainFormatLogger) Debugf(format string, args ...interface{}) {
	s.logger.Printf(format, args)
}

func (s *plainFormatLogger) Info(args ...interface{}) {
	s.logger.Print(args)
}

func (s *plainFormatLogger) Infof(format string, args ...interface{}) {
	s.logger.Printf(format, args)
}

func (s *plainFormatLogger) Warning(args ...interface{}) {
	s.logger.Print(args)
}

func (s *plainFormatLogger) Warningf(format string, args ...interface{}) {
	s.logger.Printf(format, args)
}

func (s *plainFormatLogger) Error(args ...interface{}) {
	s.logger.Print(args)
}

func (s *plainFormatLogger) Errorf(format string, args ...interface{}) {
	s.logger.Printf(format, args)
}
func (s *plainFormatLogger) Fatal(args ...interface{}) {
	s.logger.Print(args)
}

func (s *plainFormatLogger) Fatalf(format string, args ...interface{}) {
	s.logger.Printf(format, args)
}
