package SSRF

import (
	"CodeScan/FindFile"
	"fmt"
)

func PHP_SSRF(dir string) {
	FindFile.FindFileByPHP(dir, "SSRF.txt", []string{
		"curl_exec(",
	})
	fmt.Println("PHPSSRF分析完成")
}
