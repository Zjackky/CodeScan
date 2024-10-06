package Frame_Analysis

import (
	Rule2 "CodeScan/CommonVul/Rule"
	"CodeScan/FindFile"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FrameAnalysiser(dir string) {
	var result []string

	mybatis := false
	spring := false
	struts := false
	shiro := false
	CKeditor := false

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename := strings.ToLower(info.Name())

		if !info.IsDir() {

			// xml黑名单匹配
			if Rule2.MatchRule(filename, Rule2.XmlBlack) {
				return nil
			}

			if strings.HasSuffix(info.Name(), ".java") || strings.HasSuffix(info.Name(), ".xml") {

				if !mybatis && strings.Contains(filename, "mybatis") {
					result = append(result, "[+] MyBatis 框架 "+info.Name()+"\n")
					mybatis = true
				}

				if !spring && (strings.Contains(filename, "spring") || strings.Contains(filename, "controller")) {

					result = append(result, "[+] Spring 框架 "+info.Name()+"\n")
					spring = true
				}

				if !struts && strings.Contains(filename, "struts") {
					result = append(result, "[+] Struts 框架 "+info.Name()+"\n")
					struts = true
				}

				if !shiro && strings.Contains(filename, "shiro") {
					result = append(result, "[+] Shiro 框架 "+info.Name()+"\n")
					shiro = true
				}

				if !CKeditor && strings.Contains(filename, "ckeditor") {
					result = append(result, "[+] CKeditor 上传组件 "+info.Name()+"\n")
					CKeditor = true
				}

			}

			if !struts && info.IsDir() && strings.Contains(filename, "action") {
				result = append(result, "[+] Struts 框架 "+info.Name()+"\n")
				struts = true
			}

			if !CKeditor && info.IsDir() && strings.Contains(filename, "ckeditor") {
				result = append(result, "[+] CKeditor 上传组件 "+info.Name()+"\n")
				CKeditor = true
			}
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	output := strings.Join(result, "\n")
	// 创建或打开输出文件，以追加模式写入
	basedir := "./results/"

	// 检查目录是否存在
	if _, err := os.Stat(basedir); os.IsNotExist(err) {
		// 如果目录不存在，则创建
		err := os.MkdirAll(basedir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
	outputfile := "Frame_Analysiser.txt"
	outputfile = basedir + outputfile
	err = ioutil.WriteFile(outputfile, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("框架分析完成")
}

func WebXmlScan(dir string, rules []string) {

	// 创建或打开输出文件，以追加模式写入
	outputFile, err := os.OpenFile("Frame_Analysiser.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	FindFile.Check(err)
	defer outputFile.Close() // 确保文件在函数返回前被关闭

	var webXmlPath string

	// 遍历目录及其子目录下的所有文件
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果找到 "web.xml" 文件，记录其路径并停止遍历
		if !info.IsDir() && strings.EqualFold(info.Name(), "web.xml") {
			webXmlPath = path
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// 如果没有找到 "web.xml" 文件，结束函数
	if webXmlPath == "" {
		return
	}

	f, err := os.Open(webXmlPath)
	FindFile.Check(err)
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
				}
				// 如果当前行是新文件的第一行且包含规则，则输出文件完整信息
				if lastFile != webXmlPath {
					_, err := outputFile.WriteString(fmt.Sprintf("====================================================================\n\n"))
					_, err = outputFile.WriteString(fmt.Sprintf("file [%s]\n%d : %s\n\n", webXmlPath, lineNumber, line))
					FindFile.Check(err)
					lastFile = webXmlPath
				} else {
					// 如果当前行不是新文件的第一行，仅输出行号和内容
					_, err := outputFile.WriteString(fmt.Sprintf("====================================================================\n\n"))
					_, err = outputFile.WriteString(fmt.Sprintf("%d : %s\n\n", lineNumber, line))
					FindFile.Check(err)
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
