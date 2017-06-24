package main

import (
	"os"
	"strconv"
	"github.com/parlihpb/neuralnet/models"
	"fmt"
	"log"
	"bufio"
	"strings"
)

//maybe could use a defer to do the backtracing? maybe not as it would need to wait until the full trace through to get the value upon which to calculate error in the first nodes.
func main() {
	fmt.Println("Hello, world")

	networkLength, _ := strconv.Atoi(os.Args[1])
	networkHeight, _ := strconv.Atoi(os.Args[2])
	fileLocation := os.Args[3]
	file, error := os.Open(fileLocation)

	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	for lineScanner.Scan() {

		text := lineScanner.Text()
		fields := strings.Fields(text)
		fmt.Println(fields)
		fmt.Println("break")
	}

	//initializeNetwork(networkHeight, networkLength)
	models.NewNetwork(networkHeight, networkLength, 5, 1)

}
