package main

import (
	"fmt"
	"test/pkg"
)

func main() {
	uuid, _ := pkg.Hoge()
	fmt.Println(uuid, "c")
}
