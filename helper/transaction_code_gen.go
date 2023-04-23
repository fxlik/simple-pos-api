package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenCode(transactionTypeId int32) string {
	var now = time.Now()
	rand.Seed(now.UnixNano())
	typeString := strconv.Itoa(int(transactionTypeId))
	code := rand.Intn(9999999999)
	codeString := strconv.Itoa(code)
	yearString := strconv.Itoa(now.Year())
	monthString := strconv.Itoa(int(now.Month()))
	kode := typeString + codeString + "/" + monthString + "/" + yearString
	return kode
}
