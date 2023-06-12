package post

const NMAX int = 20

type Post struct {
	ArrReply            [NMAX]Reply
	Nreply              int
	TxtAddPost, TagPost string
	ID                  int
	User                string
	UserIndex           int
}

type TabPost struct {
	ArrPost [NMAX]Post
	N       int
}

type Tag struct {
	TxtTag string
	N      int
}
