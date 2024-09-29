package Fastjson

import (
	"CodeScan/FindFile"
	"fmt"
)

func Parsecheck(dir string) {
	FindFile.FindFileByJava(dir, "fastjson.txt", []string{".parseObject("})
	fmt.Println("fastjson分析完成")

}
