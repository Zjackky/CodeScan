package Rule

import "strings"

func MatchRule(str string, blackList []string) bool {
	//1.对传入的内容包含相关的黑名单关键字则不写入文件
	for _, v := range blackList {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
