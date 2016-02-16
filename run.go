package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/exp/inotify"
)

const (
	elmPort = 8081
)

func main() {
	var err error

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.AddWatch("client", inotify.IN_MODIFY)
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.AddWatch("server", inotify.IN_MODIFY)
	if err != nil {
		log.Fatal(err)
	}

	var elm *exec.Cmd
	var server *exec.Cmd

	start := func() {
		elm = exec.Command("elm", "reactor", fmt.Sprintf("--port=%v", elmPort))
		elm.Dir = "client"
		err = elm.Start()
		if err != nil {
			log.Fatal(err)
		}

		server = exec.Command("go", "run", "main.go")
		server.Dir = "server"
		err = server.Start()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("client: %#v, server: %#v", elm, server)
	}

	kill := func() {
		err = elm.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}

		err = server.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}
	}

	start()
	defer kill()

	for {
		select {
		case ev := <-watcher.Event:
			log.Printf("modified file: %v", ev.Name)
			kill()
			start()
		}
	}
}
