// Package trace is used for printf-style tracing of go
// code. Embedding Print*() calls in go code will print out a leader,
// the current source file name, the current source line number, and
// an optional message to the configured Writer (defaults to
// os.Stdout) when the line is run.
//
// There is also a Dump() function used to dump complex data
// structures in an easy to understand format. The Dump function uses
// a modified version of Dave Collins'
// https://github.com/davecgh/go-spew package to pretty-print data
// structures.
package trace

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/apatters/go-trace/spew"
)

const (
	stdoutName = "/dev/stdout"
)

var (
	// SpewCS holds the spew configurion used by the Dump
	// function. See
	// https://github.com/davecgh/go-spew#configuration-options
	// for details.
	SpewCS *spew.ConfigState

	// Leader is printed at the beginning of every trace line.
	Leader = "### "

	// Writer is used for trace output.
	Writer io.Writer = os.Stdout

	// TraceLevel is used to control output of Print*Level functions.
	TraceLevel int
)

// init inializes the spew configuration.
func init() {
	SpewCS = spew.NewDefaultConfig()
	SpewCS.Indent = "\t"
	SpewCS.SortKeys = true
	SpewCS.DisableMethods = true
	SpewCS.SortKeys = true
	SpewCS.SpewKeys = true
}

// fprint wraps output to the io.Writer. The go test command requires output go directly
// to os.Stdout.
func fprint(w io.Writer, leader string, a ...interface{}) (n int, err error) {
	var msg string

	if len(a) == 0 {
		msg = strings.TrimRight(leader, " \t\n")
	} else {
		msg = leader + fmt.Sprint(a...)
	}
	msg = strings.TrimRight(msg, " \t\n")

	switch v := w.(type) {
	case *os.File:
		if v.Name() == stdoutName {
			return fmt.Println(msg)
		}
	}

	return fmt.Fprintln(Writer, msg)
}

// fprintln wraps output to the io.Writer. The go test command
// requires output go directly to os.Stdout.
func fprintln(w io.Writer, leader string, a ...interface{}) (n int, err error) {
	var msg string

	if len(a) == 0 {
		msg = strings.TrimRight(leader, " \t\n")
	} else {
		msg = leader + fmt.Sprintln(a...)
	}
	msg = strings.TrimRight(msg, " \t\n")

	switch v := w.(type) {
	case *os.File:
		if v.Name() == stdoutName {
			return fmt.Println(msg)
		}
	default:
	}

	return fmt.Fprintln(Writer, msg)
}

// fprintf wraps output to the io.Writer. The go test command requires
// output go directly to os.Stdout.
func fprintf(w io.Writer, leader string, format string, a ...interface{}) (n int, err error) {
	var msg string

	if len(a) == 0 {
		msg = strings.TrimRight(leader, " \t\n")
	} else {
		msg = leader + fmt.Sprintf(format, a...)
	}
	msg = strings.TrimRight(msg, " \t\n")

	switch v := w.(type) {
	case *os.File:
		if v.Name() == stdoutName {
			return fmt.Println(msg)
		}
	}

	return fmt.Fprintln(Writer, msg)
}

func leader(filename string, line int) string {
	return fmt.Sprintf("%s%s:%d ", Leader, path.Base(filename), line)
}

// Print outputs the leader, source file name, and source line number
// followed by any args in a similar manner as fmt.Print.
func Print(args ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	fprint(Writer, leader(filename, line), args...)
}

// Print outputs the leader, source file name, and source line number
// followed by any args in a similar manner as fmt.Println.
func Println(args ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	_, _ = fprintln(Writer, leader(filename, line), args...)
}

// Print outputs the leader, source file name, and source line number
// followed by any args in a similar manner as fmt.Printf. A trailing
// newline is also output.
func Printf(format string, args ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	_, _ = fprintf(Writer, leader(filename, line), format, args...)

}

// PrintLevel operates identically to Print except no output is done
// if level is greater that the current trace level (TraceLevel).
func PrintLevel(level int, args ...interface{}) {
	if level > TraceLevel {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	_, _ = fprint(Writer, leader(filename, line), args...)

}

// PrintlnLevel operates identically to Println except no output is
// done if level is greater that the current trace level (TraceLevel).
func PrintlnLevel(level int, args ...interface{}) {
	if level > TraceLevel {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	_, _ = fprintln(Writer, leader(filename, line), args...)
}

// PrintfLevel operates identically to Printf except no output is done
// if level is greater that the current trace level (TraceLevel).
func PrintfLevel(level int, format string, args ...interface{}) {
	if level > TraceLevel {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	_, _ = fprintf(Writer, leader(filename, line), format, args...)
}

// Dump() outputs the leader, source file name, and source line number
// followed by pretty-printed versions of any args. It is meant to be
// used for more complex data structures. Dump output can be modified
// by changing values in the SpewCS global variable.  See
// https://github.com/davecgh/go-spew#configuration-options for
// details.
func Dump(args ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	_, _ = fprintln(Writer, leader(filename, line))
	SpewCS.Fdump(Writer, args...)
}
