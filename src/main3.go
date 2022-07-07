package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

var buf [1]byte // This buffer had come out of the function. Doing it, prevents the runtime from escaping this buffer to the heap everytime readbyte is called. It escapes the buffer because it's created on the stack and is passed to an interface. Since the runtime does not know wether the interface passed to readbyte will need the buffer later, it escapes it. Creating the buffer in the global escope, prevents this buffer from being reallocated everytime readbyte is called, and consecutively being escaped to ensure whatever interface needs it will be able to access it in the future.

func readbyte(r io.Reader) (rune, error) {
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}
	words := 0
	inword := false
	b := bufio.NewReader(f) //Created a buffer to read from file via buffer.
	for {
		r, err := readbyte(b) //reading from buffer
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words \n", os.Args[1], words)

}
