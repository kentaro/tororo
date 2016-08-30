package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/kentaro/tororo/parser"
	"github.com/kentaro/tororo/vm"
)

func main() {
    result := parser.Parse(strings.NewReader(os.Args[1]))
    fmt.Println(vm.Run(result))
}
