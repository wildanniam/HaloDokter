package pasien

import (
	"fmt"
	common "projek/common"
	pasienStruct "projek/features/pasien/structs"
)

func EditAcc(tPasien *pasienStruct.TabPasien) {
	var input int
	pasienIdx := LoginPasien(tPasien)

	if pasienIdx == -1 {
		common.ResetConsole()

		fmt.Println("=======================================================================================")
		fmt.Println("                  Mohon inputkan username dan password dengan benar!                             ")
		fmt.Println("=======================================================================================")
		fmt.Println()

		common.ShowEndAction()
	} else {
		common.ResetConsole()
		fmt.Println("Aksi")
		fmt.Println("1. Ganti Nama")
		fmt.Println("2. Ganti Username")
		fmt.Println("3. Ganti Password")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&input)
		if input == 1 {
			fmt.Print("Masukkan nama baru: ")
			common.InputMultipleString(&tPasien.ArrPasien[pasienIdx].Nama)
		} else if input == 2 {
			fmt.Print("Masukkan username baru: ")
			common.InputMultipleString(&tPasien.ArrPasien[pasienIdx].Username)
		} else if input == 3 {
			fmt.Print("Masukkan password baru: ")
			common.InputMultipleString(&tPasien.ArrPasien[pasienIdx].Password)
		} else if input == 4 {
			fmt.Println()
		}
	}
}
