package Upload

import (
	"CodeScan/CommonVul/Rule"
	"CodeScan/FindFile"
	"fmt"
)

func JavaUpload_check(dir string) {
	//FindFile.FindFileByJava(dir, "upload.txt", []string{"new File(", "MultipartFile", "upload", ".getOriginalFilename(", ".transferTo("})
	FindFile.FindFileByJava(dir, "upload.txt", Rule.JavaUploadRuleList)
	fmt.Println("上传分析完成")
}

func PHPUpload_check(dir string) {
	FindFile.FindFileByPHP(dir, "upload.txt", Rule.PHPUploadRuleList)
	fmt.Println("上传分析完成")
}
