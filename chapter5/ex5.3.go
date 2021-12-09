package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
		_, _ = stdin.ReadString('\n')
		n, err = fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n, a, b)
		}
	} else {
		fmt.Println(n, a, b)
	}
}