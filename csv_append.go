package main

import (
	"os"
	"fmt"
	"encoding/csv"
)

func csvWriter(yourSliceGoesHere []string) {
	f, err := os.OpenFile("yourFile.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(yourSliceGoesHere)
	w.Flush()
}

func main() {
	row1 := []string{"Sr.No", "a", "b", "c", "d"}
	//row2 := []string{"1", "a1", "b1", "c1", "d1"}
	csvWriter(row1)
	//csvWriter(row2)
}
