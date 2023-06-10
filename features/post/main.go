package post

import (
	"fmt"
	"projek/common"
	doctorStruct "projek/features/dokter/structs"
	pasienMenu "projek/features/pasien/menu"
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
	var input int
	pasienMenu.ShowHomePasienMenu(arrPost)
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)
	for input != 3 {
		postFunc.ShowPost(arrPost, arrDoctor, arrPatient)
		if input == 1 {
			postFunc.PoChat(arrPost)
		} else if input == 2 {
			postFunc.RepPost(arrPost, UserType, UserIndex)
		} else if input == 3 {
		}
		common.ResetConsole()
		pasienMenu.ShowHomePasienMenu(arrPost)
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}
}
