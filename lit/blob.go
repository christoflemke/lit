package lit

type Blob struct {
	Data []byte
}

func (blob *Blob) ToDbObject() *DbObject {
	return &DbObject{StringRepresentation: blob.ToString(), StorageType: "blob"}
}

func (blob *Blob) ToString() string {
	return string(blob.Data)
}
