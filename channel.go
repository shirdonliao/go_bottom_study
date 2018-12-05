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
	"runtime"
)

func main()  {
	chan_n := make(chan int ,1)
	chan_c := make(chan int, 2)
	done := make(chan struct{})


	go func() {
		char_seq := []string{"A","B","C","D","E","F","G","H","I","J","K"}
		for i := 0; i < 10; i += 2 {
			a := <-chan_n
			fmt.Print(a)
			fmt.Print(char_seq[i])
			fmt.Print(char_seq[i+1])
			chan_c <- 2
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 1; i < 11; i += 2 {
			b := <-chan_c
			fmt.Print(b)
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- 1
		}
	}()

	chan_n <- 1 //先发送到chan_n func 第一个方法永远是先接收再发送
	<-done
	t := runtime.NumGoroutine()
	fmt.Println("\n:aaa",t)

}
