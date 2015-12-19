package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// simple Trie stucture
type Trie struct {
	Next  map[uint8]Trie // next Tries, index on the int value of the char
	Match bool           // ends will match
}

// test if the trie contains the string. Tests are recursive, starting with
// the first char to at most len(string)
func (trie *Trie) Contains(item string) (answer bool) {

	items := len(item)
	// if there's still more to the string, we can see if it's contained
	if items > 1 {

		// dig deeper if trie.Next is not nil
		if t, ok := trie.Next[item[0]]; ok {
			return t.Contains(item[1:])
		}
		// end of the line
		return false
	}

	// end of the line, what say you?
	return trie.Match
}

func NewTrie() (trie Trie) {
	trie = Trie{}
	trie.Next = make(map[uint8]Trie)
	return trie
}

// add to the trie. Another recursive call to go from 0 ... len(item)
func (trie *Trie) Add(item string) {

	// if Next hasn't been initialized, do so
	if trie.Next == nil {
		trie.Next = make(map[uint8]Trie)
	}

	// see if were at the last char in the string
	size := len(item)
	if size > 1 {

		// grab the first char in the input
		idx := item[0]

		// see if it's already in the map
		if t, ok := trie.Next[idx]; ok {

			// dig deeper
			t.Add(item[1:])

			// go idom, any changes to t need to be assigned since t is a copy
			trie.Next[idx] = t
			return
		}

		// it's not already in the map, create a new trie, initialize it's map
		// and update our data structure
		t := Trie{}
		t.Next = make(map[uint8]Trie)

		// dig deeper
		t.Add(item[1:])

		// again, go idiom
		trie.Next[idx] = t
		return
	}

	// we're at the end of the line, update match to true
	trie.Match = true
}

// this function loads the trie from a file
func (trie *Trie) Load(filename string) (err error) {
	thefile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open %s read mode. %s\n", filename, err)
		return
	}
	scanner := bufio.NewScanner(thefile)
	for scanner.Scan() {
		trie.Add(scanner.Text())
	}
	thefile.Close()
	return
}

// this fuction reads in a file of strings to check
func (trie *Trie) Check(filename string) (err error) {
	thefile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open %s read mode. %s\n", filename, err)
		return
	}
	scanner := bufio.NewScanner(thefile)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Printf("%s, %t\n", s, trie.Contains(s))
	}
	thefile.Close()
	return
}

func main() {

	inputfile := flag.String("i", "", "file that has strings you want to search")
	testfile := flag.String("s", "", "filename that has strings you want to search for")
	flag.Parse()

	if *inputfile == "" {
		fmt.Println("Need you to specify the abs filename of our input file -i")
		return
	}

	if *testfile == "" {
		fmt.Println("Need you to specify the abs filename of our test file -s")
		return
	}

	t := Trie{}
	err := t.Load(*inputfile)
	if err != nil {
		fmt.Println("bad input file ", *inputfile, " ", err)
		return
	}
	err = t.Check(*testfile)
	if err != nil {
		fmt.Println("bad search file ", *checkfile, " ", err)
		return
	}
}
