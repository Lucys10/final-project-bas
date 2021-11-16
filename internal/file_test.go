package internal

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

type testDir struct {
	dir1     string
	dir2     string
	fileName string
}

func TestProcessDirChek(t *testing.T) {
	dir1 := "/Users/lokis/Documents/RebreinGO/sync_directory/internal/test1/"
	log := logrus.New()
	req := require.New(t)
	res := readDir(dir1, log)
	req.NotNil(res)
}

func TestCopyFile(t *testing.T) {
	req := require.New(t)

	goodPath := testDir{
		dir1:     "/Users/lokis/Documents/RebreinGO/sync_directory/internal/test1/",
		dir2:     "/Users/lokis/Documents/RebreinGO/sync_directory/internal/test2/",
		fileName: "test1.txt"}

	badPath := testDir{
		dir1:     "sync_directory/internal/test1/",
		dir2:     "sync_directory/internal/test1/",
		fileName: "tes.txt",
	}

	t.Run("GoodPath", func(t *testing.T) {
		res, err := copyFile(goodPath.dir1, goodPath.dir2, goodPath.fileName)
		req.NoError(err)
		req.NotNil(res)
	})

	t.Run("BadPath", func(t *testing.T) {
		res, err := copyFile(badPath.dir1, badPath.dir2, badPath.fileName)
		req.Error(err)
		req.Equal(int64(0), res)
	})
}
