package internal

import "flag"

func InputDir(dir *Dir) *Dir {
	flag.StringVar(&dir.Dir1, "dir1", "/Users", "Absolut path to the directory")
	flag.StringVar(&dir.Dir2, "dir2", "/Users", "Absolut path to the directory")
	flag.Parse()
	return dir
}
