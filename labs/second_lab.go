package labs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"
)


type SLConfig struct {
	DocumentRoot string
}

type SecondLab struct {
	Config SLConfig
}

func (sl *SecondLab) ThirdTask() {
	var filesDir = sl.Config.DocumentRoot + string(os.PathSeparator) + "filesDir"
	filesMap := getTxtFilesInfo(filesDir)
	for name, linesCnt := range filesMap {
		fmt.Println(name, ":", linesCnt)
	}
}

func (sl *SecondLab) SeventhTask() {
	var filePath = sl.Config.DocumentRoot + string(os.PathSeparator) + "filesDir" + string(os.PathSeparator) + "Sentences.txt"
	wordsMap := resolveSentenceWords(filePath)
	for word, sentences := range wordsMap {
		fmt.Printf("%s: %v\n", word, sentences)
	}
}

func getTxtFilesInfo(dir string) map[string]int {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal("Directory", dir, "not found!")
	}

	txtFilesMap := make(map[string]int)
	for _, file := range files {
		filePath := dir + string(os.PathSeparator) + file.Name()
		if file.IsDir() || filepath.Ext(filePath) != ".txt" {
			continue
		}
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal("Failed to open" + filePath)
		}
		scanner := bufio.NewScanner(file)
		txtFilesMap[file.Name()] = 0
		for scanner.Scan() {
			txtFilesMap[file.Name()] += 1
		}
		_ = file.Close()
	}

	return txtFilesMap
}

func resolveSentenceWords(fileName string) map[string][]int {
	wordsMap := make(map[string][]int)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file" + fileName)
	}
	defer file.Close()

	word := ""
	sentenceCounter := 1
	buffer := make([]byte, 80)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err.Error())
		}
		fragment := string(buffer[:n])
		for _, symbol := range fragment {
			if unicode.IsLetter(rune(symbol)) {
				word += string(symbol)
			} else {
				if string(symbol) == "." {
					sentenceCounter++
				} else if string(symbol) == " " {
					if wordsMap[word] == nil {
						wordsMap[word] = []int{sentenceCounter}
					} else {
						wordsMap[word] = append(wordsMap[word], sentenceCounter)
					}
					word = ""
				}
			}
		}
	}


	return wordsMap
}