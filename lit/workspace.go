package lit

import (
	"os"
)

var IGNORE = []string{".", "..", ".git"}

type Workspace struct {
	RootPath string
}

func (w *Workspace) ListFiles() []string {
	entries, err := os.ReadDir(w.RootPath)
	if err != nil {
		panic(err)
	}
	allFiles := make([]string, len(entries))
	for i, e := range entries {
		allFiles[i] = e.Name()
	}
	var filtered []string
	for _, name := range allFiles {
		ignored := false
		for _, ignoredName := range IGNORE {
			if ignoredName == name {
				ignored = true
				break
			}
		}
		if !ignored {
			filtered = append(filtered, name)
		}
	}
	return filtered
}

func (w *Workspace) ReadFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}
