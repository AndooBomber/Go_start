package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"io"
)

func quiz(s []string, n int) string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(n)
	return s[r]
}

func judge(r io.Reader) <-chan bool {
	subjects := []string{"animal", "cat", "dog", "egg", "orange", "bike", "car", "bus", "train", "pen", "notebook", "chair", "desk", "ball", "box", "cap", "bag", "picture", "camera", "computer", "TV", "guitar", "piano", "door", "window", "bed", "kitchen", "tree", "house", "home"}
	s := quiz(subjects, 30)
	fmt.Println(s)
	ch := make(chan bool)
	go func() {
		for {
			scan := bufio.NewScanner(r)
			answer := false
			fmt.Print(">")
			for scan.Scan() {
				if scan.Text() == s {
					answer = true
				}
				break
			}
			if answer {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	count := 0
	timer := time.After(10 * time.Second)
	for {
		ch := judge(os.Stdin)
		select {
		case <-timer:
			fmt.Println("\nScore:", count,"point")
			os.Exit(1)
		case <-ch:
			count++
		}
	}
}
