package godump

import (
	"bytes"
	"fmt"
	"io"
)

//DumpBytes returns binary data from io.Reader
func DumpBytes(r io.Reader, name string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	b := make([]byte, 1)
	var err error
	sep := fmt.Sprintf("var %s = []byte{", name)
	for true {
		if _, err = r.Read(b); err != nil {
			break
		}
		fmt.Fprintf(buf, "%s%#02x", sep, b)
		sep = ", "
	}
	fmt.Fprintln(buf, "}")
	return buf, nil
}
