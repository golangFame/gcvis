//go:build !go1.20
// +build !go1.20

package parser

import (
	"bytes"
	"github.com/golangFame/gcvis/pkg/trace"
	"reflect"
	"testing"
	"time"
)

var parser *Parser

func runParserWith(line string) *Parser {
	reader := bytes.NewReader([]byte(line))
	parser = NewParser(reader)
	go parser.Run()
	return parser
}

func TestParserWithMatchingInput(t *testing.T) {
	line := "gc 763 @77536.239s 1%: 0.11+2192+0.75 ms clock, 0.92+9269/4379/3243+6.0 ms cpu, 6370->6390->3298 MB, 6533 MB goal, 8 P"

	runParserWith(line)

	expectedGCTrace := &trace.Gctrace{
		Heap1:        6533,
		ElapsedTime:  77536.239,
		STWSclock:    0.11,
		MASclock:     2192,
		STWMclock:    0.75,
		STWScpu:      0.92,
		MASAssistcpu: 9269,
		MASBGcpu:     4379,
		MASIdlecpu:   3243,
		STWMcpu:      6.0,
	}

	select {
	case gctrace := <-parser.GcChan:
		if !reflect.DeepEqual(gctrace, expectedGCTrace) {
			t.Errorf("Expected gctrace to equal %+v. Got %+v instead.", expectedGCTrace, gctrace)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}

func TestParserWithScvgLine(t *testing.T) {
	line := "scvg1: inuse: 12, idle: 13, sys: 14, released: 15, consumed: 16 (MB)"

	runParserWith(line)

	expectedScvgTrace := &trace.Scvgtrace{
		Inuse:    12,
		Idle:     13,
		Sys:      14,
		Released: 15,
		Consumed: 16,
	}

	select {
	case scvgTrace := <-parser.ScvgChan:
		if !reflect.DeepEqual(scvgTrace, expectedScvgTrace) {
			t.Errorf("Expected scvgTrace to equal %+v. Got %+v instead.", expectedScvgTrace, scvgTrace)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}

func TestParserNonMatchingInput(t *testing.T) {
	line := "INFO: test"

	runParserWith(line)

	select {
	case <-parser.GcChan:
		t.Fatalf("Unexpected trace result. This input should not trigger gcChan.")
	case <-parser.ScvgChan:
		t.Fatalf("Unexpected trace result. This input should not trigger scvgChan.")
	case <-parser.NoMatchChan:
		return
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}

func TestParserWait(t *testing.T) {
	line := "INFO: wait"
	parser := runParserWith(line)

	select {
	case <-parser.Done:
		return
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Execution timed out.")
	}
}
