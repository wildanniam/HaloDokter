package post

import (
	"fmt"
	postStruct "projek/features/post/structs"

	common "projek/common"
)

func RepPost(tPost *postStruct.TabPost, UserType string, UserIndex int) {
	//Pengguna menginputkan balasan pada postingan
	var IDPost int
	var idx int
	var inputUser string

	fmt.Println("Pilih ID postingan yang ingin dibalas: ")
	fmt.Scan(&IDPost)
	idx = SearchPost(*tPost, IDPost)

	if idx != -1 {
		fmt.Println("Balasan anda: ")
		common.InputMultipleString(&inputUser)
		if inputUser != " " && tPost.ArrPost[idx].Nreply < postStruct.NMAX {
			tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply].Message = inputUser
			tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply].User = UserType
			tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply].UserIndex = UserIndex
			tPost.ArrPost[idx].Nreply++
			fmt.Println("Balasan anda telah diposting!")
		}
	} else {
		fmt.Print("postingan tidak ditemukan!")
	}
}
