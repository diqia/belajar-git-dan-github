package main

import "fmt"

func inputOutput() { // menambah fungsi input output
	fmt.Println("--------------")
	var umur int
	fmt.Print("Berapa umur anda? ")
	fmt.Scan(&umur)
	fmt.Print("Umur anda adalah ", umur)
}

func main() {
	fmt.Print("Hello, World!\n")
	fmt.Print("Nama saya Tafif\n")
	fmt.Print("Saya Mahasiswa Informatika Angkatan 2021\n")
	fmt.Print("Disini saya akan belajar mengenai Git dan Github\n")
	inputOutput() // Memanggil fungsi input output
}