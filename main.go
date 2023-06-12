package main

import (
	// External Package
	fmt "fmt"

	// Internal Package
	common "projek/common"

	// Feature Home
	homeMenu "projek/features/home/menu"

	// Feature Post
	doctorStruct "projek/features/dokter/structs"
	postFunc "projek/features/post/functions"
	postStruct "projek/features/post/structs"

	// Feature Pasien
	doctor "projek/features/dokter"
	pasien "projek/features/pasien"
	pasienStruct "projek/features/pasien/structs"

	//Umum
	menuUmum "projek/features/umum"
)

func main() {
	var arrPost postStruct.TabPost
	var arrPatient pasienStruct.TabPasien
	var arrDoctor doctorStruct.TabDokter

	var input int
	homeMenu.ShowHomeMenu()
	// Terima inputan dari user
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)

	for input != 4 {
		// Cek apakah menu tersedia
		// Pasien
		if input == 1 {
			// Reset console
			common.ResetConsole()
			pasien.Main(&arrPatient, &arrPost, &arrDoctor)
		} else if input == 2 {
			// Dokter
			common.ResetConsole()
			doctor.Main(&arrDoctor, &arrPost, &arrPatient)
		} else if input == 3 {
			// Forum Konsultasi
			common.ResetConsole()
			menuUmum.ShowUmumMenu()
			fmt.Print("Pilih Menu : ")
			fmt.Scan(&input)
			for input != 2 {
				if input == 1 {
					common.ResetConsole()
					postFunc.ShowPost(&arrPost, &arrDoctor, &arrPatient)
				} else {
					fmt.Println("Menu salah, coba lagi!")
				}
				menuUmum.ShowUmumMenu()
				fmt.Print("Pilih Menu : ")
				fmt.Scan(&input)
			}
		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
		homeMenu.ShowHomeMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}

	// Exit Aplikasi
	fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Kami")
}
