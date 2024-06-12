package FileRead

import (
	"CodeScan/CommonVul/Rule"
	"CodeScan/FindFile"
	"fmt"
)

func Read(dir string) {
	FindFile.FindFileByPHP(dir, "FileRead_Phar.txt", Rule.PHPFileReadList)
	fmt.Println("PHP文件读取分析完成")

}
