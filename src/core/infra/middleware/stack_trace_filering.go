package middleware

import (
	"fmt"
	"strings"

	"github.com/FUJI0130/curriculum/src/core/config" // your_project_pathを適切なものに変更してください
)

type StackTraceFilter struct{}

func (f *StackTraceFilter) FilterStackTrace(err error) string {
	// エラーからスタックトレースを文字列として取得
	fullTrace := fmt.Sprintf("%+v", err)

	// log.Printf("Full Stack Trace: %s", fullTrace)

	// 行ごとに分割
	lines := strings.Split(fullTrace, "\n")

	// 特定のキーワードでフィルタリング
	var filteredLines []string
	for _, line := range lines {
		shouldSkip := false
		for _, keyword := range config.StackTraceFilterKeywords { // 変数名を変更
			if strings.Contains(line, keyword) {
				shouldSkip = true
				break
			}
		}

		// if shouldSkip {
		// 	continue
		// }
		if shouldSkip {
			// log.Printf("Skipping line: %s", line)
			continue
		}
		// log.Printf("Keeping line: %s", line)
		filteredLines = append(filteredLines, line)
	}

	// フィルタリングしたスタックトレースを再結合
	return strings.Join(filteredLines, "\n")
}
