package main

import (
	"fmt"
	"github.com/lroc/gowi"
)

func main() {
	fmt.Println("init")

	app := gowi.New()

	w := NewMainWindow()
	w.SetWindowText("app example")
	w.Show()

	app.Run()
}
