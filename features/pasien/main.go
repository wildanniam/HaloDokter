package pasien

import (
	//
	"fmt"
	common "projek/common"

	doctorStruct "projek/features/dokter/structs"

	//
	pasienFunc "projek/features/pasien/functions"
	pasienMenu "projek/features/pasien/menu"
	pasienStruct "projek/features/pasien/structs"

	//

	post "projek/features/post"
	postStruct "projek/features/post/structs"
)

func Main(arrPasien *pasienStruct.TabPasien, arrPost *postStruct.TabPost, arrDoctor *doctorStruct.TabDokter) {
	var input int
	pasienMenu.ShowAuthPasienMenu()
	// Terima inputan dari user
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)

	for input != 3 {

		if input == 1 {
			// Register
			common.ResetConsole()
			pasienFunc.RegistrasiPasien(arrPasien)
			common.ResetConsole()
		} else if input == 2 {
			// Login
			PatientIndex := pasienFunc.LoginPasien(arrPasien)
			if PatientIndex != -1 {
				// Jika login berhasil
				common.ResetConsole()

				post.Main(arrPost, arrDoctor, arrPasien, "pasien", PatientIndex)
			} else {
				// Jika login gagal
				common.ResetConsole()

				fmt.Println("=======================================================================================")
				fmt.Println("                  Mohon inputkan username dan password dengan benar!                             ")
				fmt.Println("=======================================================================================")
				fmt.Println()

				common.ShowEndAction()
			}
		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
		pasienMenu.ShowAuthPasienMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}

	// Reset console
	common.ResetConsole()
}
