package auth

import (
	"fmt"
	common "projek/common"
	authMenu "projek/features/auth/functions/menu_auth"
	pasienFunc "projek/features/pasien/functions"
	pasienStruct "projek/features/pasien/structs"
)

func Main() {
	var input int
	var tPasien pasienStruct.TabPasien
	authMenu.ShowAuthMenu()

	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)
	common.ResetConsole()
	for input != 4 {
		if input == 1 {
			pasienFunc.ShowAuthMenu()
			fmt.Print("Pilih Menu : ")
			fmt.Scan(&input)
			for input != 3 {
				if input == 1 {
					pasienFunc.RegisPasien(&tPasien)
				} else if input == 2 {
					pasienFunc.LoginPasien(&tPasien)
				} else {
					fmt.Println("Menu Salah, coba lagi!")
				}
				fmt.Print("Pilih Menu : ")
				fmt.Scan(&input)
			}
			authMenu.ShowAuthMenu()
		} else if input == 2 {
			// Dokter
		} else if input == 3 {
			// Paisen non login
		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}
	fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Kami")

}
