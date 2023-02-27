package lit

import (
	"compress/zlib"
	"errors"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	DbPath string
}

type Storable interface {
	ToDbObject() *DbObject
}

func (d *Database) Store(storable Storable) {
	blob := storable.ToDbObject()
	objectPath := blob.FilePath(d.DbPath)
	if _, err := os.Stat(objectPath); !errors.Is(err, os.ErrNotExist) {
		return
	}
	objectFolder := blob.DirPath(d.DbPath)
	err := d.mkObjectFolder(objectFolder)
	temp := d.newTempFile(objectFolder)
	defer os.Remove(temp.Name())
	writer := zlib.NewWriter(temp)

	blobString := (*blob).StringRepresentation
	data := (*blob).StorageType + " " + strconv.Itoa(len(blobString)) + "\x00" + blobString
	_, err = io.Copy(writer, strings.NewReader(data))
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
