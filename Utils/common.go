package Utils

import (
	"github.com/cheggaaa/pb/v3"
	"strings"
	"sync"
	"time"
)

var (
	progressBar *pb.ProgressBar
	StartTime   time.Time
)

// scanDirectory 函数用于启动一个 goroutine 来扫描指定目录
func scanDirectory(scanFunc func(string), dir string, wg *sync.WaitGroup) {
	scanFunc(dir)
	progressBar.Increment()
	wg.Done()
}

func ClearDir(dir string) string {

	// 将 \ 转换为 /
	dir = strings.ReplaceAll(dir, `\\`, "/")
	dir = strings.ReplaceAll(dir, `\`, "/")

	return dir
}
