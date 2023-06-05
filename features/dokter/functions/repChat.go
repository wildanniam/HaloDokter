package dokter

import (
	"fmt"
	postStruct "projek/features/post/structs"
)

func repChat(tPost *postStruct.TabPost) {
	//Pengguna menginputkan balasan pada postingan
	var noPost int
	fmt.Println("Pilih postingan yang ingin dibalas: ")
	fmt.Scan(&noPost)
	fmt.Println("Balasan anda: ")

	// common.InputMultipleString(&tPost.ArrPost[noPost-1].ArrReply[len(tPost.ArrPost[noPost-1].ArrReply)])
	fmt.Println("Balasan anda telah diposting!")
}

// ArrReply                      [NMAX]string
// 	TxtAddPost, TagPost string
// 	NoPost                        int
// 	User

/**
-> ArrPost {
	Post []
	N = 1
}

arrPost.N ++
arrPost.[0].Id = "5JfbW"
arrPost.[0].TxtAddPost = "Hslo Gaes"
arrPost.[0].TagPost = "#sakitperut"
arrPost.[0].User = {
	name = "INdra Mahesa"
}
arrPost.[0].arrReply = [
	{
		tag = "#sakitperut"
		txt = "Wahh bermfaat sekali"
	},
	{
		tag = "#sakitperut"
		txt = "Wahh bermfaat sekali"
	}
	{
		tag = "#sakitperut"
		txt = "Wahh bermfaat sekali"
	}
]

getPostById(id) -> post

2 panjang
post.arrReply[len(arrReply)] = {
	tag = "#sakitperut"
	txt = "Wahh bermfaat sekali"
}

/**
 Post
 - TxtAddPost
 - TagPost
 - User {
	- Id
	- Nama
	- username
	-...
 }
 - ArrReply [
	{
		- tag
		- TxtAddPost
	},
	{
		- tag
		- TxtAddPost
	},
 ]
**/
