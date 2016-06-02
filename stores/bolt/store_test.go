package bolt

import (
	"bytes"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/kildevaeld/go-filestore"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	dbPath := "test.db"
	defer os.Remove(dbPath)
	fs, err := New(dbPath)

	if err != nil {
		t.Error(err)
	}

	fileId := filestore.NewFid()

	file := filestore.File{
		Fid:      fileId,
		Path:     "/",
		Filename: "readme.md",
		Mime:     "text/markdown",
	}

	reader := bytes.NewBuffer([]byte("Hello, World!"))

	l, e := fs.Create(reader, &file)

	if e != nil {
		t.Error(e)
	}

	if l != uint64(13) {
		t.Errorf("Size must be 13, was: %d\n", l)
	}

	fss := fs.(*fs_impl)
	fss.Close()

	fss.bolt.View(func(tx *bolt.Tx) error {

		meta := tx.Bucket(metaBucket)

		b := meta.Get([]byte("/readme.md"))

		file.Unmarshal(b)

		assert.Equal(t, file.Fid, fileId)
		assert.Equal(t, file.Path, "/")
		assert.Equal(t, file.Filename, "readme.md")
		assert.Equal(t, file.Filesize, uint64(13))
		assert.Equal(t, file.Mime, "text/markdown")
		return nil
	})

}
