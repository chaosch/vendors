package localdb

type BatItem struct {
	Op    string
	Key   []byte
	Value []byte
	Ttl   int
}
