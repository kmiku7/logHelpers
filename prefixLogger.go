package logHelpers

type prefixLogger struct {
	logger Logger
	prefix string
}

func NewPrefixLogger(prefix string, logger Logger) Logger {
	return &prefixLogger{logger, prefix}
}

func (s *prefixLogger) Debug(args ...interface{}) {
	s.logger.Debug(s.prefix, args...)
}

func (s *prefixLogger) Debugf(format string, args ...interface{}) {
	s.logger.Debugf(s.prefix+" "+format, args...)
}

func (s *prefixLogger) Info(args ...interface{}) {
	s.logger.Debug(s.prefix, args...)
}

func (s *prefixLogger) Infof(format string, args ...interface{}) {
	s.logger.Debugf(s.prefix+" "+format, args...)
}

func (s *prefixLogger) Warning(args ...interface{}) {
	s.logger.Debug(s.prefix, args...)
}

func (s *prefixLogger) Warningf(format string, args ...interface{}) {
	s.logger.Debugf(s.prefix+" "+format, args...)
}

func (s *prefixLogger) Error(args ...interface{}) {
	s.logger.Debug(s.prefix, args...)
}

func (s *prefixLogger) Errorf(format string, args ...interface{}) {
	s.logger.Debugf(s.prefix+" "+format, args...)
}
func (s *prefixLogger) Fatal(args ...interface{}) {
	s.logger.Debug(s.prefix, args...)
}

func (s *prefixLogger) Fatalf(format string, args ...interface{}) {
	s.logger.Debugf(s.prefix+" "+format, args...)
}
