package tsdoc

import "testing"

func Test(t *testing.T) {
	doc, err := Get(".", false)
	if err != nil {
		t.Fatal(err)
	}
	doc2, err := Get("./", false)
	if err != nil {
		t.Fatal(err)
	}
	if doc != doc2 {
		t.Fatal("Documentation is not same")
	}
}

func TestWrong(t *testing.T) {
	var err error
	_, err = Get("No folder exist", false)
	if err == nil {
		t.Fatal("folder is not exist")
	}
	_, err = Get("README.md", false)
	if err == nil {
		t.Fatal("cannot find in file")
	}
}
