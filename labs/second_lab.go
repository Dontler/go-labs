package labs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
		defer file.Close()
		scanner := bufio.NewScanner(file)
		txtFilesMap[file.Name()] = 0
		for scanner.Scan() {
			txtFilesMap[file.Name()] += 1
		}
	}

	return txtFilesMap
}