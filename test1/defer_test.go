package main

import (
	"fmt"
	"testing"
)

func printTest(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferTest1() {
	var a = 1
	defer fmt.Println(a)

	a = 2
	return
}

func deferTest2() {
	var arr = [3]int{1, 2, 3}
	defer printTest(&arr)

	arr[0] = 4
	return
}

func deferTest3() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func deferTest4() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func deferTest5() int {
	var i int
	defer func() {
		i++
	}()

	return 1
}

func deferTest6() int {
	var i int
	defer func() {
		i++
	}()

	return i
}

func deferTest7() (i int) {
	defer func() {
		i++
	}()
	return 0
}

func TestDefer(t *testing.T) {
	deferTest1()              // 1                值传递，调用defer的时候，a的值已经确定为1，后面改变a的值变为2，但是defer函数里传进去的形参值还是1
	deferTest2()              // 4,2,3            数组指针的拷贝，改变指针指向的内容
	fmt.Println(deferTest3()) // 2                第一步，将i变量赋值给result，在执行defer语句，result++,最后返回result
	fmt.Println(deferTest4()) // 2                同deferTest3()
	fmt.Println(deferTest5()) // 1                第一步，将1赋值给隐式返回值result，在执行defer语句，i++,并未修改result的值，所以值为1
	fmt.Println(deferTest6()) // 0                第一步：将变量i的值赋值给隐式返回值result，此时i为0，第二步：在执行defer语句，i++,并未修改result的值，所以值为0，第三步：返回result
	fmt.Println(deferTest7()) // 1                第一步：i=0，  第二步：i++，  第三步：return
}
