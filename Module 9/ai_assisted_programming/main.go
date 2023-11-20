package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, AI Assisted Programming!")

	// Example: Perform some calculations
	result := calculateSum(5, 10)
	fmt.Println("Sum:", result)

	// Example: Use advanced math functions
	squareRoot := calculateSquareRoot(25)
	fmt.Println("Square Root:", squareRoot)

	// Example: Use advanced string manipulation
	str := "Hello, World!"
	uppercase := convertToUppercase(str)
	fmt.Println("Uppercase:", uppercase)

	// Example: Use advanced data structures
	queue := NewQueue()
	queue.Enqueue("Apple")
	queue.Enqueue("Banana")
	queue.Enqueue("Cherry")

	fmt.Println("Queue Size:", queue.Size())

	item := queue.Dequeue()
	fmt.Println("Dequeued Item:", item)

	// Example: Use advanced file operations
	fileName := "example.txt"
	content := "This is an example file."
	err := writeToFile(fileName, content)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File written successfully.")
	}
}

func calculateSum(a, b int) int {
	return a + b
}

func calculateSquareRoot(n float64) float64 {
	return math.Sqrt(n)
}

func convertToUppercase(str string) string {
	return strings.ToUpper(str)
}

type Queue struct {
	items []string
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
	if len(q.items) == 0 {
		return ""
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Size() int {
	return len(q.items)
}

func writeToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
