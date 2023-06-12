package post

import (
	"fmt"
	common "projek/common"
	postStruct "projek/features/post/structs"
)

func PoChat(tPost *postStruct.TabPost, UserType string) {
	//Pengguna menginputkan postingan kedalam aplikasi
	if tPost.N < postStruct.NMAX {
		if UserType == "pasien" {
			fmt.Print("Posting pertanyaan anda: ")
			common.InputMultipleString(&tPost.ArrPost[tPost.N].TxtAddPost)
			fmt.Print("Tag: ")
			fmt.Scan(&tPost.ArrPost[tPost.N].TagPost)
			tPost.ArrPost[tPost.N].ID = tPost.N
			tPost.N++
			fmt.Println("Pertanyaan anda telah diposting!")
		} else {
			fmt.Print("Posting informasi kesehatan: ")
			common.InputMultipleString(&tPost.ArrPost[tPost.N].TxtAddPost)
			fmt.Print("Tag: ")
			fmt.Scan(&tPost.ArrPost[tPost.N].TagPost)
			tPost.ArrPost[tPost.N].ID = tPost.N
			tPost.N++
			fmt.Println("Informasi anda telah diposting!")
		}

	}
}
