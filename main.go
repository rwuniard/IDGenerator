package main

import (
	"fmt"
	"time"
)

var counter int

func init() {
	counter = 0
}

func main() {
	go resetCounter()
	var id int64
	start := time.Now()
	fmt.Println("Start Time:", start)
	for i := 0; i < 100000; i++ {
		id = GenerateID(1, 1)
		// fmt.Println("id:", id)
	}
	end := time.Now()
	fmt.Println("End Time:", end)
	fmt.Println("Duration:", end.Sub(start))
	fmt.Println("id:", id)
}

// Generate ID
// 1 bit for sign, 41 bits for timestamp, 5 bits for data center, 5 bits for machine, 12 bits for sequence.
// The problem right now is that it can generate more than 4096 IDs per millisecond.
func GenerateID(dataCenterId int, machineId int) int64 {
	// 1 bit for sign, 41 bits for timestamp, 5 bits for data center, 5 bits for machine, 12 bits for sequence.
	// 1. Calculate the data center id.
	dataCenterBitValue := dataCenterId << (12 + 5)
	// fmt.Println("dataCenterBitValue:", dataCenterBitValue)
	// 2. Calculate the machine id.
	machineBitValue := machineId << 12
	// fmt.Println("machineBitValue:", machineBitValue)
	id := int64(dataCenterBitValue | machineBitValue)
	// fmt.Println("id:", id)
	// 3. Get current time in milliseconds
	// 4. Subtract the epoch time (1288834974657) from the current time.
	currentTime := time.Now().UnixMilli() - time.Unix(1288834974, 657).UnixMilli()
	// fmt.Println("currentTime:", currentTime)
	// 5. Shift the timestamp to the left by 22 bits
	timeBitValue := currentTime << (22)
	// 6. Add the data center id and machine id
	id = id | timeBitValue
	// 7. Add the sequence
	counter++
	id = id | int64(counter)
	// fmt.Println("counter:", counter)
	return id
}

// This function will reset the counter to 0 every 1 millisecond.
func resetCounter() {
	for {
		time.Sleep(1 * time.Millisecond)
		counter = 0
	}
}
