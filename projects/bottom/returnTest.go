//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
)

func say(hi bool) {
	if !hi {
		fmt.Println("false")
		return
	}
	fmt.Println("没有返回值")
}
func getStatus() (num int) {
	num = 999
	return num
}

func main() {
	say(true)
	say(false)
	fmt.Println(getStatus())
}
