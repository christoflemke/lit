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
	err := os.WriteFile(o.headPath(), []byte(oid), 0644)
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
