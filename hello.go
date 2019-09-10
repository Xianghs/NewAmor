package main

import (
	"fmt"
)

const PI float32 = 3.141
const PORT int = 8080
const (
	IP   string = "127.0.0.1"
	Gate int    = 10080
)

//var b = 30 //全局变量
func main() {
	/*
		var a int = 20 //var+名称+类型
		var c int
		c = 40
		fmt.Printf("a=%d,b=%d,c=%d", a, b, c)

		//速记声明(只能放在函数内、左侧至少有一个变量是新的)
		d := 40
		sex, num := "w", 12
		fmt.Println("The value is", d)
		fmt.Println(sex, num)

		//多变量声明
		var width, height int = 30, 40 //同种类型
		var (                          //不同类型
			name string = "shenhuan"
			age  int    = 20 //若不赋值则默认为0

		)
		fmt.Println(width, height, name, age)

		e, f := 20, 30
		fmt.Println("e is", e, "f is", f)
		f, g := 40, 50
		fmt.Println("f is", f, "g is", g)
		f, g = 80, 90
		fmt.Println("changed f is", f, "g is", g)
		r := Max(10, 20)
		fmt.Println("r=", r)
	*/
	/*var a int = 12
	var b float32 = 3.14
	c1 := complex(4, 8)
	c2 := 3 + 4i
	cadd := c1 + c2

	fmt.Println("sum:", cadd)
	fmt.Printf("The %.1f's type is %T \nsize id is %d\n", b, b, unsafe.Sizeof(a))
	*/
	/*first := "marry"
	last := "steven"
	name := first + " " + last
	fmt.Println("my name is", name)
	fmt.Println(len(name))
	*/
	/*
		var str string = "123456"
		//将字符串转化为整型
		intValue, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(intValue)
		}

		var phone int = 1234567123124
		phoneStr := strconv.Itoa(phone)
		fmt.Println(phoneStr)

	*/
	i := 12
	j := 13.87
	sum := float32(i) + float32(j)
	fmt.Println(sum)
	fmt.Println(Gate)

	//var defaultName = "Sam"
	//	type myString string
	//var customName myString = "Sam"
	//	fmt.Printf("%T", customName)
	//customName = defaultName
	const (
		a = iota
		b
		c
		d = "a"
		e
		f = 100
		g
		h = iota
		q
	)
	fmt.Println(a, b, c, d, e, f, g, h, q)
}

//func 函数名（参数列表）返回值{}
/*func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}