package internal

import "testing"

func BenchmarkCopyFile(b *testing.B) {
	dirTest := testDir{
		dir1:     "/Users/lokis/Documents/RebreinGO/sync_directory/internal/test1/",
		dir2:     "/Users/lokis/Documents/RebreinGO/sync_directory/internal/test2/",
		fileName: "test1.txt",
	}
	for i := 0; i < b.N; i++ {
		copyFile(dirTest.dir1, dirTest.dir2, dirTest.fileName)
	}
}
