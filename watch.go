package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

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

	start := func() {
		elm = exec.Command("bash", "run")
		//elm.Dir = "server/out"
		elm.Stdin = os.Stdin
		elm.Stdout = os.Stdout
		elm.Stderr = os.Stdout
		err = elm.Start()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("PID: %v", elm.Process.Pid)
	}

	kill := func() {
		log.Printf("killing PID: %v", elm.Process.Pid)

		err = elm.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}

		err = elm.Process.Signal(syscall.SIGKILL)
		if err != nil {
			log.Fatal(err)
		}

		err = elm.Process.Release()
		if err != nil {
			log.Fatal(err)
		}

		elm = nil
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
