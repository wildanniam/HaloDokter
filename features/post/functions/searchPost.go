package post

import (
	postStruct "projek/features/post/structs"
)

func SearchPost(tpost postStruct.TabPost, x int) int {
	//Binary search
	//Mencari postingan berdasarkan ID post dan return indext dari post
	var med int
	var kr = 0
	var kn = tpost.N - 1
	var found = -1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if tpost.ArrPost[med].ID < x {
			kr = med + 1
		} else if tpost.ArrPost[med].ID > x {
			kn = med
		} else {
			found = med
		}
	}
	return found

}
