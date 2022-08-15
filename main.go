package main

import (
	"fmt"
	"strings"
)

func makeKey() map[string]string {
	alphabet1 := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	alphabet2 := strings.Split("nopqrstuvwxyzabcdefghijklm", "")

	key := make(map[string]string)
	key[" "] = " "
	for i := range alphabet1 {
		key[alphabet1[i]] = alphabet2[i]
		key[strings.ToUpper(alphabet1[i])] = strings.ToUpper(alphabet2[i])
	}
	return key
}

func convert(word string, key map[string]string) string {
	var convertedWord []string
	for _, letter := range strings.Split(word, "") {
		if _, ok := key[letter]; ok {
			convertedWord = append(convertedWord, key[letter])
		} else {
			return "word contains an invalid character"
		}
	}
	return strings.Join(convertedWord, "")
}

type Key struct {
	encryptionKey map[string]string
}

func (k Key) getKey() map[string]string {
	return k.encryptionKey
}

func main() {

	k := Key{makeKey()}
	fmt.Println(convert("hello", k.getKey()))
	fmt.Println(convert(convert("hello", k.getKey()), k.getKey()))

}

// find a way to not have to explicity include capital letters
// make decryption key from the encryption key (instead of having 2 functions)
// make a struct that contains both the encryption key and decryption key
// make encrypt and unecrypt be methods on that struct (your struct can be handlers)
