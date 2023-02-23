package lit

import (
	"crypto/sha1"
	"encoding/hex"
	"path"
	"strconv"
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
		hasher := sha1.New()
		hasher.Write([]byte(blob.ToString()))
		tmp := hex.EncodeToString(hasher.Sum(nil))
		blob.oid = &tmp
	}
	return *blob.oid
}

func (blob *Blob) ToString() string {
	if blob.content == nil {
		tmp := blob.Type() + " " + strconv.Itoa(len(blob.Data)) + "\x00" + string(blob.Data)
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
