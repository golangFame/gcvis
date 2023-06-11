//go:build !go1.16
// +build !go1.16

package main

import (
	"regexp"
)

const GCRegexExp = `gc #?\d+ @(?P<ElapsedTime>[\d.]+)s \d+%: [\d.+/]+ ms clock, [\d.+/]+ ms cpu, \d+->\d+->\d+ MB, (?P<Heap1>\d+) MB goal, \d+ P`

var gcRegex = regexp.MustCompile(GCRegexExp)
