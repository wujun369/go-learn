package main

import "fmt"

func main() {
	//var data = []string{"hello","","world","hello1","1","world1"}
	//fmt.Println(noneempty(data))
	//fmt.Println(data)
	//fmt.Println(noneempty2(data))
	//fmt.Println(data)
	//str := noneempty2(data)
	//top := str[len(str) - 1]
	//fmt.Println(top)
	//str = str[:len(str) - 1]
	//fmt.Println(str)
	i := []int{1,2,3,4,5,6,7}
	fmt.Println(i)
	i = remove2(i,2)
	fmt.Println(i)

	fmt.Println("---------------------------------------------------------")

	ints := make([]int, 10,20)//make([]T, len, cap)，其中len为切片长度，cap为容量且可以省略，如果没有省略，声明的值必须大于len
	fmt.Println(ints)

	reverse(i)
	fmt.Println(i)
}
func noneempty(strings []string) []string {//返回非空字符切片，返回的切片和输入切片共享底层数据
	i := 0
	for _,v := range strings{
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[:i]
}
func noneempty2(strings []string) []string{//返回非空字符切片，返回的切片为新的切片
	var str []string
	for _,v := range strings{
		if v != "" {
			str = append(str,v)
		}
	}
	return str
}
func remove(slice []int,i int) []int {//删除字符切片的某一个元素，保持原有的数据，返回的切片为和输入切片共享一个底层数据
	copy(slice[i:],slice[i+1:])
	return slice[:len(slice) - 1]
}

func remove2(slice []int,i int) []int {//删除字符切片的某一个元素，返回的切片为和输入切片共享一个底层数据
	slice[i] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}
func reverse(s []int) {//反转数据
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
	slice和数组的区别：
		1、slice不需要指定明确的长度，会隐式的创建一个合适大小的数据；
		2、slice的底层是通过数组实现的；
		3、slice不能通过 ‘==' 来判断两个slice，而数组可以，唯一可以通过 ‘==’ 比较的是和nil
 */