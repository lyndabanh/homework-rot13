package main

import (
	"fmt"
	"strings"
)

// makeKey could return Key instead of a map
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

// Make 'convert' a method of Key
func encrypt(word string, key map[string]string) (string, error) {
	var encryptedWord []string
	for _, letter := range strings.Split(word, "") {
		if _, ok := key[letter]; ok {
			encryptedWord = append(encryptedWord, key[letter])
		} else {
			return "", fmt.Errorf("word contains an invalid character")
		}
	}
	return strings.Join(encryptedWord, ""), nil
}

func decrypt(word string, key map[string]string) (string, error) {
	var decryptedWord []string
	for _, letter := range strings.Split(word, "") {
		if _, ok := key[letter]; ok {
			decryptedWord = append(decryptedWord, key[letter])
		} else {
			return "", fmt.Errorf("word contains an invalid character")
		}
	}
	return strings.Join(decryptedWord, ""), nil
}

type Key struct {
	encryptionKey map[string]string
	decryptionKey map[string]string
}

func (k Key) getEncryptionKey() map[string]string {
	return k.encryptionKey
}

func (k Key) getDecryptionKey() map[string]string {
	return k.decryptionKey
}

// k.covert("hello")
func main() {
	k := Key{makeKey(), makeKey()}
	fmt.Println(encrypt("hello", k.getEncryptionKey()))
	encryptedWord, err := encrypt("hello", k.getEncryptionKey())
	fmt.Println(decrypt(encryptedWord, k.getDecryptionKey()))
	fmt.Println(err)
}

// find a way to not have to explicity include capital letters
// make decryption key from the encryption key (instead of having 2 functions)
// make a struct that contains both the encryption key and decryption key
// make encrypt and decrypt be methods on that struct (your struct can be handlers)
// convert func could be method on key type
