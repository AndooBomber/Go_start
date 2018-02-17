package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var msg = flag.Bool("n", false, "setn")

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("入力エラー")
		os.Exit(1)
	}
	args := flag.Args()
	if *msg {
		is := 1
		for i := 0; i < len(args); i++ {
			fileName, err := os.Open(args[i])
			if err != nil {
				fmt.Println(err)
			}
			defer fileName.Close()
			scanner := bufio.NewScanner(fileName)
			for scanner.Scan() {
				fmt.Println(is, ":", scanner.Text())
				is++
			}
		}
	} else {
		for i := 0; i < len(args); i++ {
			fileName, err := os.Open(args[i])
			if err != nil {
				fmt.Println(err)
			}
			defer fileName.Close()
			scanner := bufio.NewScanner(fileName)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}
}
