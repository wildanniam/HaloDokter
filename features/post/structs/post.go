package post

import (
	pasienStruct "projek/features/pasien/structs"
)

const NMAX int = 20

type Post struct {
	ArrReply            [NMAX]Reply
	Nreply              int
	TxtAddPost, TagPost string
	ID                  string
	User                pasienStruct.Pasien
}

type TabPost struct {
	ArrPost [NMAX]Post
	N       int
}
