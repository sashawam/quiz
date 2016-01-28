/*
Longest Compound Word Finder 
created by Aleksandar Slavkovic on Jan 26, 2016
email: sashawam@gmail.com
*/
package main

// structure to save and sort compound map keys by word length
type Key struct {
	key rune
	size int
}

// array of keys
type SortedBySizeKeys []Key

// sorting functions
func (s SortedBySizeKeys) Len() int {
    return len(s)
}
func (s SortedBySizeKeys) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s SortedBySizeKeys) Less(i, j int) bool {
    return /*s[i].key < s[j].key &&*/ s[i].size > s[j].size
}

// adding a unique key utilizes the keysMap set
func (s * SortedBySizeKeys) AddKey(key Key) {
	if _, ok := keysMap[key]; ok { return }
	keysMap[key] = true;
	*s = append(*s, key)
}
