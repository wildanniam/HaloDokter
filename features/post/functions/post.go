package post

import (
	"fmt"
	common "projek/common"
	postStruct "projek/features/post/structs"
)

func PoChat(tPost *postStruct.TabPost) {
	//Pengguna menginputkan postingan kedalam aplikasi
	if tPost.N < postStruct.NMAX {
		fmt.Print("Posting pertanyaan anda: ")
		common.InputMultipleString(&tPost.ArrPost[tPost.N].TxtAddPost)
		fmt.Print("Tag: ")
		fmt.Scan(&tPost.ArrPost[tPost.N].TagPost)
		tPost.ArrPost[tPost.N].ID = common.GenerateRandomString(3)
		tPost.N++
		fmt.Println("Pertanyaan anda telah diposting!")
	}
}
