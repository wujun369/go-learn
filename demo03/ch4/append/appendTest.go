package main

import "fmt"

func main() {
	var runes []rune

	for _, r := range "hello 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q ", runes)
	runes = append(runes, '!')
	fmt.Println()
	fmt.Println("cap=", cap(runes)) //为分片分配内存，当内存不足是，将分配当前内存的两倍从而达到扩容的效果，分配后的底层数据不是原来的数据
	fmt.Println()

	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

}
