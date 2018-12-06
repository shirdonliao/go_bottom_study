//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package oop

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct{}
func (c Cat) Speak() string {
	return "miao"
}

type Dog struct{}
func (d Dog) Speak() string {
	return "wangwang"
}

type Dock struct {}

func (do Dock) Speak() string  {
	return "guaguagua"
}

func Test(params interface{}) {
	fmt.Println(params)
}


