package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// fmt.Println("hello")
	// exportStdout()
	// readFile("test")
	readLine()
}

func exportStdout() {
	w := bufio.NewWriter(os.Stdout)
	w.WriteString("hello")
	w.Flush()
}

func readFile(fileName string) {
	path := "D:/project/vscode/blog/blogs/每日随笔/2022-03-02_note.md"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(content)

}

func readLine() {
	path := "D:/project/vscode/blog/blogs/每日随笔/2022-03-02_note.md"
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("size = ", size)

	for {
		buf := bufio.NewReader(file)
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		if err != nil {
			if err == io.EOF {
				fmt.Println("file read end")
				break
			} else {
				fmt.Println("file read err")
				return
			}
		}
	}

}

func writeFile(path string, line string) {

}
