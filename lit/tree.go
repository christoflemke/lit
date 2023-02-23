package lit

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strconv"
)

const MODE = "100644"

type Tree struct {
	Entries []Entry
	oid     *string
}

func (tree *Tree) Type() string {
	return "tree"
}
func (tree *Tree) Oid() string {
	if tree.oid == nil {
		hasher := sha1.New()
		hasher.Write([]byte(tree.ToString()))
		tmp := hex.EncodeToString(hasher.Sum(nil))
		tree.oid = &tmp
	}
	return *tree.oid
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
		oidByts, err := hex.DecodeString(e.Oid())
		if err != nil {
			panic(err)
		}
		oidString := string(oidByts)
		result += MODE + " " + e.Path + "\x00" + oidString
	}
	result = tree.Type() + " " + strconv.Itoa(len(result)) + "\x00" + result
	return result
}
