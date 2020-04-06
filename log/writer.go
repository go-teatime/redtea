package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

func NewLogByDayWriter(filePath, filePrefix string) io.Writer {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		fmt.Printf("os.MkdirAll: %s", err)
		return nil
	}
	date := time.Now().Format("2006-01-02")
	p := path.Join(filePath, fmt.Sprintf("%s%s.log", filePrefix, date))
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}
	return f
}
