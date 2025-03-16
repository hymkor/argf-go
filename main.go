package argf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type autoOpenFile struct {
	Name string
	fd   *os.File
}

func (f *autoOpenFile) Read(buffer []byte) (int, error) {
	if f.fd == nil {
		var err error
		f.fd, err = os.Open(f.Name)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", f.Name, err)
		}
	}
	n, err := f.fd.Read(buffer)
	if err != nil {
		f.fd.Close()
		if err != io.EOF {
			err = fmt.Errorf("%s: %w", f.Name, err)
		}
	}
	return n, err
}

func New(filenames []string) io.Reader {
	if len(filenames) <= 0 {
		return os.Stdin
	}
	f := make([]io.Reader, 0, len(filenames))
	for _, fname := range filenames {
		if matches, err := filepath.Glob(fname); err == nil {
			for _, m := range matches {
				f = append(f, &autoOpenFile{Name: m})
			}
		} else {
			f = append(f, &autoOpenFile{Name: fname})
		}
	}
	return io.MultiReader(f...)
}
