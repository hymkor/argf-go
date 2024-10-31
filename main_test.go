package argf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew1(t *testing.T) {
	contents := []string{
		"ahaha", "ufuuf\nohoho\n", "gohoho\n",
	}
	tempDir := t.TempDir()
	filenames := make([]string, len(contents))
	for i, content := range contents {
		name := fmt.Sprintf("%d.txt", i+1)
		filenames[i] = filepath.Join(tempDir, name)
		// println(filenames[i])
		err := os.WriteFile(filenames[i], []byte(content), 0777)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	f := New(filenames)
	result, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err.Error())
	}
	expect := []byte(strings.Join(contents, ""))
	if !bytes.Equal(result, expect) {
		t.Fatal("not equals")
	}
}

func TestNew0(t *testing.T) {
	f := New([]string{})
	if _f, ok := f.(*os.File); !ok || _f != os.Stdin {
		t.Fatal("not stdin for no filenames")
	}
}
