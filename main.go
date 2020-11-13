package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//exit process immediately upon sigterm
	handleSigTerms()

	fmt.Println("hello")
}

func handleSigTerms() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
