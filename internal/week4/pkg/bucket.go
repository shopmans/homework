package roll

type Bucket struct {
	WindowStart int64
	Count       int
}

func (b *Bucket) Add() {
	b.Count++
}
