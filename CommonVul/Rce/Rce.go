package Rce

import (
	"CodeScan/CommonVul/Rule"
	"CodeScan/FindFile"
	"fmt"
)

func JavaRce(dir string) {
	FindFile.FindFileByJava(dir, "rce.txt", Rule.JavaRceRuleList)
	fmt.Println("RCE分析完成")
}

func PHPRce(dir string) {
	FindFile.FindFileByPHP(dir, "rce.txt", Rule.PHPRceRuleList)
	fmt.Println("RCE分析完成")
}
