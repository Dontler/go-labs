package labs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const spaceSymbol = 32
const dotSymbol = 46
const lowCaseA = 65
const lowCaseZ = 90
const upperCaseA = 97
const upperCaseZ = 122

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
		for i := 0; i < n; i++ {
			symbol := buffer[i]
			if symbol >= lowCaseA && symbol <= lowCaseZ || symbol >= upperCaseA && symbol <= upperCaseZ {
				word += string(symbol)
			} else {
				if symbol == dotSymbol {
					sentenceCounter++
				} else if symbol == spaceSymbol {
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