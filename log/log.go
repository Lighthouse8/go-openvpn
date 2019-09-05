package log

import "log"

func Trace(a ...interface{})  {}
func Tracef(a ...interface{}) {}
func Debug(a ...interface{})  {}
func Debugf(a ...interface{}) {}
func Info(a ...interface{})   {}
func Infof(a ...interface{})  {}
func Warn(a ...interface{}) {
	log.Printf("[WARN] %v\n", a)
}
func Warnf(a ...interface{}) {
	log.Printf("[WARN] %v\n", a)
}
func Error(a ...interface{}) {
	log.Printf("[ ERR] %v\n", a)
}
func Errorf(a ...interface{}) {
	log.Printf("[ ERR] %v\n", a)
}
func Critical(a ...interface{}) {
	log.Printf("[CRIT] %v\n", a)
}
func Criticalf(a ...interface{}) {
	log.Printf("[CRIT] %v\n", a)
}
