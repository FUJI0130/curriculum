package config

// StackTraceFilterKeywords contains keywords that are used to filter the stack trace.
var StackTraceFilterKeywords = []string{
	"gin-gonic",
	"net/http",
	"runtime.goexit",
	// その他フィルタリングしたいキーワードをここに追加
}
