package wrongEmbedding

import "sync"

type InMem struct {
	sync.Mutex
	m map[string]int
	M map[string]int
}

func New() *InMem {
	return &InMem{m: make(map[string]int)}
}
