package logHelpers

type discardLogger struct{}

func NewDiscardLogger() Logger {
	return nil
}

func (s *discardLogger) Debug(args ...interface{}) {
}

func (s *discardLogger) Debugf(format string, args ...interface{}) {
}

func (s *discardLogger) Info(args ...interface{}) {
}

func (s *discardLogger) Infof(format string, args ...interface{}) {
}

func (s *discardLogger) Warning(args ...interface{}) {
}

func (s *discardLogger) Warningf(format string, args ...interface{}) {
}

func (s *discardLogger) Error(args ...interface{}) {
}

func (s *discardLogger) Errorf(format string, args ...interface{}) {
}
func (s *discardLogger) Fatal(args ...interface{}) {
}

func (s *discardLogger) Fatalf(format string, args ...interface{}) {
}
