package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:] // Exclude the program name itself

	if len(args) > 0 {
		day := args[0]

		wd, _ := os.Getwd()
		dirPath := wd + "/day" + day

		if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
			fmt.Println("==== Found", dirPath, "removing first")
			os.RemoveAll(dirPath)
		}

		// Create folder
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}

		fmt.Println("==== Creating files at", dirPath)
		src, _ := os.Create(dirPath + "/day" + day + ".go")
		testSrc, _ := os.Create(dirPath + "/day" + day + "_test.go")
		goMod, _ := os.Create(dirPath + "/go.mod")
		inputFile, _ := os.Create(dirPath + "/input.txt")
		defer src.Close()
		defer testSrc.Close()
		defer goMod.Close()
		defer inputFile.Close()

		src.WriteString("package day" + day)
		testSrc.WriteString("package day" + day)
		goMod.WriteString("module besharati.se/advent-of-code-2023/day" + day + "\n\ngo 1.21.3\n")

		fmt.Println("==== Adding day" + day + " to go.work")
		cmd := exec.Command("go", "work", "use", "day"+day)
		cmd.Run()
	} else {
		fmt.Println("No day argument provided.")
	}
}
