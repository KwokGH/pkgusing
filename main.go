package main

import "fmt"

func main() {
	//watermilluse.RunSub()
	//watermilluse.RunPub()
	test()
}

func test(){
	defer func(){
		fmt.Println("hello")
	}()
}
