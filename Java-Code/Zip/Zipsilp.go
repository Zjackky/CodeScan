package Zip

import (
	"CodeScan/FindFile"
	"fmt"
)

func Zipsilp(dir string) {
	FindFile.FindFileByJava(dir, "zip.txt", []string{"zipEntry.getName(", "ZipUtil.unpack(", "ZipUtil.unzip(", "entry.getName()", "AntZipUtils.unzip(", "zip.getEntries()"})
	fmt.Println("Zipsilp分析完成")
}
