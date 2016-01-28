/*
Longest Compound Word Finder 
created by Aleksandar Slavkovic on Jan 26, 2016
email: sashawam@gmail.com

search() and isCompoundWord(...)  functions
*/
package main

import (
	"sort"
	"fmt"
	"time"
)

// search for the first compound word
func search() string {
	inValidWords := make(map[string]bool)
	start := time.Now()
	for i := range keys {
		for wordListIndex := range words[keys[i]] {
			word := words[keys[i]][wordListIndex]

			if _, ok := inValidWords[word]; !ok && isCompoundWord(word, &inValidWords) {
				fmt.Println("Longest word found in ", time.Now().Sub(start))
				fmt.Println("Size of invalid word set: ", len(inValidWords))
				inValidWords = nil
				return word
			}
		}
	}

	fmt.Println("Size of invalid word set: ", len(inValidWords))
	inValidWords = nil
	return ""
}

// test if a word is a compound word
func isCompoundWord(word string, invalidWords * map[string]bool) bool {

	length := len(word)
		
	for i:= length - minWordLength; i > minWordLength; i-- {
		substr := word[0:i]
		
		if _, ok := (*invalidWords)[substr]; ok {  continue } // we tested this, it does not exist
		
		key := Key{rune(substr[0]), len(substr)}
		bucketLen := len(words[key])
		if bucketLen == 0 { continue } // there is no such a bucket
		
		index := sort.SearchStrings(words[key], substr) // search the prefix in the bucket
		if index < bucketLen && words[key][index] == substr || isCompoundWord(substr, invalidWords) {
			// the prefix exists in the bucket or is a compound word
			// now check the remainder/suffix
			substr = word[i:]
			
			if _, ok := (*invalidWords)[substr]; ok {  continue } // we tested this, it does not exist
			 
			key = Key{rune(substr[0]), len(substr)}
			bucketLen = len(words[key])
			if  bucketLen == 0 { continue } // not found
			
			index = sort.SearchStrings(words[key], substr) // search the prefix in the bucket
			if index < bucketLen && words[key][index] == substr || isCompoundWord(substr, invalidWords) {
				return true
			} else {
				(*invalidWords)[substr] = true
			}
		} else {
			(*invalidWords)[substr] = true
		}
		// else the word does not exist, this is invalid, continue to the next length substring
	}

	return false;
}
