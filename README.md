go-trace
=========

Golang printf-style, code tracing package

[![Build Status](https://travis-ci.org/apatters/go-trace.svg)](https://travis-ci.org/apatters/go-trace) [![GoDoc](https://godoc.org/github.com/apatters/go-trace?status.svg)](https://godoc.org/github.com/apatters/go-trace)

The trace is used for printf-style tracing of go code. Embedding
Print*() calls in go code will print out a leader, the current source
file name, the current source line number, and an optional message to
the configured Writer (defaults to os.Stdout) when the line is run.

There is also a Dump() function used to dump complex data structures
in an easy to understand format. The Dump function uses a modified
version of Dave Collins' https://github.com/davecgh/go-spew package to
pretty-print data structures

Documentation
-------------

Documentation can be found at [GoDoc](https://godoc.org/github.com/apatters/go-trace)


Installation
------------

```bash
$ go get -u github.com/apatters/go-trace
```

## Quick Start

Add this import line to the file you are working in"

``` go
import "github.com/apatters/go-trace"
```

To trace a line of code when the program is run, use one of the
Print*() or the Dump() functions, e.g.:

``` go
trace.Println("Print file name, line number, and this message when this line line is executed.")
```

Example
-------

``` go
package main

import (
	"fmt"

	"github.com/apatters/go-trace"
)

type DumpTest struct {
	aString      string
	aStringSlice []string
	anInt        int
}

func main() {
	// Standard fmt.Print*()-style functions.
	trace.Print()
	trace.Print("Print()", "second", 3)
	trace.Println("Println()", "second", 3)
	trace.Printf("Printf() %s %d", "second", 3)

	// Trace levels.
	savedTraceLevel := trace.TraceLevel
	trace.TraceLevel = 1
	fmt.Printf("TraceLevel = %d\n", trace.TraceLevel)
	trace.PrintLevel(0, "Print at level 0")
	trace.PrintLevel(1, "Print at level 1")
	trace.PrintLevel(2, "Print at level 2") // Not printed.
	trace.TraceLevel = savedTraceLevel
	fmt.Printf("TraceLevel set back to default = %d\n", trace.TraceLevel)
	trace.PrintLevel(0, "Print at level 0")
	trace.PrintLevel(1, "Print at level 1") // Not printed.
	trace.PrintLevel(2, "Print at level 2") // Not printed.

	// Pretty-print data structures.
	trace.Dump([]string{"now", "is", "the time"})
	trace.Dump(DumpTest{
		"now is the time",
		[]string{"now", "is", "the time"},
		1})
	trace.Dump(*trace.SpewCS)
	
	// Change the leader.
	savedLeader := trace.Leader
	trace.Leader = "\t* "
	trace.Print()
}
```

Which results in the following output (assuming your source file is
called 'example_test.go':

``` go
### example_test.go:17
### example_test.go:18 Print()second3
### example_test.go:19 Println() second 3
### example_test.go:20 Printf() second 3
TraceLevel = 1
### example_test.go:26 Print at level 0
### example_test.go:27 Print at level 1
TraceLevel set back to default = 0
### example_test.go:31 Print at level 0
### example_test.go:36
([]string) (len=3 cap=3) {
	(string) (len=3) "now",
	(string) (len=2) "is",
	(string) (len=8) "the time"
}
### example_test.go:37
(trace_test.DumpTest) {
	aString: (string) (len=15) "now is the time",
	aStringSlice: ([]string) (len=3 cap=3) {
		(string) (len=3) "now",
		(string) (len=2) "is",
		(string) (len=8) "the time"
	},
	anInt: (int) 1
}
### example_test.go:41
(spew.ConfigState) {
	Indent: (string) (len=1) "\t",
	MaxDepth: (int) 0,
	DisableMethods: (bool) true,
	DisablePointerMethods: (bool) false,
	DisablePointerAddresses: (bool) false,
	DisableCapacities: (bool) false,
	ContinueOnMethod: (bool) false,
	SortKeys: (bool) true,
	SpewKeys: (bool) true
}
	* example_test.go:46
```

License
-------

MIT.


Thanks
------

Thanks to [Secure64](https://secure64.com/company/) for
contributing this code.

