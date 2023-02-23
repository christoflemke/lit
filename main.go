package main

import (
	"io/fs"
	"lit/lit"
	"log"
	"os"
	"path"
	"path/filepath"
)

var logger = log.New(os.Stderr, "", 0)

func usage(command string) {
	logger.Printf("lit: %s is not a lit command\n", command)
	os.Exit(1)
}

func main() {
	command := os.Args[1]
	switch command {
	case "init":
		initCommand()
	case "commit":
		commitCommand()
	default:
		usage(command)
	}
}

func resolveGitPath(d *string) string {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if d != nil {
		p = *d
	}
	rootPath, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	return path.Join(rootPath, ".git")
}

func commitCommand() {
	gitPath := resolveGitPath(nil)
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ws := lit.Workspace{RootPath: dir}
	db := lit.Database{DbPath: path.Join(gitPath, "objects")}
	var entries []lit.Entry
	for _, path := range ws.ListFiles() {
		data := ws.ReadFile(path)
		blob := lit.Blob{Data: data}
		var dbObject lit.DatabaseObject = &blob
		db.Store(&dbObject)
		entries = append(entries, lit.Entry{Path: path, Blob: &blob})
	}
	tree := lit.Tree{Entries: entries}
	var dbObject lit.DatabaseObject = &tree
	db.Store(&dbObject)
	logger.Println(ws.ListFiles())
}

func initCommand() {
	gitPath := resolveGitPath(nil)
	if len(os.Args) > 2 {
		gitPath = resolveGitPath(&os.Args[2])
	}
	dirs := [2]string{"objects", "refs"}
	for _, dir := range dirs {
		err := os.MkdirAll(path.Join(gitPath, dir), fs.FileMode(0755))
		if err != nil {
			panic(err)
		}
	}
	logger.Println("Initialized empty Lit repository in ", gitPath)
}
