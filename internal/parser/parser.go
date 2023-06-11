package parser

import (
	"bufio"
	"github.com/golangFame/gcvis/pkg/trace"
	"io"
	"regexp"
	"strconv"
)

const SCVGRegexp = `scvg\d+: inuse: (?P<inuse>\d+), idle: (?P<idle>\d+), sys: (?P<sys>\d+), released: (?P<released>\d+), consumed: (?P<consumed>\d+) \(MB\)`

var svgRegExp = regexp.MustCompile(SCVGRegexp)

type Parser struct {
	reader      io.Reader
	GcChan      chan *trace.Gctrace
	ScvgChan    chan *trace.Scvgtrace
	NoMatchChan chan string
	Done        chan bool

	Err error

	scvgRegexp *regexp.Regexp
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		reader:      r,
		GcChan:      make(chan *trace.Gctrace, 1),
		ScvgChan:    make(chan *trace.Scvgtrace, 1),
		NoMatchChan: make(chan string, 1),
		Done:        make(chan bool),
	}
}

func (p *Parser) Run() {
	sc := bufio.NewScanner(p.reader)

	for sc.Scan() {
		line := sc.Text()

		if result := gcRegex.FindStringSubmatch(line); result != nil {
			p.GcChan <- parseGCTrace(gcRegex, result)
			continue
		}

		if result := svgRegExp.FindStringSubmatch(line); result != nil {
			p.ScvgChan <- parseSCVGTrace(result)
			continue
		}

		p.NoMatchChan <- line
	}

	p.Err = sc.Err()

	close(p.Done)
}

func parseGCTrace(gcre *regexp.Regexp, matches []string) *trace.Gctrace {
	matchMap := getMatchMap(gcre, matches)

	return &trace.Gctrace{
		Heap1:        silentParseInt(matchMap["Heap1"]),
		ElapsedTime:  silentParseFloat(matchMap["ElapsedTime"]),
		STWSclock:    silentParseFloat(matchMap["STWSclock"]),
		MASclock:     silentParseFloat(matchMap["MASclock"]),
		STWMclock:    silentParseFloat(matchMap["STWMclock"]),
		STWScpu:      silentParseFloat(matchMap["STWScpu"]),
		MASAssistcpu: silentParseFloat(matchMap["MASAssistcpu"]),
		MASBGcpu:     silentParseFloat(matchMap["MASBGcpu"]),
		MASIdlecpu:   silentParseFloat(matchMap["MASIdlecpu"]),
		STWMcpu:      silentParseFloat(matchMap["STWMcpu"]),
	}
}

func parseSCVGTrace(matches []string) *trace.Scvgtrace {
	matchMap := getMatchMap(svgRegExp, matches)

	return &trace.Scvgtrace{
		Inuse:    silentParseInt(matchMap["inuse"]),
		Idle:     silentParseInt(matchMap["idle"]),
		Sys:      silentParseInt(matchMap["sys"]),
		Released: silentParseInt(matchMap["released"]),
		Consumed: silentParseInt(matchMap["consumed"]),
	}
}

// Transform our matches in a readable hash map.
//
// The resulting hash map will be something like { "Heap1": 123 }
func getMatchMap(re *regexp.Regexp, matches []string) map[string]string {
	matchingNames := re.SubexpNames()[1:]
	matchMap := map[string]string{}
	for i, value := range matches[1:] {
		if matchingNames[i] == "" {
			continue
		}
		matchMap[matchingNames[i]] = value
	}
	return matchMap
}

func silentParseInt(value string) int64 {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return intVal
}

func silentParseFloat(value string) float64 {
	floatVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return float64(0)
	}

	return float64(floatVal)
}
