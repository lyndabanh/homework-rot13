package main

import (
	"fmt"
	"reflect"
	"strings"
)

// using *Key in function signature and return &Key
func makeKey() Key {
	alphabet1 := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	alphabet2 := strings.Split("nopqrstuvwxyzabcdefghijklm", "")

	key := make(map[string]string)
	key[" "] = " "
	for i := range alphabet1 {
		key[alphabet1[i]] = alphabet2[i]
		key[strings.ToUpper(alphabet1[i])] = strings.ToUpper(alphabet2[i])
	}

	revKey := make(map[string]string)
	for k, v := range key {
		revKey[v] = k
	}

	return Key{
		encryptionKey: key,
		decryptionKey: revKey,
	}
}

func (k Key) encrypt(word string, key map[string]string) (string, error) {
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

func (k Key) decrypt(word string, key map[string]string) (string, error) {
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
	k := makeKey()

	if reflect.DeepEqual(k.encryptionKey, k.decryptionKey) {
		fmt.Println("encryption and decryption keys are the same")
	} else {
		fmt.Println("encryption and decryption keys are not the same")
	}

	fmt.Println(k.encrypt("hello", k.getEncryptionKey()))
	encryptedWord, err := k.encrypt("hello", k.getEncryptionKey())
	fmt.Println(k.decrypt(encryptedWord, k.getDecryptionKey()))
	fmt.Println(err)
}

// find a way to not have to explicity include capital letters
// make decryption key from the encryption key (instead of having 2 functions)
// make a struct that contains both the encryption key and decryption key
// make encrypt and decrypt be methods on that struct (your struct can be handlers)
// Make 'convert' a method of Key
