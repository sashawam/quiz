Longest Compound Word Finder 
created by Aleksandar Slavkovic on Jan 26, 2016
email: sashawam@gmail.com

INPUT: a text file with one word per line
OUTPUT: longest compound word consisting of words found in INPUT

ASSUMPTIONS: 
	input file is a text file 
	input file fits into memory
	strings do not have to be ASCII
	word list does not have to be alhabetically sorted
	word list does not have to be unique
	
GOAL: minimize the search time for finding the longest compound word in the list
	
ALGORITHMIC APPROACH:
	let's start with the longest word
			the testing is accomplished by checking {prefix, suffix} pairs of the word
				prefix is a substring of the word starting with 0 index and having a variable length	
				suffix is the remaining substring of the word with the starting index of length(prefix)
			if both prefix and suffix are found in the word list or are compound words, then the word is a compound word
				this approach requires reccursion
			let's use sort.SearchStrings([]string, string) to determine if prefix/suffix is in the word list
			SearchStrings(...) is quite expensive for large string arrays
	
	by utilizing additional data structures, we segment the words list and minimize the search space for each iteration; 
	in other words,
		the []string input to the sort.SearchStrings will be based on the size of the substring being searched;
		since the words are stored in buckets, we will search much smaller arrays during each iteration;
		if a substring is not found in the list, we will store it in a lookup table (as a map[string]bool)	
			and check that lookup table before a search
	
DATA STRUCTURES:
	The idea is to reduce the search space at each iteration
	input words are loaded into buckets defined as a map of string slices
		a bucket is accessed via a unique compound key and contains an array of words of the same length
			the unique compound key = {<first letter of a word>, <word length>}
			thus, all of the words with the same first letter and the same length are stored in the same sub-array
	unique keys are stored in an array
		the keys array is sorted by <word length> field
	set of inValidWords implemented as a map[string]bool; 
	
JUSTIFICATION FOR ADDITIONAL DATA STRUCTURES:
	Keys array is much smaller than the words list; for the test file, we have 263533 records and 538 buckets
	Time needed to create the Keys array and to sort it is negligable compared to loading the words into memory
	
	Keys array is sorted by word size in descending order -- thus, the first compound word found is the longest one
		
	Set of invalid words (words searched and not found in the words list) is maintained to avoid searching them repeatedly
		during reccursions. This structure can be kept to a certain size if needed at the expense of repeated searches;
		It could also be localized per bucket if using too much space.
		The search algorithm is 200x faster when using the invalidWords set
		
	This approach finds the longest compound word in under 0.5 milliseconds 
	on an Intel i7 2.2GH machine (for the sample data set; not counting the time to load the words into memory from the disk)

OUTPUT EXAMPLE for the sample INPUT file of 263533 words:
	Reading input file <local path>/wordlist.txt
	Imported 263533 words
	Loading the strings took  123.636619ms
		Number of buckets =   538
	(Sorting the key array took  46.609µs)
		testing word: pneumonoultramicroscopicsilicovolcanoconiosis
		testing word: dichlorodiphenyltrichloroethanes
		testing word: dichlorodiphenyltrichloroethane
		testing word: floccinaucinihilipilifications
		testing word: antidisestablishmentarianisms
	Longest compound word found: antidisestablishmentarianisms
	Longest compound word size:  29
	Elapsed:  374.702µs
	Size of invalidWords Set is  174

PARALLEL CODE:
  uses a pool of initalized go routines; task requests are sent via 'input' channel
  performance for the provided input file is not better than the serial version; this is 
  due to chan communcation overhead with the parallel version.
  
BUILD:
	go build -o lcword_search ./compword
	
RUN:
	./lcword_search <local path>/wordlist.txt
	./lcword_search <local path>/wordlist.txt parallel
