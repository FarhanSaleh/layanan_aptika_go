package helper

import (
	"encoding/json"
	"fmt"
)

func PrintData(data any) {
	d, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(d))
}