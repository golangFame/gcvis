//go:build !go1.15
// +build !go1.15

package main

import (
	"regexp"
)

const GCRegexExp = `gc\d+\(\d+\): ([\d.]+\+?)+ us, \d+ -> (?P<Heap1>\d+) MB, \d+ \(\d+-\d+\) objects,( \d+ goroutines,)? \d+\/\d+\/\d+ sweeps, \d+\(\d+\) handoff, \d+\(\d+\) steal, \d+\/\d+\/\d+ yields`

var gcRegex = regexp.MustCompile(GCRegexExp)
