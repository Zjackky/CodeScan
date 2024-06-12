package Rule

import (
	"strings"
)

func RemoveStaticVar(content string, rule string) bool {
	// 找到rule在content的位置
	index := strings.Index(content, rule)
	if index == -1 {
		// 如果rule不在content中，返回false
		return false
	}

	// 截取遇到第一个)之前的数据
	substr := content[index : strings.Index(content[index:], ")")+index+1]

	// 判断该数据内容是否存在"
	if strings.Contains(substr, "\"") {
		// 如果存在"，再一层判断，是否存在+
		if strings.Contains(substr, "+") {
			// 如果满足+和"，则返回true
			return true
		} else {
			// 如果存在"，不存在+，则返回false
			return false
		}
	} else {
		// 如果都没有， 则返回true
		return true
	}
}
