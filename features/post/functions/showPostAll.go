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
	common.ResetConsole()

	for i := 0; i < arrPost.N; i++ {
		fmt.Printf("ID Postingan :%s\nTag: %s \n%s \n", arrPost.ArrPost[i].ID, arrPost.ArrPost[i].TagPost, arrPost.ArrPost[i].TxtAddPost)
		for j := 0; j < arrPost.ArrPost[i].Nreply; j++ {
			// Pengecekan tipe user
			if arrPost.ArrPost[i].ArrReply[j].User == "dokter" {

				doctor := arrDoctor.ArrDokter[arrPost.ArrPost[i].ArrReply[j].UserIndex]
				fmt.Println("==== Balasan =====")
				fmt.Println("Nama Dokter: ", doctor.Nama)
				fmt.Println("Spesialis: ", doctor.Spesialis)

			} else if arrPost.ArrPost[i].ArrReply[j].User == "pasien" {

				patient := arrPatient.ArrPasien[arrPost.ArrPost[i].ArrReply[j].UserIndex]
				fmt.Println("==== Balasan =====")
				fmt.Println("Nama Pasien: ", patient.Nama)

			}
			fmt.Println("Balasan : ", arrPost.ArrPost[i].ArrReply[j].Message)
		}
		fmt.Println()

	}

}
