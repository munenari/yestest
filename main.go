package main

import (
	"fmt"
	"os"
)

func main() {
	yes := "y"
	if len(os.Args) > 1 {
		yes = os.Args[1]
	}
	for {
		fmt.Println(yes)
	}
}
