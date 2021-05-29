//Package data file allows to read data from files to an array
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([9]float64, error) {
	fmt.printf("Changed")
	var numbers [9]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
