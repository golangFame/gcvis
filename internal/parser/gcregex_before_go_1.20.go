//go:build !go1.20
// +build !go1.20

package parser

import "regexp"

const GCRegexExp = `gc #?\d+ @(?P<ElapsedTime>[\d.]+)s \d+%: (?P<STWSclock>[^+]+)\+(?P<MASclock>[^+]+)\+(?P<STWMclock>[^+]+) ms clock, (?P<STWScpu>[^+]+)\+(?P<MASAssistcpu>[^+]+)/(?P<MASBGcpu>[^+]+)/(?P<MASIdlecpu>[^+]+)\+(?P<STWMcpu>[^+]+) ms cpu, \d+->\d+->\d+ MB, (?P<Heap1>\d+) MB goal, \d+ P`

var gcRegex = regexp.MustCompile(GCRegexExp)
