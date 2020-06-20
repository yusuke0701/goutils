package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

var traceID string

// SetTraceID set the trace id for the Cloud Logging.
// ref: https://cloud.google.com/trace/docs/setup/go?hl=ja
func SetTraceID(r *http.Request) {
	traceID = strings.SplitN(r.Header.Get("X-Cloud-Trace-Context"), "/", 2)[0]
}

type severity string

// severity list
// ref: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry?hl=ja#LogSeverity
const (
	severityInfo     severity = "INFO"
	severityError    severity = "ERROR"
	severityWarn     severity = "WARNING"
	severityCritical severity = "CRITICAL"
	severityDebug    severity = "DEBUG"
)

// LogDebugf デバッグログを出力する
func LogDebugf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityDebug, fmt.Sprintf(format, v...))
}

// LogDebug デバッグログを出力する
func LogDebug(ctx context.Context, v interface{}) {
	logPrintf(ctx, severityDebug, fmt.Sprint(v))
}

// LogInfof 情報ログを出力する
func LogInfof(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityInfo, fmt.Sprintf(format, v...))
}

// LogInfo 情報ログを出力する
func LogInfo(ctx context.Context, v interface{}) {
	logPrintf(ctx, severityInfo, fmt.Sprint(v))
}

// LogErrorf エラーログを出力する
func LogErrorf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityError, fmt.Sprintf(format, v...))
}

// LogError エラーログを出力する
func LogError(ctx context.Context, v interface{}) {
	logPrintf(ctx, severityError, fmt.Sprint(v))
}

// LogWarningf 警告ログを出力する
func LogWarningf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityWarn, fmt.Sprintf(format, v...))
}

// LogWarning 警告ログを出力する
func LogWarning(ctx context.Context, v interface{}) {
	logPrintf(ctx, severityWarn, fmt.Sprint(v))
}

// LogCriticalf 重大エラーログを出力する
func LogCriticalf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityCritical, fmt.Sprintf(format, v...))
}

// LogCritical 重大エラーログを出力する
func LogCritical(ctx context.Context, v interface{}) {
	logPrintf(ctx, severityCritical, fmt.Sprint(v))
}

func logPrintf(ctx context.Context, s severity, msg string) {
	// プレフィックスに余計な文字列がつかないようにLoggerオブジェクトを作成
	logger := log.New(getWriter(), "", 0)

	// 構造化ロギングに出力する情報
	var info, _ = getCallerInfo(3)

	// 設定可能な特殊フィールドについては次を参照
	// https://cloud.google.com/logging/docs/agent/configuration?hl=ja#special-fields
	entry := map[string]interface{}{
		"message":                      msg,
		"severity":                     s,
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", ProjectID, traceID),
		"logging.googleapis.com/sourceLocation": map[string]interface{}{
			"file":     info.File,
			"line":     info.Line,
			"function": info.FnName,
		},
	}
	payload, _ := json.Marshal(entry)
	logger.Println(string(payload))
}

var getWriter = func() io.Writer {
	return os.Stdout
}

type callerInfo struct {
	File   string
	Line   int
	FnName string
}

var getCallerInfo = func(skipFrame int) (callerInfo, bool) {
	pc, file, line, ok := runtime.Caller(skipFrame + 1)
	if !ok {
		return callerInfo{}, ok
	}
	f := runtime.FuncForPC(pc)
	return callerInfo{
		File:   path.Base(file),
		Line:   line,
		FnName: f.Name(),
	}, true
}
