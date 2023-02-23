package test

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"io"
	"lit/lit"
	"os"
	"os/exec"
	"path"
	"testing"
)

func TestObjectWrite(t *testing.T) {
	dbpath := os.TempDir()
	blob := lit.Blob{Data: []byte("hello\n")}

	if blob.Oid() != "ce013625030ba8dba906f756967f9e9ca394464a" {
		t.Errorf("blob.Oid() = %s, want: ce013625030ba8dba906f756967f9e9ca394464a", blob.Oid())
	}

	got := blob.DirPath(dbpath)
	if got != path.Join(dbpath, "ce") {
		t.Errorf("blob.DirPath(dbpath) = %s, want: ce", got)
	}

	got = blob.FilePath(dbpath)
	if got != path.Join(dbpath, "ce", "013625030ba8dba906f756967f9e9ca394464a") {
		t.Errorf("blob.FilePath(dbpath) = %s, want: ce/013625030ba8dba906f756967f9e9ca394464a", got)
	}

	database := lit.Database{DbPath: dbpath}
	var dbObject lit.DatabaseObject = &blob
	database.Store(&dbObject)
	bs, err := os.ReadFile(blob.FilePath(dbpath))
	if err != nil {
		panic(err)
	}
	expected := []byte("\x78\x9c\x4a\xca\xc9\x4f\x52\x30\x63\xc8\x48\xcd\xc9\xc9\xe7\x02\x04\x00\x00\xff\xff\x1d\xc5\x04\x14")
	if bytes.Compare(expected, bs) != 0 {
		t.Errorf("encoded bytes did not match, got %s", hex.EncodeToString(bs))
	}
}

func TestTree(t *testing.T) {
	helloBlob := lit.Blob{Data: []byte("hello\n")}
	worldBlob := lit.Blob{Data: []byte("world\n")}
	helloEntry := lit.Entry{Blob: &helloBlob, Path: "hello.txt"}
	worldEntry := lit.Entry{Blob: &worldBlob, Path: "world.txt"}
	tree := lit.Tree{Entries: []lit.Entry{helloEntry, worldEntry}}
	println(tree.ToString())
}

func compareFileBytes(t *testing.T, relativePath string) {
	referenceFilePath := path.Join("reference-repo", ".git", "objects", relativePath)
	referenceBytes, err := readAndInflateFile(t, referenceFilePath)
	if err != nil {
		t.Errorf("Unable to read file: %s", referenceFilePath)
		return
	}
	myFilePath := path.Join("my-repo", ".git", "objects", relativePath)
	myBytes, err := readAndInflateFile(t, myFilePath)

	if bytes.Compare(myBytes, referenceBytes) != 0 {
		t.Errorf("Bytes did not match\nfile: %s\nexpected: %s\nactual:   %s", myFilePath, hex.EncodeToString(referenceBytes), hex.EncodeToString(myBytes))
		return
	}
}

func readAndInflateFile(t *testing.T, referenceFilePath string) ([]byte, error) {
	referenceFile, err := os.Open(referenceFilePath)
	if err != nil {
		t.Errorf("Unable to open reference file: %s", referenceFilePath)
		return nil, err
	}
	referenceReader, err := zlib.NewReader(referenceFile)
	if err != nil {
		t.Errorf("Unable to create reader for file: %s", referenceFilePath)
		return nil, err
	}
	referenceBytes, err := io.ReadAll(referenceReader)
	return referenceBytes, err
}

func TestCompareWithReference(t *testing.T) {
	command := exec.Command("./setup-test-data.sh")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Errorf("Output: %s", string(output))
		t.Error(err)
	}
	compareFileBytes(t, "cc/628ccd10742baea8241c5924df992b5c019f71")
	compareFileBytes(t, "ce/013625030ba8dba906f756967f9e9ca394464a")
	compareFileBytes(t, "88/e38705fdbd3608cddbe904b67c731f3234c45b")
}
