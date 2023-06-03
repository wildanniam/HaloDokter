package pasien

const NMAX int = 20

type Post struct {
	TxtAddPost, TxtReply, TagPost string
	NoPost                        int
}

type TabPost struct {
	ArrPost [NMAX]Post
	N       int
}
