//package landarea
package main

import (
	"fmt"
	"os"
)

var(
	i int
	ii float64
	iii string
)

const (
	x = iota
	y = iota
	z = iota
	w
)

const v = iota

var enabled, disabled = true, false

func add1(a *int) (b int) {
	*a = *a+1
	b = *a
	return
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

var user = os.Getenv("USER")

//func init() {
//	if user == "" {
//		panic("no value for $USER")
//	}
//}
//
//func throwsPanic(f func()) (b bool) {
//
//	defer func() {
//		if x := recover(); x != nil {
//			b = true
//		}
//	}()
//	f()
//	return
//}

func f() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b")
}

func main() {
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)

	//c:=5+5i
	//fmt.Printf("Complex64 value: %v\n", c)

	username:="root"
	rename :="angle"
	rename ="jane"
	fmt.Printf("String: %v, %v\n", username, rename)

	s := "hello"
	s = "world"
	fmt.Printf("String: %v\n", s)

	fmt.Printf("x=%d\ny=%d\nz=%d\nw=%d\nv=%d\n", x, y, z, w, v)

	var arr [5]int
	arr[0] = 10
	arr[4] = 20
	fmt.Printf("This first is %d\n", arr[0])
	fmt.Printf("This last is %d\n", arr[4])

	//Slice
	test := [5]byte{'a', 'b', 'c', 'd', 'e'}
	var a, b []byte
	a = test[1:3]
	b = test[3:5]

	fmt.Printf("Slice: %v  %v\n", a, b)

	//Map
	numbers := make(map[string]int)
	numbers["test"] = 1
	numbers["test1"] = 2
	numbers["test2"] = 3
	fmt.Println("test2: ", numbers["test2"])

	maptest := map[string]float32{"a": 1, "b": 2, "c": 3}
	testArr, mapresult := maptest["a"]
	fmt.Printf("Test: %v %v\n", testArr, mapresult)
	if mapresult {
		fmt.Println("value: ", testArr, mapresult)
	}else {
		fmt.Println("Key不存在")
	}

	// make(T, args)
	// make用于内建类型(map、slice、channel)的内存分配，并返回一个有初始值(非0)的T类型

	// new用于各种类型的内存分配，返回指针
	//

	// func
	x := 3
	fmt.Println("X = ", x)
	x1 := add1(&x)
	fmt.Println("X + 1 = ", x1)
	fmt.Println("X = ", x)

	for i:=0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Printf("\n");

	slice := []int {1,2,3,4,5,7}
	fmt.Println("Slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)

	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)

	fmt.Println("c")
	defer func() {
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("e")
	}()
	f()
	fmt.Println("f")
}