package main

import (
	"flag"
	"fmt"
	"github.com/quminzhi/rail"
)

var (
	option  string // encryption or decryption
	message string // plaintext or ciphertext
	depth   int    // depth of rail cipher
	repeat  int    // how many times to repeat
	help    bool   // show help
)

func main() {
	flag.StringVar(&option, "option", "encryption", "encryption or decryption")
	flag.StringVar(&message, "message", "hello world", "input message")
	flag.IntVar(&depth, "depth", 2, "depth of rail fence encryption")
	flag.IntVar(&repeat, "repeat", 1, "the number of times to repeat encryption or decryption")
	flag.BoolVar(&help, "help", false, "display usage information")

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	switch option {
	case "encryption":
		ciphertext := rail.Encode(message, []int{depth, repeat})
		fmt.Printf("Ciphertext: %v\n", ciphertext)
	case "decryption":
		plaintext := rail.Decode(message, []int{depth, repeat})
		fmt.Printf("Plaintext: %v\n", plaintext)
	default:
		fmt.Println("Unknown option: encryption or decryption")
		return
	}

	return
}
