package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func Err(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func Sha256File(filePath string) string {
	file, err := os.Open(filePath)
	Err(err)
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	Err(err)

	return fmt.Sprintf("%x", hash.Sum(nil))

}

func PregReplace(expression string, replace string, input string) string {

	reg, err := regexp.Compile(expression)
	Err(err)

	var bk = []byte(input)
	var out []byte

	if reg.Match(bk) {
		for _, sub := range reg.FindAllSubmatchIndex(bk, -1) {
			out = reg.Expand(out, []byte(replace), bk, sub)
		}
	}
	return string(out)
}
