package internal

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Dir struct {
	Dir1 string
	Dir2 string
}

func copyFile(src, dst, fileName string) (int64, error) {
	srcOpen, err := os.Open(src + fileName)
	if err != nil {
		return 0, err
	}
	defer srcOpen.Close()

	dstCreate, err := os.Create(dst + fileName)
	if err != nil {
		return 0, err
	}
	defer dstCreate.Close()

	_, err = io.Copy(dstCreate, srcOpen)
	if err != nil {
		return 0, err
	}

	fileStat, err := dstCreate.Stat()
	if err != nil {
		return 0, err
	}
	size := fileStat.Size()
	return size, nil
}

func noContainsMapCache(dirFiles []os.DirEntry, val string) bool {

	set := make(map[string]struct{}, len(dirFiles))
	for _, file := range dirFiles {
		set[file.Name()] = struct{}{}
	}

	_, ok := set[val]

	return !ok
}

func processDeleteFile(dir *Dir, inputDir string, dirFiles []os.DirEntry, mapCache *MapCache, logger *logrus.Logger) {
	for keyFileName := range mapCache.Mpc {

		if noContainsMapCache(dirFiles, keyFileName) {
			mapCache.delete(keyFileName)
			if inputDir == dir.Dir1 {
				err := os.Remove(dir.Dir2 + keyFileName)
				logger.WithFields(logrus.Fields{
					"Dir2":     dir.Dir2,
					"FileName": keyFileName,
				}).Info("Delete file")
				if err != nil {
					logger.WithFields(logrus.Fields{
						"Dir2":     dir.Dir2,
						"FileName": keyFileName,
					}).Errorf("failed to delete file: %s", err)
				}
			} else {
				err := os.Remove(dir.Dir1 + keyFileName)
				logger.WithFields(logrus.Fields{
					"Dir1":     dir.Dir1,
					"FileName": keyFileName,
				}).Info("Delete file")
				if err != nil {
					logger.WithFields(logrus.Fields{
						"Dir1":     dir.Dir1,
						"FileName": keyFileName,
					}).Errorf("failed to delete file: %s", err)
				}

			}
		}

	}
}

func readDir(inputDir string, logger *logrus.Logger) []os.DirEntry {

	dirFiles, err := os.ReadDir(inputDir)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"Dir": inputDir,
		}).Fatalf("Directory failed: %s", err)
	}

	return dirFiles
}

func processCopyFile(dirFiles []os.DirEntry, dir *Dir, mapCache *MapCache, inputDir string, logger *logrus.Logger) {

	for _, file := range dirFiles {

		if !mapCache.chekAddFile(file.Name()) {
			mapCache.add(file.Name())
			if inputDir == dir.Dir1 {
				size, err := copyFile(dir.Dir1, dir.Dir2, file.Name())
				logger.WithFields(logrus.Fields{
					"Dir1":     dir.Dir1,
					"Dir2":     dir.Dir2,
					"FileName": file.Name(),
					"Size":     size,
				}).Info("Copy file")
				if err != nil {
					logger.WithFields(logrus.Fields{
						"Dir1":     dir.Dir1,
						"Dir2":     dir.Dir2,
						"FileName": file.Name(),
					}).Errorf("file failed to copy: %s", err)
				}
			} else {
				size, err := copyFile(dir.Dir2, dir.Dir1, file.Name())
				logger.WithFields(logrus.Fields{
					"Dir2":     dir.Dir2,
					"Dir1":     dir.Dir1,
					"FileName": file.Name(),
					"Size":     size,
				}).Info("Copy file")
				if err != nil {
					logger.WithFields(logrus.Fields{
						"Dir2":     dir.Dir2,
						"Dir1":     dir.Dir1,
						"FileName": file.Name(),
					}).Errorf("file failed to copy: %s", err)
				}

			}
		}

	}
}

func ProcessDirChek(dir *Dir, inputDir string, mapCache *MapCache, syncChan chan struct{}, logger *logrus.Logger) {
	dirFiles := readDir(inputDir, logger)
	processCopyFile(dirFiles, dir, mapCache, inputDir, logger)
	processDeleteFile(dir, inputDir, dirFiles, mapCache, logger)
	syncChan <- struct{}{}
}
