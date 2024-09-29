package JNDI

import (
	"CodeScan/FindFile"
	"fmt"
)

func Jndi(dir string) {
	FindFile.FindFileByJava(dir, "jndi.txt", []string{".lookup("})
	fmt.Println("JNDI分析完成")
}
