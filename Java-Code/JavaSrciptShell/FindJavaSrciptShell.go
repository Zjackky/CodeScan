package JavaSrciptShell

import (
	"CodeScan/FindFile"
	"fmt"
)

func FindJavaSrciptShell(dir string) {
	FindFile.FindFileByJava(dir, "jshell.txt", []string{".getEngineByName(\"JavaScript\""})
	fmt.Println("JavaSrciptShell 分析完成")
}
