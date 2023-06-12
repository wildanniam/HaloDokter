package dokter

import (
	"fmt"

	common "projek/common"
	// struct dokter
	dokterFunc "projek/features/dokter/functions"
	inputDoctor "projek/features/dokter/inputDoctor"
	dokterMenu "projek/features/dokter/menu"
	dokterStruct "projek/features/dokter/structs"
	"projek/features/post"

	// Struct patient
	patientStruct "projek/features/pasien/structs"

	//struct post
	//post "projek/features/post"
	postStruct "projek/features/post/structs"
)

func Main(
	arrDokter *dokterStruct.TabDokter,
	arrPost *postStruct.TabPost,
	arrPatient *patientStruct.TabPasien,
) {
	var input int

	dokterMenu.ShowAuthDokterMenu()
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)

	for input != 2 {
		if input == 1 {
			//memasukkan data data dokter kedalam array
			inputDoctor.InputDoctor(arrDokter)
			//memasukkan index dokter kedalam variable
			doctorIndex := dokterFunc.LoginDokter(arrDokter)
			if doctorIndex != -1 {
				common.ResetConsole()
				post.Main(arrPost, arrDokter, arrPatient, "dokter", doctorIndex)
			} else {
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
		dokterMenu.ShowAuthDokterMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}
	common.ResetConsole()
}
