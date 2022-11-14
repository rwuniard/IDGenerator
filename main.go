package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var counter int32
var resetCount int32
var prevId int64
var prevTime int64
var prevCounter int32

func init() {
	counter = 0
	resetCount = 0
	prevId = 0
	prevTime = 0
	prevCounter = 0
}

func main() {
	go resetCounter()
	var id int64
	start := time.Now()
	fmt.Println("Start Time:", start)
	for i := 0; i < 10; i++ {
		id = GenerateID(1, 1)
		// fmt.Println("id:", id)
	}
	end := time.Now()
	fmt.Println("End Time:", end)
	fmt.Println("Duration:", end.Sub(start))
	fmt.Println("id:", id)
	currentTime := time.Now().UnixNano() - time.Unix(1288834974, 657).UnixNano()
	fmt.Println(time.Unix(1288834974, 657).UnixMilli())
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(currentTime)
	a := currentTime & 2199023255551
	fmt.Println("a:", a)
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
	currentTime := time.Now().UnixNano() - time.Unix(1288834974, 657).UnixNano()
	// fmt.Println("currentTime:", currentTime)
	// 5. Shift the timestamp to the left by 22 bits
	timeBitValue := currentTime << (22)
	// 6. Add the data center id and machine id
	id = id | timeBitValue
	// 7. Add the sequence
	atomic.AddInt32(&counter, 1)
	// fmt.Println("This counter:", counter)
	if counter > 4095 {
		fmt.Println("counter:", counter)
	}
	id = id | int64(counter)
	// fmt.Println("counter:", counter)
	if prevId == id {
		fmt.Println("prevId:", prevId)
		fmt.Println("id:    ", id)
		fmt.Println("prevTime:", prevTime)
		fmt.Println("currentTime:", currentTime)
		fmt.Println("prevCounter:", prevCounter)
		fmt.Println("current counter:", counter)
	} else {
		prevId = id
		prevTime = currentTime
		prevCounter = counter
	}
	return id
}

// This function will reset the counter to 0 every 1 millisecond.
func resetCounter() {
	for {
		time.Sleep(300000 * time.Nanosecond)
		atomic.StoreInt32(&counter, 0)
		atomic.AddInt32(&resetCount, 1)
		runtime.Gosched()
	}
}
