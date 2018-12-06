//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"./projects/oop"
	"fmt"
)

func main()  {
	//user := oop.Attrs{1,"sss","ddd","ddd"}
	//	//uid := user.GetUid()
	//	//user.Login()
	//	//
	//	//fmt.Println(user)
	//	//fmt.Println(uid)

	//animals := []oop.Animal{oop.Cat{}, oop.Dog{},oop.Dock{}}
	//for _, animal := range animals {
	//	fmt.Println(animal.Speak())
	//}
	//
	//oop.Test("string")
	//oop.Test(123)
	//oop.Test(true)

	names := []string{"a", "0", "c","DDD"}
	fmt.Println(len(names))
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	oop.PrintAll(vals)
}