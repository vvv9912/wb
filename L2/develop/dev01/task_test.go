package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	curTime, err := gettime()
	if err != nil {
		fmt.Println(err)
	}
	cutTime2 := time.Now()

	curTimeSec := curTime.Hour()*60*60 + curTime.Minute()*60 + curTime.Second()

	curTimeSec2 := cutTime2.Hour()*60*60 + cutTime2.Minute()*60 + cutTime2.Second()

	deltaTime := curTimeSec2 - curTimeSec

	if math.Abs(float64(deltaTime)) > 30 { //Сравнение приблизительно с точностью до 30 с
		t.Error("Time delay", deltaTime, "\nNTP time:", curTime, "\nCurrent Time:", cutTime2)
	}

	fmt.Println("NTP time:", curTime, "\nCurrent Time:", cutTime2)
}
