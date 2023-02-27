package lit

import (
	"encoding/hex"
	"sort"
)

const MODE = "100644"

type Tree struct {
	Entries []Entry
}

func (tree *Tree) ToDbObject() *DbObject {
	return &DbObject{StringRepresentation: tree.ToString(), StorageType: "tree"}
}

type byPath []Entry

func (s byPath) Len() int {
	return len(s)
}
func (s byPath) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPath) Less(i, j int) bool {
	return s[i].Path < s[j].Path
}

func (tree *Tree) ToString() string {
	sort.Sort(byPath(tree.Entries))
	result := ""
	for _, e := range tree.Entries {
		oidBytes, err := hex.DecodeString(e.Oid())
		if err != nil {
			panic(err)
		}
		oidString := string(oidBytes)
		result += MODE + " " + e.Path + "\x00" + oidString
	}
	return result
}
