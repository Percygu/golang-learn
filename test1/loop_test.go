package main

import (
	"fmt"
	"testing"
)

type Guest struct {
	id      int
	name    string
	surname string
	friends []int
}

func (self Guest) removeFriend(id int) {
	for i, other := range self.friends {
		if other == id {
			self.friends = append(self.friends[:i], self.friends[i+1:]...)
			break
		}
	}
}

func (self Guest) AddFriend(id int) {
	self.friends = append(self.friends, id)
	fmt.Println("-------------------------")
	fmt.Println(self.friends)
	fmt.Println("-------------------------")
}

func TestForRange(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	for i := 0; i < len(slice1); i++ {
		slice1[i]++
	}
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3, 4}
	for i := range slice2 {
		slice2[i]++
	}
	fmt.Println(slice1)

	slice3 := []int{1, 2, 3, 4}
	for _, val := range slice3 {
		val++
	}
	fmt.Println(slice3)

	test1 := Guest{0, "Echo", "大叔", []int{1, 2, 3, 4, 5}}
	fmt.Println(test1)
	test1.removeFriend(4)
	fmt.Println(test1)

	test2 := Guest{0, "ssss", "aaa", []int{1, 2, 3, 4, 5}}
	fmt.Println(test2)
	test2.AddFriend(10) // 函数内和函数外不一样，虽然传进去的是Guest的之拷贝，但是friend是指针类型，执行同一块内存，但是在append之后，出发了扩容，两个指针指向了不同的地址，原指针直指向原来的，函数里的指向扩容后的，所以打印内容不一样
	fmt.Println(test2)  // 还是打印原内容
}
