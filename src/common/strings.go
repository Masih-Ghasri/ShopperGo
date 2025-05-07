package common

import (
	"github.com/Masih-Ghasri/GolangBackend/config"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	lowerCharSet   = regexp.MustCompile("[^a-z]")
	upperCharSet   = regexp.MustCompile("[^A-Z]")
	specialCharSet = regexp.MustCompile("[^a-z]")
	numberSet      = regexp.MustCompile("[0-9]+")
	allCharSet     = lowerCharSet.String() + upperCharSet.String() + specialCharSet.String() + numberSet.String()
)

func CheckPassword(password string) bool {
	cfg := config.Getconfig()
	if len(password) < cfg.Password.MinLength {
		return false
	}

	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}

	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}

	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}

	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}

	return true
}

func GeneratePassword() string {
	var password strings.Builder

	cfg := config.Getconfig()
	passwordLength := cfg.Password.MinLength + 2
	minSpecialChar := 2
	minNum := 3
	if !cfg.Password.IncludeDigits {
		minNum = 0
	}

	minUpperCase := 3
	if !cfg.Password.IncludeUppercase {
		minUpperCase = 0
	}

	minLowerCase := 3
	if !cfg.Password.IncludeLowercase {
		minLowerCase = 0
	}

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet.String()))
		password.WriteString(string(specialCharSet.String()[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet.String()))
		password.WriteString(string(numberSet.String()[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet.String()))
		password.WriteString(string(upperCharSet.String()[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(lowerCharSet.String()))
		password.WriteString(string(lowerCharSet.String()[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func GenerateOtp() string {
	cfg := config.Getconfig()
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))   // 10^d-1 100000
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1) // 999999 = 1000000 - 1 (10^d) -1

	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
