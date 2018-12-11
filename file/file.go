//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"io/ioutil"
	"os"
)


func main() {
	f, err := os.Open("./file.txt")
	if err != nil {
		return
	}
	defer f.Close()

	//utility 公共的
	b, err := ioutil.ReadAll(f)
	println(string(b))
}
