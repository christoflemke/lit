package lit

import (
	"strings"
)

type Commit struct {
	Tree    *Tree
	Author  *Author
	Message string
	Parent  *string
}

func (commit *Commit) ToDbObject() *DbObject {
	return &DbObject{StringRepresentation: commit.ToString(), StorageType: "commit"}
}

func (commit *Commit) ToString() string {
	var result []string
	result = append(result, "tree "+commit.Tree.ToDbObject().Oid())
	if commit.Parent != nil {
		result = append(result, "parent "+*commit.Parent)
	}
	result = append(result, "author "+commit.Author.ToString())
	result = append(result, "committer "+commit.Author.ToString())
	return strings.Join(result, "\n")
}
