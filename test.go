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

	//names := []string{"a", "0", "c","DDD"}

	//num := []int{1,2,3,4}

	//twoArr := [2][2] int{
	//	{1,2},
	//	{3,4},
	//}
	//vals := make([]interface{},len(twoArr))

	shirdon := map[string]string {
		"name":"shirdon",
		"age":"28",
	}

	for _,va:= range shirdon {
		fmt.Println(va)
	}


	//for _,val := range &twoArr {
	//	for _,v := range &val {
	//		v = v+1
	//		fmt.Println(v)
	//	}
	//}
	//fmt.Println(twoArr)

	//vals := make([]interface{}, len(names))
	//for i, v := range twoArr {
	//	vals[i] = v
	//}
	//base.Foreach(vals)
	//oop.PrintAll(vals)

	//vals := make([]interface{}, len(num))
	//for _, v := range names {
	//	//vals[i] = v
	//	fmt.Println(v)
	//}
	//oop.Params(vals)
}