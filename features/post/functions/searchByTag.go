package post

import (
	"fmt"
	common "projek/common"
	doctorStruct "projek/features/dokter/structs"
	patientStruct "projek/features/pasien/structs"
	postStruct "projek/features/post/structs"
)

func SearchByTag(arrPost *postStruct.TabPost, arrPatient *patientStruct.TabPasien, arrDoctor *doctorStruct.TabDokter, X string) {
	var tabFound postStruct.TabPost
	var j int = 0
	var i int
	fmt.Println()
	for j < arrPost.N {
		if arrPost.ArrPost[j].TagPost == X {
			//Show Postingan
			tabFound.ArrPost[i] = arrPost.ArrPost[j]
			fmt.Println("ID: ", tabFound.ArrPost[i].ID)
			fmt.Println(tabFound.ArrPost[i].TxtAddPost)
			fmt.Println("Tag : ", tabFound.ArrPost[i].TagPost)
			// Show Balasan
			for k := 0; k < tabFound.ArrPost[i].Nreply; k++ {
				if arrPost.ArrPost[i].ArrReply[k].User == "dokter" {

					doctor := arrDoctor.ArrDokter[arrPost.ArrPost[i].ArrReply[k].UserIndex]
					fmt.Println("==== Balasan =====")
					fmt.Println("dr.", doctor.Nama)
					fmt.Println("Spesialis", doctor.Spesialis)

				} else if arrPost.ArrPost[i].ArrReply[k].User == "pasien" {

					patient := arrPatient.ArrPasien[arrPost.ArrPost[i].ArrReply[k].UserIndex]
					fmt.Println("==== Balasan =====")
					fmt.Println(patient.Nama)

				}
				fmt.Println("Balasan : ", arrPost.ArrPost[i].ArrReply[k].Message)
			}
			i++
		}
		j++
	}
	common.ResetConsole()
}
