package main

import "fmt"
import "github.com/PMoneda/dianadb/compress"

func main() {
	fmt.Println("Hello BoletoApi")
	compress.Zip([]byte("Ola"))

}
