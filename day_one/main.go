package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var array1, array2 []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		num1, err1 := strconv.Atoi(fields[0])
		num2, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Erro ao converter número:", err1, err2)
			continue
		}

		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura do arquivo:", err)
		return
	}

	fmt.Println("Array 1:", array1)
	fmt.Println("Array 2:", array2)

	result := array1[0] + array2[0]
	fmt.Println("Soma:", result)
}
