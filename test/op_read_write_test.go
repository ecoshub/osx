package osx

import (
	"osx"
	"strings"
	"testing"
)

const (
	writerDirectory        string = "../test_assets/writerTest.txt"
	testfileDirectory      string = "../test_assets/testfile.txt"
	testfileFalseDirectory string = "../test_assets/testfile.json"
	testfileContent        string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
incididunt ut labore et dolore magna aliqua. 
Lectus urna duis convallis convallis tellus id. 
Urna condimentum mattis pellentesque id nibh tortor id aliquet lectus. 
Arcu vitae elementum curabitur vitae nunc sed velit dignissim. 
Magna sit amet purus gravida quis blandit turpis. 
Eget dolor morbi non arcu risus quis varius quam. 
In hac habitasse platea dictumst. 
Sagittis aliquam malesuada bibendum arcu vitae. 
Enim praesent elementum facilisis leo vel fringilla. 
Faucibus turpis in eu mi bibendum neque egestas congue quisque. 
Proin nibh nisl condimentum id venenatis a. 
Donec adipiscing tristique risus nec. 
Turpis tincidunt id aliquet risus feugiat in ante metus. 
Fusce ut placerat orci nulla pellentesque dignissim enim. 
Quis varius quam quisque id diam vel quam. 
Sodales ut etiam sit amet. 
Diam vulputate ut pharetra sit amet aliquam id diam. 
Justo nec ultrices dui sapien eget. 
Ipsum a arcu cursus vitae congue mauris.
`
)

func TestRead(t *testing.T) {
	var err error
	var file []byte
	_, err = osx.ReadFile(testfileFalseDirectory)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			t.FailNow()
		}
	}
	file, err = osx.ReadFile(testfileDirectory)
	if err != nil {
		t.Error(err)
		return
	}
	if string(file) != testfileContent {
		t.FailNow()
	}
}

func TestWrite(t *testing.T) {
	var err error
	var file []byte
	err = osx.WriteFile(writerDirectory, []byte(testfileContent))
	if err != nil {
		t.Error(err)
		return
	}
	file, err = osx.ReadFile(testfileDirectory)
	if err != nil {
		t.Error(err)
		return
	}
	if string(file) != testfileContent {
		t.FailNow()
	}
}
