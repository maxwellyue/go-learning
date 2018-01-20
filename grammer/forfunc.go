package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf( "最大值是 : %d\n", ret )


	//演示返回多个值
	c, d := swap("Mahesh", "Kumar")
	fmt.Println(c, d)

	//演示数组
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}


	//演示指针
	var e int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &e  /* 指针变量的存储地址 */

	fmt.Printf("e 变量的地址是: %x\n", &e  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	//演示空指针
	var  ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr  )

	if(ptr != nil) {}    /* ptr 不是空指针 */
	if(ptr == nil) {}  /* ptr 是空指针 */
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

