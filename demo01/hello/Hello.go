package hello

import "fmt"

type Hello struct {
	Name string
}

func (h *Hello) PrintName()  {
	fmt.Println(h.Name)
}