package hook

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Commands struct {
	file io.Writer
}

func NewCommands() *Commands {
	return &Commands{file: os.Stdout}
}

func NewTestCommands() (*bytes.Buffer, *Commands) {
	buffer := bytes.NewBufferString("")
	return buffer, &Commands{file: buffer}
}

func (h *Commands) SetEnv(name, value string) {
	fmt.Fprintf(h.file, "export %s=\"%s\"\n", name, value)
}

func (h *Commands) UnsetEnv(name string) {
	fmt.Fprintf(h.file, "unset %s\n", name)
}
