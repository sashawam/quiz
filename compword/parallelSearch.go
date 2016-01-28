/* 
Longest Compound Word Finder 
created by Aleksandar Slavkovic on Jan 26, 2016
email: sashawam@gmail.com

concurrent code with worker pool and task generator for compound word search

parallelSearch(...)
taskGenerator(...)
worker(...)

*/
package main

import (
	"time"
	"fmt"
)

// using a goroutine pool and sending task requests via an input channel
// signaling the end of work is done via 'quit' channel
// we block waiting for the first word to be returned from the workers
// the 'result' channel is buffered to ensure we do not deadlook
// the taskGenerator sends buckets' indexes via 'input' channel; thus, each worker is assigned a unique bucket
func parallelSearch(numProcessors int, words MyMapList, keys SortedBySizeKeys) string {
	
	fmt.Println("Running in parallel with ", numProcessors, " logical CPUs...")
	
	var result = make (chan string, 3 * numProcessors) // make channel for returning the result, buffer it***
	// *** we can have additional goroutines return results as we are shutting down
	//		thus, to avoid a deadlock, we will allow enough space in the channel for all goroutines to potentially return results
	var input = make (chan int, numProcessors) // make channel for inputing data into worker threads
	var quit = make (chan bool) // signal goroutine completion
	var quitSendData = make (chan bool) // signal goroutine completion
	
	inValidWords := make(map[string]bool)
	// set the worker goroutines, one per processor
	for j := 0; j< numProcessors; j++ {
		go worker(&inValidWords, j, input, result, quit)
	}
	
	start := time.Now()
	// call the workers
	go taskGenerator(len(keys), numProcessors, input, result, quitSendData)

	// block until the first result comes in
	resultString := <-result
	
	fmt.Println("Longest word found in ", time.Now().Sub(start))
	
	// notify the task generator to stop
	quitSendData <- true
	
	// notify all worker goroutines to stop	
	for j := 0; j< numProcessors; j++ {
		quit <- true
	}
			
	return resultString
}

// test buckets in parallel by sending bucket id to a worker goroutine via input channel
func taskGenerator(keysArrayLen int, numProcessors int, in chan int, result chan string, quitdata chan bool ) {
	for i := 0; i < keysArrayLen; i = i + numProcessors {
		select {
			case <- quitdata: {
				fmt.Println("Done sending data to workers.")
				return
			}
			default:
				for j := 0; j< numProcessors; j++ {
					in <- i + j // send request to goroutine
				}
		}
		time.Sleep(time.Microsecond * 1)
	}
	
	// if nothing found, return an empty string so the program can terminate
	result <- ""
}

// this is where the search is taking place
// each worker maintains its own copy of the invalidWords set
// result channel is used only if the worker found a compound word
// it blocks on in and quit channels
func worker(invalidWords * map[string]bool, wid int, in <-chan int, result chan<- string, quit <-chan bool ) {
	//invalidWords := make(map[string]bool)

	threadid := wid
	//fmt.Println("Worker goroutine created: ", threadid)
	
	start := time.Now()
	for {
		select {
			case <- quit: {
				// finish goroutine
				//fmt.Println("Worker goroutine ", threadid, " finished. Elapsed time ", time.Now().Sub(start))
				//invalidWords = nil
				return
			}
			case i := <- in: {	
				//start := time.Now()			
				//if i < len(keys) {
					for wordListIndex := range words[keys[i]] {
						word := words[keys[i]][wordListIndex]
						 
						fmt.Println("\t\tWorker goroutine ", threadid, " testing word: ", word)
						if _, found := (*invalidWords)[word]; !found && isCompoundWord(word, invalidWords) {
							fmt.Println("Longest word found in ", time.Now().Sub(start), " word is ", word)
							result <- word 
							break
						} 
					}
			}
		}
		time.Sleep(time.Microsecond * 1)
	}	
}
