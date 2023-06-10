package post

import (
	"fmt"
	postStruct "projek/features/post/structs"

	common "projek/common"
)

func RepPost(tPost *postStruct.TabPost, UserType string, UserIndex int) {
	//Pengguna menginputkan balasan pada postingan
	var IDPost string
	var idx int
	var inputUser string
	var response int

	fmt.Println("Pilih ID postingan yang ingin dibalas: ")
	fmt.Scan(&IDPost)
	idx = SearchPost(*tPost, IDPost)

	/*if idx != -1 {
		fmt.Println("Balasan anda: ")
		common.InputMultipleString(&inputUser)
		for tPost.ArrPost[idx].Nreply < len(tPost.ArrPost[idx].ArrReply) && inputUser != "STOP" {
			tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply] = inputUser
			tPost.ArrPost[idx].Nreply++
			fmt.Println("Balasan anda telah diposting!")

			fmt.Println("Balasan anda: ")
			common.InputMultipleString(&inputUser)
		}
	} else {
		fmt.Print("postingan tidak ditemukan!")
	}
	// for j := 0; j < tPost.ArrPost[idx].Nreply; j++ {
	// 	fmt.Println("Result : ", tPost.ArrPost[idx].ArrReply[j])
	// }*/

	if idx != -1 {
		fmt.Println("Balasan anda: ")
		common.InputMultipleString(&inputUser)
		fmt.Println("[1] Kirim          [2]Batal")
		fmt.Print("Action : ")
		fmt.Scan(&response)
		for response == 1 {
			if tPost.ArrPost[idx].Nreply < postStruct.NMAX {
				tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply].Message = inputUser
				tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply].User = UserType
				tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply].UserIndex = UserIndex
				tPost.ArrPost[idx].Nreply++
				fmt.Println("Balasan anda telah diposting!")
			}
			fmt.Println("Balasan anda: ")
			common.InputMultipleString(&inputUser)
			fmt.Println("[1] Balas          [2]Batal")
			fmt.Print("Action : ")
			fmt.Scan(&response)
		}
	} else {
		fmt.Print("postingan tidak ditemukan!")
	}
}
