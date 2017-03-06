package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main()  {
	// get the book
	res, err := http.Get("http://www.gutenberg.org/cache/epub/54281/pg54281.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scan operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 300)
	// loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}

func HashBucket(word string) int {
	return int(word[0])
}
