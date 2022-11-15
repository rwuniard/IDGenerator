package algorithms

import (
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

var counter int32 = 0

func init() {
	go resetCounter()
}

// Generate ID
// 1 bit for sign, 41 bits for timestamp, 5 bits for data center, 5 bits for machine, 12 bits for sequence.
// The problem right now is that it can generate more than 4096 IDs per millisecond.
func GenerateID(dataCenterId int, machineId int) int64 {
	if (dataCenterId > 31) || (machineId > 31) {
		log.Println("Data center ID and machine ID must be between 0 and 31")
		return -1
	}
	// 1 bit for sign, 41 bits for timestamp, 5 bits for data center, 5 bits for machine, 12 bits for sequence.
	// 1. Calculate the data center id.
	dataCenterBitValue := dataCenterId << (12 + 5)

	// 2. Calculate the machine id.
	machineBitValue := machineId << 12

	id := int64(dataCenterBitValue | machineBitValue)

	// 3. Get current time in milliseconds
	// 4. Subtract the epoch time (1288834974657) from the current time.
	currentTime := time.Now().UnixNano() - time.Unix(1288834974, 657).UnixNano()
	_41_bitTime := currentTime & 2199023255551
	// fmt.Println("currentTime:", currentTime)
	// fmt.Println("_41_bitTime:", _41_bitTime)
	// 5. Shift the timestamp to the left by 22 bits
	timeBitValue := _41_bitTime << (22)

	// 6. Add the data center id and machine id
	id = id | timeBitValue

	// 7. Add the sequence
	atomic.AddInt32(&counter, 1)
	// If the cpu is too fast and it count up to 4095 already
	// sleep for 0.1 millisecond, so the counter will reset to 0 for the next id.
	// The counter is only 12 bits, so it can only count up to 4095.
	if counter >= 4095 {
		log.Println("counter:", counter)
		time.Sleep(100000 * time.Nanosecond)
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
