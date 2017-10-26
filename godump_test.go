package godump

import (
	"bytes"
	"io"
	"os"
)

func ExampleDumpBytes() {
	dump, err := DumpBytes(bytes.NewBufferString("hello world"), "foobar")
	if err != nil {
		return
	}
	io.Copy(os.Stdout, dump)
	// Output:
	// var foobar = []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}
}
