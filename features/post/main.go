package post

import (
	"fmt"
	"projek/common"
	doctorMenu "projek/features/dokter/menu"
	doctorStruct "projek/features/dokter/structs"
	patienMenu "projek/features/pasien/menu"
	patientStruct "projek/features/pasien/structs"
	postFunc "projek/features/post/functions"
	postStruct "projek/features/post/structs"
)

func Main(
	arrPost *postStruct.TabPost,
	arrDoctor *doctorStruct.TabDokter,
	arrPatient *patientStruct.TabPasien,
	UserType string,
	UserIndex int,
) {
	if UserType == "pasien" {
		var input int
		patienMenu.ShowHomePasienMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
		for input != 4 {
			postFunc.ShowPost(arrPost, arrDoctor, arrPatient)
			if input == 1 {
				postFunc.PoChat(arrPost, "pasien")
				common.ResetConsole()
			} else if input == 2 {
				postFunc.RepPost(arrPost, UserType, UserIndex)
				common.ResetConsole()
			} else if input == 3 {
				postFunc.ShowPost(arrPost, arrDoctor, arrPatient)
			}
			patienMenu.ShowHomePasienMenu()
			fmt.Print("Pilih Menu : ")
			fmt.Scan(&input)
		}
	} else if UserType == "dokter" {
		var input int
		doctorMenu.ShowHomeDokterMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
		//Kalau input == 3 akan keluar
		for input != 4 {
			postFunc.ShowPost(arrPost, arrDoctor, arrPatient)
			if input == 1 {
				postFunc.PoChat(arrPost, "dokter")
			} else if input == 2 {
				postFunc.RepPost(arrPost, UserType, UserIndex)
			}
			common.ResetConsole()
			doctorMenu.ShowHomeDokterMenu()
			fmt.Print("Pilih Menu : ")
			fmt.Scan(&input)
		}
	}

}
