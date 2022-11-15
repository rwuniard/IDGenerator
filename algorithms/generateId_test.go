package algorithms

import (
	"fmt"
	"testing"
)

func TestGenerateID(t *testing.T) {
	counter = 0
	// go resetCounter()
	var id int64
	ids := make(map[int64]int64, 0)
	numberLoop := 100000000
	for i := 0; i < numberLoop; i++ {
		id = GenerateID(1, 1)
		ids[id] = id
		tempId := id >> 12
		machineId := tempId & 31 // 00000000000000000000000000011111
		if machineId != 1 {
			t.Errorf("machineId is not 1, it is %d", machineId)
		}
		dataCenterId := (tempId >> 5) & 31 // 00000000000000000000000000011111

		if dataCenterId != 1 {
			t.Errorf("dataCenterId is not 1, it is %d", dataCenterId)
		}
	}
	if len(ids) != numberLoop {
		fmt.Println("len(ids):", len(ids))
		fmt.Println("numberLoop:", numberLoop)
		fmt.Println("delta:", numberLoop-len(ids))
		t.Errorf("There are duplicate ids")
	}
	fmt.Println("id:    ", id)
}
