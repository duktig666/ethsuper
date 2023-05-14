// description:
// @author renshiwei
// Date: 2022/11/16 10:44

package util

import "strings"

// TrimLeft0x 去掉字符串左侧 0x
func TrimLeft0x(s string) string {
	return strings.TrimLeft(s, "0x")
}
