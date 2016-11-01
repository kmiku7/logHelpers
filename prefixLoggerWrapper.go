package logHelpers

type prefixLoggerWrapper struct {
	logger Logger
	prefix string
}

func NewPrefixLoggerWrapper(prefix string, logger Logger) Logger {
	return &prefixLoggerWrapper{logger, prefix}
}

func (s *prefixLoggerWrapper) Debug(args ...interface{}) {
	s.logger.Debug(append([]interface{}{s.prefix}, args...)...)
}

func (s *prefixLoggerWrapper) Debugf(format string, args ...interface{}) {
	s.logger.Debugf(s.prefix+" "+format, args...)
}

func (s *prefixLoggerWrapper) Info(args ...interface{}) {
	s.logger.Info(append([]interface{}{s.prefix}, args...)...)
}

func (s *prefixLoggerWrapper) Infof(format string, args ...interface{}) {
	s.logger.Infof(s.prefix+" "+format, args...)
}

func (s *prefixLoggerWrapper) Warning(args ...interface{}) {
	s.logger.Warning(append([]interface{}{s.prefix}, args...)...)
}

func (s *prefixLoggerWrapper) Warningf(format string, args ...interface{}) {
	s.logger.Warningf(s.prefix+" "+format, args...)
}

func (s *prefixLoggerWrapper) Error(args ...interface{}) {
	s.logger.Error(append([]interface{}{s.prefix}, args...)...)
}

func (s *prefixLoggerWrapper) Errorf(format string, args ...interface{}) {
	s.logger.Errorf(s.prefix+" "+format, args...)
}
func (s *prefixLoggerWrapper) Fatal(args ...interface{}) {
	s.logger.Fatal(append([]interface{}{s.prefix}, args...)...)
}

func (s *prefixLoggerWrapper) Fatalf(format string, args ...interface{}) {
	s.logger.Fatalf(s.prefix+" "+format, args...)
}
