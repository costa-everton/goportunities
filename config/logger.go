package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	err *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger{
	writer := io.Writer(os.Stdout)
	logger := log.New(writer,prefix, log.Ldate|log.Ltime )

	return &Logger{
		debug: log.New(writer, "DEBUG: ", logger.Flags()),
		info: log.New(writer, "INFO: ", log.Flags()),
		warning: log.New(writer, "WARNING: ", log.Flags()),
		err: log.New(writer, "ERROR: ", log.Flags()),
		writer: writer,
	}
}

// Create Non-Formatted Logs
func (c *Logger) Debug(v ...interface{}){
	c.debug.Println(v ...)
}
func (c *Logger) Info(v ...interface{}){
	c.info.Println(v ...)
}
func (c *Logger) Warning(v ...interface{}){
	c.warning.Println(v ...)
}
func (c *Logger) Error(v ...interface{}){
	c.err.Println(v ...)
}

// Create Format Enable Logs
func (c *Logger) Debugf(format string, v ...interface{}){
	c.debug.Printf(format, v ...)
}
func (c *Logger) Infof(format string, v ...interface{}){
	c.info.Printf(format, v ...)
}
func (c *Logger) Warningf(format string, v ...interface{}){
	c.warning.Printf(format, v ...)
}
func (c *Logger) Errorf(format string, v ...interface{}){
	c.err.Printf(format, v ...)
}