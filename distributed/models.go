package distributed

// VectorModel is used by server.go and matrix.go as a
// contract between client and server.
type VectorModel struct {
	RowIndex int
	Vector   []int
}
