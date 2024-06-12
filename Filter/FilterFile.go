package Filter

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FilterFile(filterContent string, dir string) {
	outfile := "FilterResult.txt"

	// 打开或创建FilterResult.txt文件
	resultFile, err := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening result file: %v\n", err)
		return
	}
	defer resultFile.Close()

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", dir, err)
		return
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略目录
		if info.IsDir() {
			return nil
		}

		// 获取文件扩展名
		ext := filepath.Ext(path)
		// 仅处理后缀为jsp, java, php的文件
		if ext == ".jsp" || ext == ".java" || ext == ".php" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			containsContent := false
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), filterContent) {
					containsContent = true
					break
				}
			}

			// 如果文件不包含filterContent，则写入结果文件
			if !containsContent {
				absPath, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				_, err = resultFile.WriteString(absPath + "\n")
				if err != nil {
					return err
				}
			}
		}
		// 跳过其他文件类型
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
	}
}
