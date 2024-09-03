package main

import "fmt"

func main() {
	var name string

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %s !\n", name)
}
