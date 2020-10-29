package osx

import (
	"fmt"
)

func main() {
}

func listtest() {
	dir := "/home/eco/Desktop/ecoshubjin"
	fmt.Println(dir)
	// filepath.
	list, err := list(dir)
	fmt.Println(err)
	fmt.Println(list)
}

func moveTest() {
	src := getHome() + "/Desktop/jin"
	dest := getHome() + "/Desktop/ecoshubjin"
	err := moveFolder(src, dest)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func testLineReader() {
	dir := "a.text"
	lr, err := NewLineReader(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, hasNext := lr.Next()
	for hasNext {
		fmt.Println(string(data))
		data, hasNext = lr.Next()
	}

	lr.Refresh()

	data, hasNext = lr.Next()
	for hasNext {
		fmt.Println(string(data))
		data, hasNext = lr.Next()
	}
}

func testFileBatcher() {
	dir := "a.text"
	batcher, err := NewFileBatcher(dir, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, hasNext := batcher.Next()
	for hasNext {
		fmt.Println(string(data))
		data, hasNext = batcher.Next()
	}

	batcher.Refresh()

	data, hasNext = batcher.Next()
	for hasNext {
		fmt.Println(string(data))
		data, hasNext = batcher.Next()
	}
}

func testByteBatcher() {
	dir := "a.text"
	buffer, err := ReadFile(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	batcher, err := NewByteBatcher(buffer, 7)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, hasNext := batcher.Next()
	for hasNext {
		fmt.Print(string(data))
		data, hasNext = batcher.Next()
	}
}
