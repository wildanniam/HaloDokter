package post

import (
	"fmt"
	"projek/common"
	doctorStruct "projek/features/dokter/structs"
	patientStruct "projek/features/pasien/structs"
	postStruct "projek/features/post/structs"
)

func ShowPost(arrPost *postStruct.TabPost, arrDoctor *doctorStruct.TabDokter, arrPatient *patientStruct.TabPasien) {
	//Menampilkan isi postingan yang telah diinputkan pengguna
	var i int
	fmt.Println()
	for i = 0; i < arrPost.N; i++ {
		fmt.Printf("ID: %d\n%s \nTag: %s \n",
			arrPost.ArrPost[i].ID, arrPost.ArrPost[i].TxtAddPost,
			arrPost.ArrPost[i].TagPost,
		)

		for j := 0; j < arrPost.ArrPost[i].Nreply; j++ {
			// Pengecekan tipe user
			if arrPost.ArrPost[i].ArrReply[j].User == "dokter" {

				doctor := arrDoctor.ArrDokter[arrPost.ArrPost[i].ArrReply[j].UserIndex]
				fmt.Println("==== Balasan =====")
				fmt.Println("dr.", doctor.Nama)
				fmt.Println("Spesialis", doctor.Spesialis)

			} else if arrPost.ArrPost[i].ArrReply[j].User == "pasien" {

				patient := arrPatient.ArrPasien[arrPost.ArrPost[i].ArrReply[j].UserIndex]
				fmt.Println("==== Balasan =====")
				fmt.Println(patient.Nama)

			}
			fmt.Println("Balasan : ", arrPost.ArrPost[i].ArrReply[j].Message)
		}
		common.ResetConsole()
	}

}
