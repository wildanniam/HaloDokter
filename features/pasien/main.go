package pasien

import (
	//
	"fmt"
	"projek/common"

	//
	pasienFunc "projek/features/pasien/functions"
	pasienMenu "projek/features/pasien/menu"
	pasienStruct "projek/features/pasien/structs"

	//
	postStruct "projek/features/post/structs"
)

func Main(arrPasien *pasienStruct.TabPasien, arrPost *postStruct.TabPost) {
	var input int
	pasienMenu.ShowAuthPasienMenu()

	for input != 3 {

		// Terima inputan dari user
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)

		// Reset console
		common.ResetConsole()

		// Cek apakah menu tersedia
		if input == 1 {

			// Register
			pasienFunc.RegistrasiPasien(arrPasien)
			pasienMenu.ShowAuthPasienMenu()

		} else if input == 2 {

			// Login
			if pasienFunc.LoginPasien(arrPasien) {
				// Jika login berhasil
				common.ResetConsole()

				pasienMenu.ShowHomePasienMenu()
			} else {
				// Jika login gagal
				common.ResetConsole()

				fmt.Println("=======================================================================================")
				fmt.Println("                  Mohon inputkan username dan password dengan benar!                             ")
				fmt.Println("=======================================================================================")
				fmt.Println()

				common.ShowEndAction()
			}

			pasienMenu.ShowAuthPasienMenu()

		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
	}

	// Reset console
	common.ResetConsole()
}
