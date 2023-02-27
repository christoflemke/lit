package lit

import "fmt"

type Commit struct {
	Tree    *Tree
	Author  *Author
	Message string
}

func (commit *Commit) ToDbObject() *DbObject {
	return &DbObject{StringRepresentation: commit.ToString(), StorageType: "commit"}
}

func (commit *Commit) ToString() string {
	return fmt.Sprintf(
		"tree %s\nauthor %s\ncommitter %s\n%s",
		commit.Tree.ToDbObject().Oid(),
		commit.Author.ToString(),
		commit.Author.ToString(),
		commit.Message,
	)
}
