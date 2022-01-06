package main

import "fmt"

func main() {
	type Map1 map[string]interface{}
	type Map2 map[string]int
	m := Map1{"foo": Map2{"first": 1}, "boo": Map2{"second": 2}}
	for k, v := range m {
		fmt.Printf("%v:  type %T   ", k, k)
		fmt.Printf("%v:  type %T\n", v, v)
	}
}
