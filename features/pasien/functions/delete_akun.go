package pasien

import (
	"fmt"
	common "projek/common"
	patientStruct "projek/features/pasien/structs"
)

func DeleteAccount(arrPatient *patientStruct.TabPasien) {
	var input int
	i := LoginPasien(arrPatient)
	if i != -1 {
		fmt.Println("===================================================================")
		fmt.Println("        Aksi ini akan menghapus data anda secara permanen          ")
		fmt.Println("                         Tetap lanjutkan?                          ")
		fmt.Println("===================================================================")
		fmt.Println("[1] Lanjut             [2]Batal                                    ")
		fmt.Scan(&input)
		if input == 1 {
			for i < arrPatient.N-1 {
				arrPatient.ArrPasien[i] = arrPatient.ArrPasien[i+1]
				i++
			}
			arrPatient.N--
		} else if input == 2 {
			fmt.Println()
		}
	} else {
		common.ResetConsole()

		fmt.Println("=======================================================================================")
		fmt.Println("                  Mohon inputkan username dan password dengan benar!                             ")
		fmt.Println("=======================================================================================")
		fmt.Println()

		common.ShowEndAction()
	}

}
