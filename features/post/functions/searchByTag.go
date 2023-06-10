package post

import (
	"fmt"
	structTag "projek/features/post/structs"
)

func SearchByTag(tag *structTag.TabPost, X string) int {
	var found int = -1
	var j int = 0
	for j < tag.N && found == -1 {
		if tag.ArrPost[j].TagPost == X {
			found = j
		}
		j++
	}
	return found
}

func ShowResultTag(tag *structTag.TabPost, Xin string) {
	// var found int = SearchByTag(tag, Xin)
	fmt.Scan(&Xin)
	// for i < tag.N {
	// 	if found != -1
	// }

}
