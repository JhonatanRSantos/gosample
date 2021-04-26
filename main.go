// Package main it's the base package of a Go program
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/JhonatanRSantos/gosample/util"
)

// main it's the entry point of a Go program
func main() {
	startTime := time.Now()
	writer := initLogger()
	defer writer.Close()

	log.Println("The system started correctly :)")
	lineReader := util.NewLineReader()

	log.Println("Total cores", runtime.NumCPU())
	log.Println("Total Goroutines", runtime.NumGoroutine())

	username := lineReader.Read("What's your name?")
	log.Println("The user name is", username)
	goRoutines(1000)
	log.Println("The total elasped time is", time.Now().Sub(startTime).Seconds(), "seconds")
}

// initLogger points all system logs to a file on current dir
func initLogger() *os.File {
	writer, err := os.Create(
		fmt.Sprintf("logger-%s.log", "x"),
	)

	if err != nil {
		log.Fatalln("Ops, we can't create the logger file. Cause:", err)
	}

	log.SetOutput(writer)
	return writer
}

// goRoutines provide an example of how create goroutines
func goRoutines(totalRoutines int) {
	wg := sync.WaitGroup{}

	wg.Add(totalRoutines)

	for routineIndex := 1; routineIndex <= totalRoutines; routineIndex++ {
		go func(routineIndex int) {
			log.Println(fmt.Sprintf("The routine %d started correctly", routineIndex))
			runtime.Gosched()
			log.Println("Let's finish the routine", routineIndex)
			wg.Done()
		}(routineIndex)
	}

	wg.Wait()
}
