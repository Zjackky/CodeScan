package Log4j

import (
	"CodeScan/FindFile"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Log4j(dir string) {

	log4j2 := false
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		filename := strings.ToLower(info.Name())

		if !info.IsDir() {

			if !log4j2 && strings.Contains(filename, "log4j") {

				log4j2 = true

				// 执行 FindFile.FindFileByJava 方法
				FindFile.FindFileByJava(dir, "log4j.txt", []string{"logger.info(", "log.info("})
				fmt.Println("Log4j2分析完成")
			}

		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}

}
