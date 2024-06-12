package Include

import (
	"CodeScan/FindFile"
	"fmt"
)

func Include(dir string) {
	FindFile.FindFileByPHP(dir, "Include.txt", []string{
		"include(",
	})
	fmt.Println("PHP文件包含分析完成")
}
