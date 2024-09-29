package JarStatic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Jarstaticer(dir string) {

	// 检查目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", dir)
		return
	}

	// 存储找到的jar文件名
	jarFiles := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".jar" {
			jarFiles = append(jarFiles, filepath.Base(path))
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dir, err)
		return
	}

	// 读取 EvilJarList.txt 文件中的每一行
	configLines, err := ioutil.ReadFile("EvilJarList.txt")
	if err != nil {
		fmt.Println("Error reading EvilJarList.txt:", err)
		return
	}
	lines := strings.Split(string(configLines), "\n")

	// 打开 SuccessAttack.txt 文件准备写入
	file, err := os.Create("SuccessAttack.txt")
	if err != nil {
		fmt.Println("Error creating SuccessAttack.txt:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 遍历 jarFiles 中的每个JAR文件名
	for _, jarName := range jarFiles {
		// 遍历 config.txt 中的每一行
		for _, line := range lines {
			// 去除行尾的换行符
			trimmedLine := strings.TrimSpace(line)
			if jarName == trimmedLine {
				// 如果找到匹配项，写入 SuccessAttack.txt
				_, err := writer.WriteString(jarName + "\n")
				if err != nil {
					fmt.Println("Error writing to SuccessAttack.txt:", err)
					return
				}
				break // 找到匹配项后跳出内层循环
			}
		}
	}

	// 检查

	// 打开 jarFiles.txt 文件准备写入
	file1, err := os.Create("jarFiles.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file1.Close()

	// 使用文件写入操作
	for _, jarName := range jarFiles {
		if _, err := file1.WriteString(jarName + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

}
