package scandir

//import "fmt"
import "path/filepath"
import "os"
import "github.com/golang/glog"

type FileFoundInterface interface {
	Process(fname string) error
	LookFor() string
}

type ScanDirError struct {
	Msg  string
	Code int
}

func (e ScanDirError) Error() string { return e.Msg }

type allPaths struct {
	paths []string
}

func (a *allPaths) Add(p string) { a.paths = append(a.paths, p) }

// ScanDir recursively scans a directory tree using the FileFoundInterface
// to look for files conforming to the search string specified by the function
// LookFor(), calling Process(fname string) when found.  Returns a []string containing
// the absolute path to the file.
func ScanDir(f FileFoundInterface, root string) error {
	fname := f.LookFor()
	if fname == "" {
		glog.Fatalf("No file name to look for.")
		return ScanDirError{"No file name to look for", 100}
	}
	p := &allPaths{}
	fnWalk := makeWalkFunc(p, fname)
	filepath.Walk(root, fnWalk)

	for i := 0; i < len(p.paths); i++ {
		f.Process(p.paths[i])
	}
	return nil
}

func makeWalkFunc(p *allPaths, fname string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		_, filename := filepath.Split(path)
		if filename == fname {
			p.Add(path)
		}
		return nil
	}
}
