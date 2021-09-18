package main

import (
	"Lab1/labs"
	"fmt"
)

func main() {
	config := labs.Config{N: 16, A: 10, B: 15}
	firstLab := labs.FirstLab{Config: config}

	fmt.Println("Task 3:")
	firstLab.ThirdTask()
	fmt.Println("Task 7:")
	firstLab.SeventhTask()
	fmt.Println("Task 10:")
	firstLab.TenthTask()
	fmt.Println("Task 11:")
	firstLab.EleventhTask()

	fmt.Println("Second lab")

	slConfig := labs.SLConfig{DocumentRoot: "C:\\Users\\sofya\\GolandProjects\\GoLabs\\Lab1"}
	secondLab := labs.SecondLab{Config: slConfig}

	fmt.Println("Tsk 3:")
	secondLab.ThirdTask()
}