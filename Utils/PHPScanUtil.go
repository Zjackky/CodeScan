package Utils

import (
	"CodeScan/CommonVul/Rce"
	"CodeScan/CommonVul/Upload"
	"CodeScan/PHP-Code/FileRead"
	"CodeScan/PHP-Code/Include"
	"CodeScan/PHP-Code/PHPSql"
	"CodeScan/PHP-Code/SSRF"
	"CodeScan/PHP-Code/Unserialize"
	"github.com/cheggaaa/pb/v3"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func PHP_Codeing() {
	StartTime = time.Now()

	// 所有要执行的扫描函数
	scanFuncs := []func(string){
		Upload.PHPUpload_check,
		Rce.PHPRce,
		PHPSql.Sqlcheck,
		FileRead.Read,
		Unserialize.Unserialize,
		SSRF.PHP_SSRF,
		Include.Include,
	}

	var wg sync.WaitGroup
	wg.Add(len(scanFuncs)) // 根据方法数量动态调整 goroutine 数量
	progressBar = pb.New(len(scanFuncs)).SetRefreshRate(time.Millisecond * 100).Start()
	// 启动 goroutine 来执行扫描任务
	for _, scanFunc := range scanFuncs {
		go scanDirectory(scanFunc, *Dir, &wg)
	}

	wg.Wait()
	progressBar.Finish()
	// 清理空文件
	root := "./" // 设置要检查的目录
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			if info.Size() == 0 {
				os.Remove(path)
			}
		}
		return nil
	})
}
