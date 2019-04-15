package getinput

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func Call(msg string, t string) interface{} {
	if !strings.HasSuffix(msg, " ") {
		msg = msg + " "
	}
	fmt.Print(msg)

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return Call(msg, t)
	}

	input = strings.TrimSpace(input)

	retInput, err := convertType(input, t)
	if err != nil {
		fmt.Println(err)
		return Call(msg, t)
	}

	return retInput
}

func convertType(input string, t string) (interface{}, error) {
	switch t {
	case "float64":
		retInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return nil, err
		}
		return retInput, nil
	case "float32":
		retInput, err := strconv.ParseFloat(input, 32)
		if err != nil {
			return nil, err
		}
		return retInput, nil
	case "int":
		retInput, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		return retInput, nil
	default:
		return input, nil
	}
}

