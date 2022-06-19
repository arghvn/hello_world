package main

// note
// Edit this code from the source
import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Go offers built-in support for regular expressions.
	//  Here are some examples of common regexp-related tasks in Go.This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	// Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.
	// This finds the match for the regexp.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// This also finds the first match but returns the start and end indexes for the match instead of the matching text.
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches.
	// For example this will return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// Similarly this will return information about the indexes of matches and submatches.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	// The All variants of these functions apply to all matches in the input, not just the first.
	// For example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// These All variants are available for the other functions we saw above as well.
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

// output :

// true
// true
// peach
// idx: [0 5]
// [peach ea]
// [0 5 1 3]
// [peach punch pinch]
// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
// [peach punch]
// true
// regexp: p([a-z]+)ch
// a <fruit>
// a PEACH
