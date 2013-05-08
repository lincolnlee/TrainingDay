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
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rn := r.Intn(100)

	fmt.Println("Random number is：" + fmt.Sprintf("%d", rn))

	temp := 0

	for {
		fmt.Scanf("%d", &temp)
		if temp > rn {
			fmt.Println("more")
		} else if temp < rn {
			fmt.Println("less")
		} else {
			fmt.Println("right")
			return
		}
	}
}
