package dokter

const NMAX int = 20

type Post struct {
	ArrReply            [NMAX]string
	TxtAddPost, TagPost string
	NoPost              int
}

type TabPost struct {
	ArrPost [NMAX]Post
	N       int
}
