package main

import (
	// External Package
	fmt "fmt"

	// Internal Package
	common "projek/common"

	// Feature Home
	homeMenu "projek/features/home/menu"

	// Feature Post
	postStruct "projek/features/post/structs"

	// Feature Pasien
	pasien "projek/features/pasien"
	pasienStruct "projek/features/pasien/structs"
)

func main() {
	var arrPost postStruct.TabPost
	var arrPasien pasienStruct.TabPasien

	var input int
	homeMenu.ShowHomeMenu()

	for input != 4 {

		// Terima inputan dari user
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)

		// Reset console
		common.ResetConsole()

		// Cek apakah menu tersedia
		if input == 1 {
			// Pasien
			pasien.Main(&arrPasien, &arrPost)
			homeMenu.ShowHomeMenu()
		} else if input == 2 {
			// Dokter
			fmt.Println("Menu Dokter")
		} else if input == 3 {
			// Forum Konsultasi
			fmt.Println("Menu Konsultasi")
		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
	}

	// Exit Aplikasi
	fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Kami")
}
