/**
 * Created with IntelliJ IDEA.
 * User: leelin
 * Date: 13-3-30
 * Time: 下午3:51
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func main() {
	for {
		var str string
		var str2 string = ""
		fmt.Scanf("%s", &str)
		u := []rune(str)
		for i := len(u) - 1; i > -1; i-- {
			str2 += string(u[i])
		}
		fmt.Println(str2)
		return
	}
}
