package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync_directory/internal"
	"syscall"
	"time"
)

func main() {

	logger, err := internal.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	dir := new(internal.Dir)

	mapCache := &internal.MapCache{
		Mpc: make(map[string]struct{}),
	}
	syncCh := make(chan struct{})

	inputDir := internal.InputDir(dir)

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for _ = range ticker.C {
			go internal.ProcessDirChek(inputDir, inputDir.Dir1, mapCache, syncCh, logger)
			<-syncCh
			go internal.ProcessDirChek(inputDir, inputDir.Dir2, mapCache, syncCh, logger)
			<-syncCh
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sig:
			fmt.Println("Ctrl+C: Завершения работы")
			return
		}
	}

}
