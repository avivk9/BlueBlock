package main

import "fmt"

func main() {
	fmt.Println("Welcome to BlueBlocks by avivk9")

	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()

}
