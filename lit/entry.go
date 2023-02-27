package lit

type Entry struct {
	Blob *Blob
	Path string
}

func (e *Entry) Oid() string {
	return e.Blob.ToDbObject().Oid()
}
