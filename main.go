package main

import "fmt"


func main() {
	names := [] string {"nico", "lynn"}
	names = append(names, "lalala")
	fmt.Println(names)
}