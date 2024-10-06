package FindFile

import (
	Rule2 "CodeScan/CommonVul/Rule"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFileByPHP(dir string, outputfile string, rules []string) {
	var fileList []string

	// 使用filepath.Walk遍历目标目录，跳过黑名单中的目录，收集所有.java文件的路径
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//如果f是一个文件夹
		if f.IsDir() {
			//继续进行遍历，如果在黑名单中的话就进行跳过
			if Rule2.MatchRule(path, Rule2.PathBlackPhp) {
				return filepath.SkipDir
			}
			//如果文件存在的话就进行遍历 否则就进行判断，如果是java或者jsp后缀就添加到文件列表
		} else if strings.HasSuffix(f.Name(), ".php") || strings.HasSuffix(f.Name(), ".mds") {
			fileList = append(fileList, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", dir, err)
		return
	}

	// 检查遍历目录过程中的错误
	Check(err)

	// 创建或打开输出文件，以追加模式写入
	basedir := "./results/"
	err1 := os.MkdirAll(basedir, os.ModePerm)
	if err1 != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	outputfile = basedir + outputfile
	outputFile, err := os.OpenFile(outputfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)
	defer outputFile.Close() // 确保文件在函数返回前被关闭

	for _, file := range fileList {
		f, err := os.Open(file)
		Check(err)
		defer f.Close() // 确保文件在处理完后被关闭

		// 使用bufio.Scanner读取文件内容，为大文件读取优化
		scanner := bufio.NewScanner(f)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 10*1024*1024) // 设置一个更大的最大行大小

		lineNumber := 1     // 行号，用于标识匹配行的位置
		var lastFile string // 记录上一次输出的文件，用于控制输出格式
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text()) // 去除行首尾空白字符
			for _, rule := range rules {
				// 检查当前行是否包含规则，且规则匹配函数Rule.MatchRule返回true
				if strings.Contains(strings.ToLower(line), strings.ToLower(rule)) {
					if Rule2.MatchRule(line, Rule2.LineBlack) {
						break // 如果规则匹配，则跳出内层循环，处理下一行
					} //xxx    x
					if !Rule2.RemoveStaticVar(strings.ToLower(line), strings.ToLower(rule)) {
						break // 如果是静态变量则不做匹配
					}

					// 如果当前行是新文件的第一行且包含规则，则输出文件完整信息
					if lastFile != file {
						_, err := outputFile.WriteString(fmt.Sprintf("====================================================================\n\n"))
						_, err = outputFile.WriteString(fmt.Sprintf("file [%s]\n%d : %s\n\n", file, lineNumber, line))
						Check(err)
						lastFile = file
					} else {
						// 如果当前行不是新文件的第一行，仅输出行号和内容
						_, err := outputFile.WriteString(fmt.Sprintf("====================================================================\n\n"))
						_, err = outputFile.WriteString(fmt.Sprintf("%d : %s\n\n", lineNumber, line))
						Check(err)
					}

				}
			}
			lineNumber++
		}

		// 检查扫描过程是否出错
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

}
