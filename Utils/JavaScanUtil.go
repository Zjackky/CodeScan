package Utils

import (
	"CodeScan/CommonVul/Rce"
	"CodeScan/CommonVul/Upload"
	"CodeScan/Java-Code/AMF"
	"CodeScan/Java-Code/Auth_Bypass"
	"CodeScan/Java-Code/El"
	"CodeScan/Java-Code/Fastjson"
	"CodeScan/Java-Code/Frame_Analysis"
	"CodeScan/Java-Code/JDBC"
	"CodeScan/Java-Code/JNDI"
	"CodeScan/Java-Code/JS"
	"CodeScan/Java-Code/JarStatic"
	"CodeScan/Java-Code/JavaSrciptShell"
	"CodeScan/Java-Code/Log4j"
	"CodeScan/Java-Code/ReadObject"
	"CodeScan/Java-Code/Reflect"
	"CodeScan/Java-Code/SSTI/FreeMarker"
	"CodeScan/Java-Code/Sql"
	"CodeScan/Java-Code/Zip"
	"github.com/cheggaaa/pb/v3"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func Java_Codeing() {
	StartTime = time.Now()
	// 所有要执行的扫描函数
	scanFuncs := []func(string){
		Frame_Analysis.FrameAnalysiser,
		Auth_Bypass.Auth,
		Zip.Zipsilp,
		JNDI.Jndi,
		Sql.Sqlcheck,
		Rce.JavaRce,
		Upload.JavaUpload_check,
		ReadObject.Readobjectcheck,
		El.Elcheck,
		Fastjson.Parsecheck,
		Reflect.ReflectCheck,
		Log4j.Log4j,
		AMF.AmfCheck,
		FreeMarker.FreeSsti,
		JDBC.FindJDBC,
		JavaSrciptShell.FindJavaSrciptShell,
		JarStatic.Jarstaticer,
		JS.Eval,
	}

	var wg sync.WaitGroup
	wg.Add(len(scanFuncs)) // 根据方法数量动态调整 goroutine 数量
	progressBar = pb.New(len(scanFuncs)).SetRefreshRate(time.Millisecond * 100).Start()
	// 启动 goroutine 来执行扫描任务
	for _, scanFunc := range scanFuncs {
		go scanDirectory(scanFunc, *Dir, &wg)
	}

	wg.Wait()
	progressBar.Finish()

	// 处理web.xml
	Frame_Analysis.WebXmlScan(*Dir, []string{"*.htm", "*.do", "*.action", "exclude"})

	// 清理空文件
	root := "./" // 设置要检查的目录
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			if info.Size() == 0 {
				os.Remove(path)
			}
		}
		return nil
	})
}
