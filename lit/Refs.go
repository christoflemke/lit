package lit

import (
	"os"
	"path"
)

type Refs struct {
	GitPath string
}

func (o *Refs) headPath() string {
	return path.Join(o.GitPath, "HEAD")
}

func (o *Refs) UpdateHead(oid string) {
	temp, err := os.CreateTemp(o.GitPath, "tmp-*")
	if err != nil {
		panic(err)
	}
	defer os.Remove(temp.Name())
	_, err = temp.Write([]byte(oid))
	if err != nil {
		panic(err)
	}
	err = os.Rename(temp.Name(), o.headPath())
	if err != nil {
		panic(err)
	}
}

func (o *Refs) ReadHead() *string {
	file, err := os.ReadFile(o.headPath())
	if err != nil {
		return nil
	}
	tmp := string(file)
	return &tmp
}
