package main

import (
	"fmt"
	"reflect"
	"strings"
)

func makeKey() *Key {
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

	return &Key{
		encryptionKey: key,
		decryptionKey: revKey,
	}
}

func transform(word string, key map[string]string) (string, error) {
	var transformedWord []string
	for _, letter := range strings.Split(word, "") {
		if _, ok := key[letter]; ok {
			transformedWord = append(transformedWord, key[letter])
		} else {
			return "", fmt.Errorf("word contains an invalid character")
		}
	}
	return strings.Join(transformedWord, ""), nil
}

func (k Key) encrypt(word string) (string, error) {
	return transform(word, k.encryptionKey)
}

func (k Key) decrypt(word string) (string, error) {
	return transform(word, k.decryptionKey)
}

type Key struct {
	encryptionKey map[string]string
	decryptionKey map[string]string
}

func main() {
	k := makeKey()

	if reflect.DeepEqual(k.encryptionKey, k.decryptionKey) {
		fmt.Println("encryption and decryption keys are the same")
	} else {
		fmt.Println("encryption and decryption keys are not the same")
	}

	fmt.Println(k.encrypt("hello"))
	encryptedWord, err := k.encrypt("hello")
	fmt.Println(k.decrypt(encryptedWord))
	fmt.Println(err)
}
