package post

import (
	"fmt"
	"projek/common"
	doctorMenu "projek/features/dokter/menu"
	doctorStruct "projek/features/dokter/structs"
	patientFunc "projek/features/pasien/functions"
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
		var searchTag string
		var input int
		patienMenu.ShowHomePasienMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
		for input != 7 {
			postFunc.ShowPost(arrPost, arrDoctor, arrPatient)
			if input == 1 {
				postFunc.PoChat(arrPost, "pasien")
				common.ResetConsole()
			} else if input == 2 {
				postFunc.RepPost(arrPost, UserType, UserIndex)
				common.ResetConsole()
			} else if input == 3 {
				postFunc.ShowPost(arrPost, arrDoctor, arrPatient)
			} else if input == 4 {
				fmt.Print("Masukkan tag yang ingin anda cari: ")
				fmt.Scan(&searchTag)
				postFunc.SearchByTag(arrPost, arrPatient, arrDoctor, searchTag)
			} else if input == 5 {
				patienMenu.ShowDeleteAccount()
				patientFunc.EditAcc(arrPatient)
			} else if input == 6 {
				patienMenu.ShowDeleteAccount()
				patientFunc.DeleteAccount(arrPatient)

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
