package post

import (
	postStruct "projek/features/post/structs"
)

func SearchPost(tpost postStruct.TabPost, x string) int {
	//Mencari postingan berdasarkan ID post dan return indext dari post
	var i int
	var found int = -1

	for i < tpost.N {
		if tpost.ArrPost[i].ID == x {
			found = i
		}
		i++
	}
	return found
}
