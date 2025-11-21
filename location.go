package codelocationmatch

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseLocations(s string, filePathPattern *regexp.Regexp) []*Location {
	re := regexp.MustCompile(`(\S+?)(?:#L|:)(\d+)(?:-(\d+))?`)
	allMatches := re.FindAllStringSubmatch(s, -1)
	if allMatches == nil {
		return []*Location{}
	}
	var locations []*Location
	for _, matches := range allMatches {
		if len(matches) < 3 {
			continue
		}
		filePath := matches[1]
		if filePathPattern != nil {
			ret := filePathPattern.FindStringSubmatch(filePath)
			if len(ret) == 2 {
				filePath = ret[1]
			}
		}
		lineNo, _ := strconv.Atoi(matches[2])
		endLineNo := lineNo
		if len(matches) > 3 && matches[3] != "" {
			endLineNo, _ = strconv.Atoi(matches[3])
		}
		for i := lineNo; i <= endLineNo; i++ {
			locations = append(locations, NewLocation(filePath, i))
		}
	}
	return locations
}

func NewLocation(filePath string, lineNo int) *Location {
	return &Location{
		filePath: filePath,
		lineNo:   lineNo,
	}
}

type Location struct {
	filePath string
	lineNo   int
}

func LocationsToString(s []*Location) string {
	var sb strings.Builder
	for _, location := range s {
		sb.WriteString(fmt.Sprintf("%s:%d\n", location.filePath, location.lineNo))
	}
	return strings.TrimSpace(sb.String())
}
