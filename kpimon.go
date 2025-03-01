package main

import "example.com/kpimon/control"
import "fmt"
import "time"

func main() {
	time.Sleep(15* time.Second)
	fmt.Printf("In Main ------")
	_ = control.NewControl()
	fmt.Printf("c.Run()")
	// c.Run()
	/*
	err:=control.Test()
	if err != nil {
		fmt.Println(err)
	}
	*/
}
