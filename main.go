package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const needTime = 420 // seconds

func main() {
	arTimes, needTimeInt := readArguments()
	fmt.Println(arTimes)
	showTimes(arTimes, needTimeInt)
}

func readArguments() ([]int, int) {
	var arTimes []int
	var needTimeInt int
	var times, needTimeStr string

	flag.StringVar(&times, "times", "", "Times (min): 1,2,3,4...")
	flag.StringVar(&needTimeStr, "needTime", fmt.Sprintf("%d", needTime), "What time need (min)")
	flag.Parse()

	if times == "" {
		panic("empty values")
	}

	strTimes := strings.Split(times, ",")

	for _, v := range strTimes {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("%s convert to int error: %s", i, err))
		}

		arTimes = append(arTimes, i)
	}

	needTimeInt, err1 := strconv.Atoi(needTimeStr)

	if err1 != nil {
		panic("parse need time failed")
	}

	return arTimes, needTimeInt
}

func showTimes(times []int, needTimeInt int) {
	commonTime := 0
	var floatTimes []float32
	for _, v := range times {
		commonTime += v
		floatTimes = append(floatTimes, float32(v))
	}

	remainTime := needTimeInt - commonTime

	for _, v := range floatTimes {
		calc := (v / float32(commonTime)) * float32(remainTime)
		fmt.Printf("%dm + %dm\n", int(math.Round(float64(v))), int(math.Round(float64(calc))))
	}
}
