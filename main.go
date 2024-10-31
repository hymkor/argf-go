package argf

import (
	"fmt"
	"io"
	"os"
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
	f := make([]io.Reader, len(filenames))
	for i, fname := range filenames {
		f[i] = &autoOpenFile{Name: fname}
	}
	return io.MultiReader(f...)
}
