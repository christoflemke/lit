package lit

import (
	"compress/zlib"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

type Database struct {
	DbPath string
}

type DatabaseObject interface {
	Oid() string
	ToString() string
	Type() string
}

func FilePath(dbObject *DatabaseObject, dbPath string) string {
	oid := (*dbObject).Oid()
	return path.Join(dbPath, oid[0:2], oid[2:len(oid)])
}

func DirPath(dbObject *DatabaseObject, dbPath string) string {
	oid := (*dbObject).Oid()
	return path.Join(dbPath, oid[0:2])
}

func (d *Database) Store(blob *DatabaseObject) {
	objectPath := FilePath(blob, d.DbPath)
	objectFolder := DirPath(blob, d.DbPath)
	err := d.mkObjectFolder(objectFolder)
	temp := d.newTempFile(objectFolder)
	defer os.Remove(temp.Name())
	writer := zlib.NewWriter(temp)

	_, err = io.Copy(writer, strings.NewReader((*blob).ToString()))
	if err != nil {
		panic(err)
	}
	writer.Close()
	temp.Close()
	err = os.Rename(temp.Name(), objectPath)
	if err != nil {
		panic(err)
	}
}

func (d *Database) mkObjectFolder(objectFolder string) error {
	err := os.MkdirAll(objectFolder, fs.FileMode(0755))
	if err != nil {
		panic(err)
	}
	return err
}

func (d *Database) newTempFile(objectFolder string) *os.File {
	temp, err := os.CreateTemp(objectFolder, "tmp-*")
	if err != nil {
		panic(err)
	}
	return temp
}