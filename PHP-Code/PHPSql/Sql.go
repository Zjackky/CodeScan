package PHPSql

import (
	"fmt"
	"os"
)

// check 函数用于检查错误，如果错误不为 nil 则触发 panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Sqlcheck 函数是我们的主函数，负责执行 SQL 检查的逻辑
func Sqlcheck(dir string) {
	// 检查是否存在 @Select 注解
	findSqlByCode(dir)

	fmt.Println("sql分析完成")
}

// writeToFile 函数用于将信息写入文件
func writeToFile(filename string, lines []string) {
	// 打开文件，如果文件不存在则创建，如果存在则追加写入
	outputFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer outputFile.Close()

	// 将每一行信息写入文件
	for _, line := range lines {
		_, err = outputFile.WriteString(fmt.Sprintf("%s\n", line))
		check(err)
	}
}
