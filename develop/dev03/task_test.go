package main

import (
	"os"
	"os/exec"
	"testing"
)

type addTest struct {
	arg1 string
}

var addTests = []addTest{
	addTest{"-n"},
	addTest{"-r"},
	addTest{"-u"},
	//addTest{"-k"},
}

func TestSort(t *testing.T) {

	file1, _ := os.Create("a")
	defer file1.Close()

	file2, _ := os.Create("b")
	defer file2.Close()

	for i, test := range addTests {
		t.Log("TEST", i+1, " -- sort", test.arg1, "file.txt")
		a, _ := exec.Command("./task", test.arg1, "file.txt").CombinedOutput()
		file1.WriteString(string(a))

		b, _ := exec.Command("sort", test.arg1, "file.txt").CombinedOutput()
		file2.WriteString(string(b))

		res, err := exec.Command("diff", "a", "b").CombinedOutput()
		if len(res) != 0 {
			t.Errorf("Error %s", string(res))
			t.Errorf(err.Error())
		}
		//fmt.Print(string(res))

	}

}
