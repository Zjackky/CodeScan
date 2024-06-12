package AMF

import (
	"CodeScan/FindFile"
	"fmt"
)

func AmfCheck(dir string) {
	FindFile.FindFileByJava(dir, "AmfCheck.txt", []string{".readMessage("})
	fmt.Println("AMF检查完成")

}
