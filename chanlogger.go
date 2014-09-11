/*
 Package chanlogger is a tiny fmt clone, but all message will send to a channel befor write out, for in-order message output
*/

package chanlogger

import (
	"fmt"
	"io"
	"os"
)

type Clogger struct {
	in chan string // msg input channel
	w  io.Writer   // output File
}

// NewLogger new ChanLogger, default write message to os.Stderr
func NewLogger() *Clogger {
	l := &Clogger{
		in: make(chan string, 256),
		w:  os.Stderr,
	}
	go l.run()
	return l
}

// run start in goroutine, read msg from channel and write out
// goroutine exit if logger closed
func (l *Clogger) run() {
	var m string
	for m = range l.in {
		fmt.Fprint(l.w, m)
	}
	return
}

// Close set closed flag, close input channel, all new message will send to blackhole
// output goroutine exit if logger closed
func (l *Clogger) Close() {
	defer func() {
		// catch panic("close on closed channel")
		recover()
		return
	}()
	close(l.in)
	return
}

// SetWriter set logger output to io.Writer
func (l *Clogger) SetWriter(w io.Writer) {
	l.w = w
	return
}

// GetWriter return current io.Writer of logger
func (l *Clogger) GetWriter() io.Writer {
	return l.w
}

// Printf formats according to a format specifier and writes to standard output.
func (l *Clogger) Printf(format string, a ...interface{}) {
	defer func() {
		// catch panic("write on closed channel")
		recover()
		return
	}()
	l.in <- fmt.Sprintf(format, a...)
	return
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
func (l *Clogger) Print(a ...interface{}) {
	defer func() {
		// catch panic("write on closed channel")
		recover()
		return
	}()
	l.in <- fmt.Sprint(a...)
	return
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
func (l *Clogger) Println(a ...interface{}) {
	defer func() {
		// catch panic("write on closed channel")
		recover()
		return
	}()
	l.in <- fmt.Sprintln(a...)
	return
}

//
//
//
//
//
