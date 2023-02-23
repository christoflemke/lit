package lit

import (
	"path"
)

type Blob struct {
	Data    []byte
	oid     *string
	content *string
}

func (blob *Blob) Type() string {
	return "blob"
}

func (blob *Blob) Oid() string {
	if blob.oid == nil {
		var dbObject DatabaseObject = blob
		tmp := ComputeOid(&dbObject)
		blob.oid = &tmp
	}
	return *blob.oid
}

func (blob *Blob) ToString() string {
	if blob.content == nil {
		tmp := string(blob.Data)
		blob.content = &tmp
	}
	return *blob.content
}

func (blob *Blob) FilePath(dbPath string) string {
	oid := blob.Oid()
	return path.Join(dbPath, oid[0:2], oid[2:len(oid)])
}

func (blob *Blob) DirPath(dbPath string) string {
	oid := blob.Oid()
	return path.Join(dbPath, oid[0:2])
}
