/*
Longest Compound Word Finder 
created by Aleksandar Slavkovic on Jan 26, 2016
email: sashawam@gmail.com

// loadWords function is
// using bufio.Scanner to load the words into memory from a text file
// while the words are being loaded into buckets, the keys array is also being created
*/

package main
/* load the words from a text file
   store the words in buckets with a compound key of [first letter of the word, word size]
   add each unique key to the keys array
   once loaded, sort the keys array by the length of the word
*/

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	//"time"
	"fmt"
)

func loadWords(path string) error {
	var word string
	var key Key
	
	fmt.Println("Reading input file " + path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer f.Close()
	
	//var start time.Time
	size := 0;
	
	// initialize the words and keys structures, 
	// segment the words into buckets by first character/rune and the length of the string
	
	
	//addKeysDuration := time.Duration(0)
    //addWordsDuration := time.Duration(0)
	
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
		
	for scanner.Scan() {
		size++
		word = strings.ToLower(scanner.Text())
		key = Key{rune(word[0]), len(word)}
		
		//start = time.Now();
		// add unique key to the Keys array
		keys.AddKey(key)
		//addKeysDuration += time.Now().Sub(start)
		
		//start = time.Now();
		// append the word to the bucket
		words[key] = append(words[key], word) 
		//addWordsDuration += time.Now().Sub(start)
		
	}

	fmt.Println("Imported " + strconv.Itoa(size) + " words")
	//fmt.Println("add keys duration  = ", addKeysDuration)
	//fmt.Println("add words duration = ", addWordsDuration)
	
	return scanner.Err()
}

