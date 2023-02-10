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
	addTest{"-c"},
	addTest{"-i"},
	addTest{"-v"},
	addTest{"-F"},
	addTest{"-n"},
}

type addTest2 struct {
	arg1, arg2 string
}

var addTests2 = []addTest2{
	addTest2{"-A", "5"},
	addTest2{"-A", "1"},
	addTest2{"-B", "3"},
	addTest2{"-B", "2"},
	addTest2{"-C", "6"},
	addTest2{"-C", "4"},
}

func TestGrep(t *testing.T) {

	file1, _ := os.Create("a")
	defer file1.Close()

	file2, _ := os.Create("b")
	defer file2.Close()

	i := 1
	filename := "task.go"
	for _, test := range addTests {

		t.Log("TEST", i, " -- grep", test.arg1, filename)
		a, _ := exec.Command("./task", test.arg1, filename).CombinedOutput()
		file1.WriteString(string(a))

		b, _ := exec.Command("grep", test.arg1, filename).CombinedOutput()
		file2.WriteString(string(b))

		res, err := exec.Command("diff", "a", "b").CombinedOutput()
		if len(res) != 0 {
			t.Errorf("Error %s", string(res))
			t.Errorf(err.Error())
		} else {
			t.Log("OK")
		}
		i++
	}

	for _, test := range addTests2 {

		t.Log("TEST", i, " -- grep", test.arg1, test.arg2, filename)
		a, _ := exec.Command("./task", test.arg1, test.arg2, filename).CombinedOutput()
		file1.WriteString(string(a))

		b, _ := exec.Command("grep", test.arg1, test.arg2, filename).CombinedOutput()
		file2.WriteString(string(b))

		res, err := exec.Command("diff", "a", "b").CombinedOutput()
		if len(res) != 0 {
			t.Errorf("Error %s", string(res))
			t.Errorf(err.Error())
		} else {
			t.Log("OK")
		}
		i++
	}

}
