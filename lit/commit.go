package lit

import "fmt"

type Commit struct {
	Tree    *Tree
	Author  *Author
	Message string
	oid     *string
}

func (commit *Commit) Oid() string {
	if commit.oid == nil {
		var dbObject DatabaseObject = commit
		tmp := ComputeOid(&dbObject)
		commit.oid = &tmp
	}
	return *commit.oid
}

func (commit *Commit) Type() string {
	return "commit"
}

func (commit *Commit) ToString() string {
	return fmt.Sprintf(
		"tree %s\nauthor %s\ncommitter %s\n%s",
		commit.Tree.Oid(),
		commit.Author.ToString(),
		commit.Author.ToString(),
		commit.Message,
	)
}
