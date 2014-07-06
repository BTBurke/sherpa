package scandir

import "testing"
import "io/ioutil"
import "os"
import "fmt"
import "errors"

type TestingContainer struct {
	paths []string
}

func (t *TestingContainer) LookFor() string { return ".sherparc" }
func (t *TestingContainer) Process(s string) error {
	t.paths = append(t.paths, s)
	fmt.Println("paths: ", t.paths)
	return nil
}

func setup(start string, putFile bool) (string, error) {
	if start == "" {
		start = "/tmp"
	}
	dName, err := ioutil.TempDir(start, "sherpa")
	if err != nil {
		return "", err
	}
	fname := dName + "/.sherparc"

	if putFile {
		f, err := os.Create(fname)
		if err != nil {
			return "", err
		}
		defer f.Close()

		f.WriteString("test")
	}
	return dName, nil
}

func recursiveSetup(recur int, start string) ([]string, error) {
	dirNames := make([]string, recur)
	var name string
	for i := 0; i < recur; i++ {
		if i == 0 {
			name, _ = setup(start, false)
		} else {
			name, _ = setup(name, true)
		}
		dirNames[i] = name
	}
	return dirNames, nil
}

func cleanup(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return errors.New("Failed to remove testing directory.")
	}
	return nil
}

func TestWalk2Level(t *testing.T) {
	top, err := recursiveSetup(2, "")
	if err != nil {
		t.Fatalf("Failed on setup")
	}
	n := TestingContainer{}
	scanerr := ScanDir(&n, top[0])
	if scanerr != nil {
		t.Fatalf("Failed on ScanErr")
	}

	fmt.Println("n: ", n)
	expected := top[1] + "/" + n.LookFor()
	if len(n.paths) == 0 {
		t.Fatalf("Failed on length 0 path. Should be 1.")
	}
	if n.paths[0] != expected {
		t.Fatalf("Expected: %s, Got: %s", expected, n.paths[0])
	}
	cleanup(top[0])
}
