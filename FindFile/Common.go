package FindFile

// check函数用于检查错误，如果错误不为nil，则会触发panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
