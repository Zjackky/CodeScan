package FileWrite

import (
	"CodeScan/FindFile"
	"fmt"
)

func Write(dir string) {
	FindFile.FindFileByPHP(dir, "FileWrite.txt", []string{
		"file_put_contents(",
	})
	fmt.Println("PHP文件写入分析完成")

}
