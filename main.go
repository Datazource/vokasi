package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan Umur Anda: ")
	fmt.Scan(&age)
	fmt.Printf("Hello %s !\n", name)
	fmt.Printf("Umur anda %s !\n", age)
}
