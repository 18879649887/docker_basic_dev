package flag

import (
	"fmt"
	"strings"
	"strconv"
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
	fmt.Println("wocker 2.0");
}

func CheckName(name string) string {
	return name;
}

func CheckPort(port string) (int, int) {
	s := strings.Split(port, ":");
	if len(s) == 2 {
		basicport, err1 := strconv.Atoi(s[0]);
		containerport, err2 := strconv.Atoi(s[1]);
		if (err1 != nil) && (err2 != nil) {
			fmt.Println(err1);
			return -1, -1;		
		} else {
			return basicport, containerport;
		}
	} else {
		fmt.Println("useless ports");
		return -1, -1;
	}
}

func CheckParams(flagstring *SFlag) (string, int, int) {
	//正则表达优化
	var name string;
	var port1 int;
	var port2 int;

	name_key := flagstring.Key[0];
	name_value := flagstring.Value[0];
	port_key := flagstring.Key[1];
	port_value := flagstring.Value[1];

	if name_key == "-n" {
		name = CheckName(name_value);
	} else {
		name = "";
	}

	if port_key == "-p" {
		port1, port2 = CheckPort(port_value);   
	} else {
		port1 = -1;
		port2 = -1;
	}

	return name, port1, port2;
}

func Check(err error) {
	if err != nil {
		fmt.Println(err);
	}
}