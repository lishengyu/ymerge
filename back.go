package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(path string) (bool, []int) {
	var list []int

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return false, list
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, gKey) {
			list = append(list, count)
		}
	}

	if len(list) == 0 {
		return false, list
	} else {
		return true, list
	}
}

func genNewLine(add string, level int) string {
	switch level {
	case 0:
		return add
	case 1:
		return "    " + add
	case 2:
		return "        " + add
	case 11:
		return "      " + add
	case 12:
		return "          " + add
	case 13:
		return "              " + add
	default:
		return add
	}
}

func matchReplace(line int, lists []int) bool {
	for _, list := range lists {
		if line == list {
			return true
		}
	}
	return false
}

func replace(path string, lists []int, newline string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tmpPath := path + ".tmp"
	wfile, err := os.Create(tmpPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		line := scanner.Text()

		if flag := matchReplace(count, lists); flag {
			fmt.Printf("replace [%s] to [%s]\n", line, newline)
			wfile.WriteString(newline + "\n")
		} else {
			wfile.WriteString(line + "\n")
		}
	}

	os.Rename(tmpPath, path)
}

func add(path, newline string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tmpPath := path + ".tmp"
	wfile, err := os.Create(tmpPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		line := scanner.Text()

		if strings.Contains(line, gMatch) {
			fmt.Printf("add [%s]\n", newline)
			if gOffset == "up" {
				wfile.WriteString(newline + "\n")
				wfile.WriteString(line + "\n")
			} else {
				wfile.WriteString(line + "\n")
				wfile.WriteString(newline + "\n")
			}
		} else {
			wfile.WriteString(line + "\n")
		}
	}

	os.Rename(tmpPath, path)
}

var (
	gFile   string //修改文件
	gKey    string //查找关键词
	gValue  string //添加字段
	gMatch  string //匹配特征，在特征上或下添加字段
	gOffset string //添加字段位置，上或下
	gLevel  int    //第几级字段
)

func main() {
	flag.StringVar(&gFile, "f", "", "需要修改的配置文件")
	flag.StringVar(&gKey, "k", "", "查找关键词")
	flag.StringVar(&gAdd, "v", "", "添加字段")
	flag.StringVar(&gMatch, "m", "", "匹配特征，在特征上或下添加字段")
	flag.StringVar(&gOffset, "o", "", "添加字段位置，上或下")
	flag.IntVar(&gLevel, "l", 0, "第几级字段")
	flag.Parse()

	if gFile == "" || gKey == "" || gAdd == "" || gMatch == "" || gOffset == "" {
		//flag.Usage()
		//os.Exit(-1)
	}

	newLine := genNewLine(gAdd, gLevel)
	if flag, lists := check(gFile); flag {
		replace(gFile, lists, newLine)
	} else {
		add(gFile, newLine)
	}

	return
}
