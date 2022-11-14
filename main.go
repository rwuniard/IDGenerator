package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var counter int32

func init() {
	counter = 0
}

func main() {
	go resetCounter()
	var id int64
	start := time.Now()
	fmt.Println("Start Time:", start)
	for i := 0; i < 10000; i++ {
		id = GenerateID(1, 1)
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

	// 2. Calculate the machine id.
	machineBitValue := machineId << 12

	id := int64(dataCenterBitValue | machineBitValue)

	// 3. Get current time in milliseconds
	// 4. Subtract the epoch time (1288834974657) from the current time.
	currentTime := time.Now().UnixNano() - time.Unix(1288834974, 657).UnixNano()

	// 5. Shift the timestamp to the left by 22 bits
	timeBitValue := currentTime << (22)

	// 6. Add the data center id and machine id
	id = id | timeBitValue

	// 7. Add the sequence
	atomic.AddInt32(&counter, 1)
	if counter > 4095 {
		fmt.Println("counter:", counter)
	}
	id = id | int64(counter)

	return id
}

// This function will reset the counter to 0 every 0.3 millisecond.
func resetCounter() {
	for {
		time.Sleep(300000 * time.Nanosecond)
		atomic.StoreInt32(&counter, 0)
		runtime.Gosched()
	}
}
