package consumer

type Article struct {
	Title string
}

// Article will implement Stringer now
func (a *Article) String() string {
	return a.Title
}
