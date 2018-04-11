package flag

import (
	"fmt"
)

type SFlag struct {
	Command string
	Key []string
	Value []string
}

func (flagstr *SFlag) Resolve(arr []string) *SFlag {
	//flagstr := new(SFlag);

	//解析出指令
	flagstr.Command = arr[0];

	//解析出key-value
	for v, out := range arr[1:] {
		if num := v % 2; num == 0 {
			flagstr.Key = append(flagstr.Key, out);
		} else {
			flagstr.Value = append(flagstr.Value, out);
		}
	}
	return  flagstr;
} 

func ShowVersion() {
	fmt.Println("wocker 1.0");
}
