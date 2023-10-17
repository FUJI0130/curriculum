package support

import (
	"fmt"
	"strings"

	"github.com/FUJI0130/curriculum/src/core/config"
)

type StackTraceFilter struct{}

func (f *StackTraceFilter) RemoveLinesFromKeywords(err string) string {
	fullTrace := fmt.Sprintf("%+v", err)
	lines := strings.Split(fullTrace, "\n")
	var filteredLines []string

	for _, line := range lines {
		shouldSkip := false
		for _, keyword := range config.StackTraceFilterKeywords {
			if strings.Contains(line, keyword) {
				shouldSkip = true
				break
			}
		}
		if !shouldSkip {
			filteredLines = append(filteredLines, line)
		}
	}

	return strings.Join(filteredLines, "\n")
}

func (f *StackTraceFilter) ExtractLinesFromKeywords(err string) string {
	fullTrace := fmt.Sprintf("%+v", err)
	lines := strings.Split(fullTrace, "\n")
	var extractedLines []string

	for _, line := range lines {
		for _, keyword := range config.StackTraceFilterKeywords {
			if strings.Contains(line, keyword) {
				extractedLines = append(extractedLines, line)
				break
			}
		}
	}

	return strings.Join(extractedLines, "\n")
}

func (f *StackTraceFilter) ExtractNLinesFromStart(err string, n int) string {
	fullTrace := fmt.Sprintf("%+v", err)
	lines := strings.Split(fullTrace, "\n")

	if len(lines) < n {
		return fullTrace
	}
	return strings.Join(lines[:n], "\n")
}
