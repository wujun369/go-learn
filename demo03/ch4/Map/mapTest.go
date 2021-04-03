package main

import (
	"fmt"
	"sort"
)

func main() {

	//m := make(map[string]int) //初始化方式一
	//m["a"] = 1
	//m["b"] = 2
	//m["c"] = 3
	//m["d"] = 4
	//m["e"] = 5
	//m["f"] = 6
	//
	//m1 := map[string]int{//初始化方式二
	//	"a": 1,
	//	"b": 2,
	//	"c": 3,
	//	"d": 4,
	//	"e": 5,
	//	"f": 6,
	//}
	//
	//for key, value := range m { //map默认情况下的遍历顺序为乱序的
	//	fmt.Printf("%s:%d ", key, value)
	//}
	//fmt.Println()
	//
	//sortMapAndPrint(m1)
	//fmt.Println("------------------------------------------------------")

	//m2 := make(map[string]bool)
	//scan := bufio.NewScanner(os.Stdin)
	//
	//for scan.Scan() {//只打印第一次输入的字符串
	//	line := scan.Text()
	//	if !m2[line] {
	//		m2[line] = true
	//		fmt.Println(line)
	//	}
	//}
	//if err := scan.Err(); err != nil {
	//	fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println("------------------------------------------------------")
	//
	//m3 := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	line := input.Text()
	//	add(line,m3)
	//	fmt.Println(count(line,m3))
	//}

	var graph map[string]map[string]bool
	graph = make(map[string]map[string]bool)
	graph["v0"] = map[string]bool{
		"v1":true,
		"v3":false,

	}
	graph["v1"] = map[string]bool{
		"v0":true,
		"v2":false,
		"v3":true,

	}
	graph["v2"] = map[string]bool{
		"v0":true,
		"v1":false,
		"v3":false,

	}
	graph["v3"] = map[string]bool{}

	//fmt.Println(hasEdge("v0","v2",graph))
	addEdge("v3","v2",graph)
	//fmt.Println(hasEdge("v0","v2",graph))
	//addEdge("v1","v2",graph)
	//fmt.Println(hasEdge("v1","v2",graph))
	for k , v := range graph{
		fmt.Println(k,":",v)
	}
}

func sortMapAndPrint(m map[string]int) { //按key的顺序输出map
	var names []string
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s:%d ", name, m[name])
	}
}
func k(list []string) string {//将slice转换为string
	return fmt.Sprintf("%q",list)
}
func add(str string,m map[string]int) {//输入字符串的个数 + 1
	m[str]++
}
func count(str string,m map[string]int) int {//返回某行输入字符串的个数
	return m[str]
}

func addEdge(from,to string,graph map[string]map[string]bool){//为图的节点 from 到 to 添加一条边
	edges := graph[from]
	fmt.Println(edges == nil)
	if edges == nil { //若没有边，则添加边
		edges := make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from ,to string,graph map[string]map[string]bool) bool {//判断两点之间存不存在边
	return graph[from][to]
}

/*
	1、通过 fmt.Sprintf 可以将slice转换为字符串并返回
 */