package main

import  "fmt"

func init()  {
	fmt.Println("先执行init?")
}

func main()  {
	fmt.Println("老岳开始学习go")
}

/*

该程序会输出：

先执行init?
老岳开始学习go

 */
