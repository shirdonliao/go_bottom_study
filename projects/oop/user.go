//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package oop

import "fmt"

type Attrs struct {
	Uid int
	Username string
	Photo string
	Wechat string
}

func (attr *Attrs) GetUid() int {
	Uid := attr.Uid
	return Uid
}

func (attr *Attrs) Login()  {
	fmt.Println("aaaaa","logining")

}



