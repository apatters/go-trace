package trace_test

import (
	"fmt"

	"github.com/apatters/go-trace"
)

type DumpTest struct {
	aString      string
	aStringSlice []string
	anInt        int
}

func Example() {
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
	trace.Leader = savedLeader

	// Output:
	// ### example_test.go:17
	// ### example_test.go:18 Print()second3
	// ### example_test.go:19 Println() second 3
	// ### example_test.go:20 Printf() second 3
	// TraceLevel = 1
	// ### example_test.go:26 Print at level 0
	// ### example_test.go:27 Print at level 1
	// TraceLevel set back to default = 0
	// ### example_test.go:31 Print at level 0
	// ### example_test.go:36
	// ([]string) (len=3 cap=3) {
	// 	(string) (len=3) "now",
	// 	(string) (len=2) "is",
	// 	(string) (len=8) "the time"
	// }
	// ### example_test.go:37
	// (trace_test.DumpTest) {
	// 	aString: (string) (len=15) "now is the time",
	// 	aStringSlice: ([]string) (len=3 cap=3) {
	// 		(string) (len=3) "now",
	// 		(string) (len=2) "is",
	// 		(string) (len=8) "the time"
	// 	},
	// 	anInt: (int) 1
	// }
	// ### example_test.go:41
	// (spew.ConfigState) {
	// 	Indent: (string) (len=1) "\t",
	// 	MaxDepth: (int) 0,
	// 	DisableMethods: (bool) true,
	// 	DisablePointerMethods: (bool) false,
	// 	DisablePointerAddresses: (bool) false,
	// 	DisableCapacities: (bool) false,
	// 	ContinueOnMethod: (bool) false,
	// 	SortKeys: (bool) true,
	// 	SpewKeys: (bool) true
	// }
	// 	* example_test.go:46
}
