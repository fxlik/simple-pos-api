package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestGenerateCode(t *testing.T) {
	var now = time.Now()
	rand.Seed(now.UnixNano())
	code := rand.Intn(9999999999)
	codeString := strconv.Itoa(code)
	yearString := strconv.Itoa(now.Year())
	monthString := strconv.Itoa(int(now.Month()))
	kode := codeString + "/" + monthString + "/" + yearString
	fmt.Println(kode)
}
