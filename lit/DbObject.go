package lit

import (
	"crypto/sha1"
	"encoding/hex"
	"path"
	"strconv"
)

type DbObject struct {
	StringRepresentation string
	StorageType          string
	oid                  *string
}

func (o *DbObject) Oid() string {
	if o.oid == nil {
		hasher := sha1.New()
		data := o.StorageType + " " + strconv.Itoa(len(o.StringRepresentation)) + "\x00" + o.StringRepresentation
		hasher.Write([]byte(data))
		tmp := hex.EncodeToString(hasher.Sum(nil))
		o.oid = &tmp
	}
	return *o.oid
}

func (o *DbObject) FilePath(dbPath string) string {
	oid := (*o).Oid()
	return path.Join(dbPath, oid[0:2], oid[2:len(oid)])
}

func (o *DbObject) DirPath(dbPath string) string {
	oid := (*o).Oid()
	return path.Join(dbPath, oid[0:2])
}
