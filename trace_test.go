// The MIT License (MIT)
//
// Copyright (c) 2018, Secure64 Software Corporation
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package trace_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/apatters/go-trace"
	"github.com/stretchr/testify/assert"
)

type Dumper struct {
	Str  string
	Num  int
	Ptr  *int
	Strs []string
}

const (
	testDataStr = "hello, world"
	testDataNum = 1
)

var (
	dumpRegExpr = regexp.MustCompile(`(?s)^### trace_test.go:\d+.*` +
		`\(trace_test\.Dumper\) {.*` +
		`\sStr: \(string\) \(len=12\) "hello, world",.*` +
		`\sNum: \(int\) 1,.*` +
		`\sPtr: \(\*int\)\(0x[[:xdigit:]]+\)\(1\),.*` +
		`\sStrs: \(\[\]string\) \(len=3 cap=3\) {.*` +
		`\s\(string\) \(len=15\) "Now is the time",.*` +
		`\s\(string\) \(len=16\) "For all good men",.*` +
		`\s\(string\) \(len=36\) "To come to the aid of their country.".*` +
		`\s+}.*` +
		`\s+}.*`)
)

func TestPrint(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(
		`^### trace_test.go:[\d]+ %s%d\n$`,
		testDataStr,
		testDataNum))
	trace.Print(testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())
	trace.Writer = savedWriter
}

func TestPrintln(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(
		`^### trace_test.go:[\d]+ %s %d\n$`,
		testDataStr,
		testDataNum))
	trace.Println(testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())
	trace.Writer = savedWriter
}

func TestPrintf(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(`^### trace_test.go:[\d]+ %s %d\n$`, testDataStr, testDataNum))
	trace.Printf("%s %d", testDataStr, testDataNum)
	trace.Writer = savedWriter
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())
	trace.Writer = savedWriter
}

func TestPrintLevel(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(
		`^### trace_test.go:[\d]+ %s%d\n$`,
		testDataStr,
		testDataNum))
	savedTraceLevel := trace.TraceLevel
	trace.TraceLevel = 1

	trace.PrintLevel(0, testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())

	out.Reset()
	trace.PrintLevel(1, testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())

	out.Reset()
	trace.PrintLevel(2, testDataStr, testDataNum)
	t.Logf("out = %s", out)
	assert.Empty(t, out.String())
	trace.TraceLevel = savedTraceLevel
	trace.Writer = savedWriter
}

func TestPrintlnLevel(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(
		`^### trace_test.go:[\d]+ %s %d\n$`,
		testDataStr,
		testDataNum))
	savedTraceLevel := trace.TraceLevel
	trace.TraceLevel = 1

	trace.PrintlnLevel(0, testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())

	out.Reset()
	trace.PrintlnLevel(1, testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())

	out.Reset()
	trace.PrintlnLevel(2, testDataStr, testDataNum)
	t.Logf("out = %s", out)
	assert.Empty(t, out.String())
	trace.TraceLevel = savedTraceLevel
	trace.Writer = savedWriter
}

func TestPrintfLevel(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(
		`^### trace_test.go:[\d]+ %s %d\n$`,
		testDataStr,
		testDataNum))
	savedTraceLevel := trace.TraceLevel
	trace.TraceLevel = 1

	trace.PrintfLevel(0, "%s %d", testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())

	out.Reset()
	trace.PrintfLevel(1, "%s %d", testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())

	out.Reset()
	trace.PrintfLevel(2, "%s %d", testDataStr, testDataNum)
	t.Logf("out = %s", out)
	assert.Empty(t, out.String())
	trace.TraceLevel = savedTraceLevel
	trace.Writer = savedWriter
}

func TestDump(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	num := 1
	dumper := Dumper{
		Str: "hello, world",
		Num: 1,
		Ptr: &num,
		Strs: []string{
			"Now is the time",
			"For all good men",
			"To come to the aid of their country.",
		},
	}

	trace.Dump(dumper)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", dumpRegExpr)
	assert.Regexp(t, dumpRegExpr, out.String())
	trace.Writer = savedWriter
}

func TestLeader(t *testing.T) {
	var buf = make([]byte, 0, 256)
	out := bytes.NewBuffer(buf)
	savedWriter := trace.Writer
	trace.Writer = out

	cmpRegExpr := regexp.MustCompile(fmt.Sprintf(
		`^\*\*\* trace_test.go:[\d]+ %s%d\n$`,
		testDataStr,
		testDataNum))
	savedLeader := trace.Leader
	trace.Leader = "*** "
	trace.Print(testDataStr, testDataNum)
	t.Logf("out = %s", out)
	t.Logf("cmp = %s", cmpRegExpr)
	assert.Regexp(t, cmpRegExpr, out.String())
	trace.Leader = savedLeader
	trace.Writer = savedWriter
}
