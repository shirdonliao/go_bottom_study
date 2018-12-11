//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import "fmt"


//对于go语言来说，设计最精妙的应该是interface了，直白点说interface是一组method的组合。至于更加详细的描述，本文不做介绍，今天谈谈空接口。
//空interface(interface{})不包含任何的method，因此所有的类型都实现了空interface。空interface在我们需要存储任意类型的数值的时候相当有用，
//有点类似于C语言的void*类型。请看下面的代码：

//这段代码声明了一个空接口的slice，这意味着它的值可以是任意类型，然后我们声明了两个map，一个是map[string]string，一个是map[string]int，
//然后在声明一个map的map类型，将这三个类型赋值给slice，使得slice可以存贮各种不同类型的数据，想想看，一个可变数组中，存储了一个key为string类型，
//value为int类型的map，又存储了一个key为string类型，value为string类型的map，还存储了一个map的map，这对c/c++转go的程序员们来说是多么让人吃惊。


func main()  {
	slice := make([]interface{}, 10)
	map1 := make(map[string]string)
	map2 := make(map[string]int)
	map1["Command"] = "ping"
	map2["TaskID"] = 1
	map3 := make(map[string]map[string]string)
	map3["mapValue"] = map1

	outMap := make(map[string]interface{},5)
	slice[0] = map2
	slice[1] = map1
	slice[3] = map3

	outMap["a"] = map1
	outMap["b"] = map2
	outMap["c"] = map3
	//fmt.Println(slice[0])
	//fmt.Println(slice[1])
	//fmt.Println(slice[3])

	fmt.Println(outMap["a"])
	fmt.Println(outMap["b"])
	fmt.Println(outMap["c"])
}
