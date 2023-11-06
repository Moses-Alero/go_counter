package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	err string
}

var display = true

func (s MyError) Error() string {
	return fmt.Sprintf("Error: %v", s.err)
}

// start timer
func startTimer(c chan time.Time, duration time.Duration) (string, error) {
	timerChan := time.NewTimer(duration)
	val := <-timerChan.C
	c <- val
	close(c)
	return "", MyError{"error"}
}

// stop timer
func stopTimer() {
	display = false
	return
}

func displayTime() {
	secondCounter := 0 * time.Second
	for {
		secondCounter += time.Second
		fmt.Printf("\rTime-> %v", secondCounter)
		time.Sleep(time.Second)
		if display == false {
			break
		}
	}
	return
}

func fun() {
	fmt.Println("Hello, World!")
	var input string
	var num int
	var duration time.Duration
	fmt.Print("Enter a how long timer should last: ")
	for {
		fmt.Scan(&input)
		val := string(input[0])
		number, err := strconv.Atoi(val)
		if len(input) < 2 {
			num = number
			duration = time.Duration(num * int(time.Second))
			break
		}
		if err != nil {
			err := MyError{"Invalid Number"}
			fmt.Println(err)
		} else {
			num = number
			timerDuration := string(input[1])
			switch string(timerDuration[0]) {
			case "d":
				duration = time.Duration(num * 24 * int(time.Hour))
			case "h":
				duration = time.Duration(num * int(time.Hour))
			case "m":
				duration = time.Duration(num * int(time.Minute))
			case "s":
				duration = time.Duration(num * int(time.Second))
			default:
				duration = time.Duration(num * int(time.Second))
			}
			break
		}

	}

	c := make(chan time.Time)

	go startTimer(c, duration)
	go displayTime()
	<-c
	stopTimer()

}
