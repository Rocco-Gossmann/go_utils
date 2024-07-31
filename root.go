package go_utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
)

const DATETIME_PRINT = "02. Jan. 2006 15:04"

type ControlledPanic struct {
	Msg      string
	ExitCode int
}

func Exitf(statement string, args ...any) {
	panic(ControlledPanic{
		Msg:      fmt.Sprintf(statement, args...),
		ExitCode: 0,
	})
}

func Failf(statement string, args ...any) {
	panic(ControlledPanic{
		Msg:      fmt.Sprintf(statement, args...),
		ExitCode: 1,
	})
}
func Err(err any) {
	if err != nil {
		panic(err)
	}
}

func Suffix(cnt int, singular string, plural string) string {
	ret := plural
	if cnt == 1 || cnt == -1 {
		ret = singular
	}

	return ret
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

func SecToTimePrint(secondCount float64) string {
	var mins = math.Floor(secondCount / 60)
	var hrs = math.Floor(mins / 60)

	var secs = secondCount - (mins * 60)
	mins -= hrs * 60
	return fmt.Sprintf("%02d:%02d:%02d", int(hrs), int(mins), int(secs))
}
