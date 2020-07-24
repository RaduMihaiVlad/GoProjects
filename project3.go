package main

import (
	"fmt"
	"bufio"
	"os"
)

func WaitForInput(chIn chan string, chOut chan int) {
	
	for true {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		
		if (text[:(len(text) - 2)] != "exit") {
			chIn <- text
		} else {
			chOut <- 1
			return
		}
	}
}

func Write(chIn chan string, chOut chan int) {

	for true {
		select {
		case x:= <- chIn:
			fmt.Println(x)
		case <- chOut:
			return
		}
	}

}

/*
	I am creating two theads, except the main one. The first one listen for the user's input, and post 
	it into a channel. The second one is listening for the first channel and prints what the it returns.
	When the user write "exit" command, the first thread closes and is announcing the main thread to close.
	Then the main thread closes the Write thread and also closes itself.
*/

func main() {

	chInput := make(chan string)
	chOutWrite := make(chan int)
	chOutMain := make(chan int)

	go WaitForInput(chInput, chOutMain)
	go Write(chInput, chOutWrite)

	<- chOutMain
	chOutWrite <- 1
	
}