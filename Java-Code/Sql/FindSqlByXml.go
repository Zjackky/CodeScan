package Sql

import (
	Rule2 "CodeScan/CommonVul/Rule"
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// findKeywordsInXMLFiles 函数用于检查 XML 文件中的关键字
func findSqlByXml(dir string) {
	xmlList := []string{}
	var lastFile string // 记录上一次输出的文件，用于控制输出格式

	// 使用 Walk 函数遍历目录，查找所有的 .xml 文件
	err := filepath.Walk(dir, func(path string, f fs.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".xml") {
			// xml黑名单匹配
			if Rule2.MatchRule(f.Name(), Rule2.XmlBlack) {
				return nil
			}
			xmlList = append(xmlList, path)
		}
		return nil
	})
	check(err)

	// 定义需要搜索的关键字
	keywords := []string{"${", "like '%${", "order by ${"} // 这里可以添加更多关键字

	// 遍历 XML 文件列表
	for _, file := range xmlList {
		foundKeywords := []string{}

		lineNumber := 1
		// 打开 XML 文件
		f, err := os.Open(file)
		check(err)
		defer f.Close()

		// 逐行扫描文件内容
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			// 检查每一行是否包含关键字，并且不包含黑名单中的关键字
			if Rule2.MatchRule(line, Rule2.XmlSqlBlack) {
				continue
			}
			// 检查每一行是否包含需要搜索的关键字
			for _, keyword := range keywords {
				if strings.Contains(line, keyword) {
					if lastFile != f.Name() {
						foundKeywords = append(foundKeywords, fmt.Sprintf("====================================================================\n"))
						foundKeywords = append(foundKeywords, fmt.Sprintf("file [%s]\n%d: %s", f.Name(), lineNumber, line))
						lastFile = f.Name()
					} else {
						foundKeywords = append(foundKeywords, fmt.Sprintf("====================================================================\n"))
						foundKeywords = append(foundKeywords, fmt.Sprintf("%d : %s", lineNumber, line))
					}
				}
			}
			lineNumber++
		}

		// 如果找到关键字，则将相关信息写入到 sql.txt 文件中
		if len(foundKeywords) > 0 {
			writeToFile("sql.txt", foundKeywords)
		}
	}
}
