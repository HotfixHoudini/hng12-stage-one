package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

const NUMBERS_API_URL = "http://numbersapi.com"

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Number Classification API")
	})

	r.GET("/api/classify-number", classifyNumber)

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func classifyNumber(c *gin.Context) {
	numberStr := c.Query("number")
	if numberStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"number": "invalid input", "error": true})
		return
	}

	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"number":  numberStr,
			"message": "Invalid number format",
		})
		return
	}

	// Convert negative numbers to positive
	var positiveNumber int

	parsedNumber := int(number)
	if number < 0 {
		positiveNumber = int(math.Abs(number)) // Convert to int
	} else {
		positiveNumber = parsedNumber
	}

	funcFact, err := fetchFunFact(parsedNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": fmt.Sprintf("Error fetching from Numbers API: %s", err.Error()),
		})
		return
	}

	numberProperties := getNumberProperties(positiveNumber)

	response := Response{
		Number:     parsedNumber,
		IsPrime:    isPrime(positiveNumber),
		IsPerfect:  isPerfect(positiveNumber),
		Properties: numberProperties,
		DigitSum:   sumOfDigits(positiveNumber),
		FunFact:    funcFact,
	}

	c.JSON(http.StatusOK, response)
}

func fetchFunFact(num int) (string, error) {
	client := resty.New()
	formattedUrl := fmt.Sprintf("%s/%d/math", NUMBERS_API_URL, num)
	response, err := client.R().Get(formattedUrl)

	if err != nil {
		return "", err
	}

	return response.String(), nil
}

func getNumberProperties(num int) []string {
	properties := []string{}
	if isArmstrong(num) {
		properties = append(properties, "armstrong")
	}
	if num%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	return properties
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(num int) bool {
	if num < 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i != num/i {
				sum += num / i
			}
		}
	}
	return sum == num && num != 1
}

func isArmstrong(num int) bool {
	digits := intToDigits(num)
	power := len(digits)
	sum := 0
	for _, digit := range digits {
		sum += intPow(digit, power)
	}
	return sum == num
}

func sumOfDigits(num int) int {
	digits := intToDigits(num)
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	return sum
}

func intToDigits(num int) []int {
	strNum := strconv.Itoa(num)
	digits := make([]int, len(strNum))
	for i, digit := range strNum {
		digits[i], _ = strconv.Atoi(string(digit))
	}
	return digits
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
