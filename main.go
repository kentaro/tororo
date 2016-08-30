package main

import (
	"fmt"
	"github.com/kentaro/tororo/parser"
	"github.com/kentaro/tororo/vm"
	"os"
	"strings"
)

func main() {
	result := parser.Parse(strings.NewReader(os.Args[1]))
	fmt.Println(vm.Run(result))
}
