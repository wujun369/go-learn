package main

import "demo01/hello"

func main() {
	h := hello.Hello{}
	h.Name = "hello"
	h.PrintName()
}
