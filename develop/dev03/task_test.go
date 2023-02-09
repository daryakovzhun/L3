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
}

func testK(t *testing.T, i *int, file1, file2 *os.File, column string) {
	t.Log("TEST", *i, " -- sort", "k", "file.txt")
	a, _ := exec.Command("./task", "-k", column, "file.txt").CombinedOutput()
	file1.WriteString(string(a))

	b, _ := exec.Command("sort", "-k", column, "file.txt").CombinedOutput()
	file2.WriteString(string(b))

	res, err := exec.Command("diff", "a", "b").CombinedOutput()
	if len(res) != 0 {
		t.Errorf("Error %s", string(res))
		t.Errorf(err.Error())
	} else {
		t.Log("OK")
	}
	*i += 1
}

func TestSort(t *testing.T) {

	file1, _ := os.Create("a")
	defer file1.Close()

	file2, _ := os.Create("b")
	defer file2.Close()

	i := 1
	for _, test := range addTests {

		t.Log("TEST", i, " -- sort", test.arg1, "file.txt")
		a, _ := exec.Command("./task", test.arg1, "file.txt").CombinedOutput()
		file1.WriteString(string(a))

		b, _ := exec.Command("sort", test.arg1, "file.txt").CombinedOutput()
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

	testK(t, &i, file1, file2, "1")
	testK(t, &i, file1, file2, "2")
	testK(t, &i, file1, file2, "3")

}
