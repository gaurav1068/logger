package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Logger struct {
	prefix string
	target map[string]bool
	buff   *bufio.Writer
}

func (l *Logger) SetPrefix(str string) {
	l.prefix = str
}

func (l *Logger) SetTarget(target string) {
	l.target[target] = true
}
func (l *Logger) RemoveTarget(target string) {
	l.target[target] = false
}

func (l *Logger) INFO(str string) {
	txt := fmt.Sprintf("INFO " + l.prefix + " " + str + "\n")
	l.write(txt)
}
func (l *Logger) WARNING(str string) {
	txt := fmt.Sprintf("WARNING " + l.prefix + " " + str + "\n")
	l.write(txt)
}
func (l *Logger) TRACE(str string) {
	txt := fmt.Sprintf("TRACE " + l.prefix + " " + str + "\n")
	l.write(txt)
}
func (l *Logger) ERROR(str string) {
	txt := fmt.Sprintf("ERROR " + l.prefix + " " + str + "\n")
	l.write(txt)
}
func (l *Logger) FATAL(str string) {
	txt := fmt.Sprintf("FATAL " + l.prefix + " " + str + "\n")
	l.write(txt)
	os.Exit(1)
}
func (l *Logger) write(txt string) {
	val, _ := l.target["stdin"]
	if val {
		fmt.Print(txt)
	}
	val, _ = l.target["file"]
	if val {
		fmt.Fprintf(l.buff, txt)
		l.buff.Flush()
	}
}
func (l *Logger) SetFile(fileName string) error {
	return l.openFile(fileName)

}

func (l *Logger) openFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	l.buff = bufio.NewWriter(f)
	return nil
}

func NewLogger() *Logger {
	l := Logger{}
	t := time.Now()
	l.SetPrefix(t.String())
	l.openFile("logger")
	l.target = make(map[string]bool)
	l.target["stdin"] = true
	l.target["file"] = false
	return &l
}
