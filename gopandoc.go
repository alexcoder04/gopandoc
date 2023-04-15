package gopandoc

import (
	"bytes"
	"errors"
	"os/exec"
)

func BytesToBytes(input []byte, output []byte, inFormat string, outFormat string) ([]byte, error) {
	_, err := exec.LookPath("pandoc")
	if err != nil {
		return []byte{}, errors.New("pandoc not installed")
	}

	inPFormat, err := toPandocInputFormat(inFormat)
	if err != nil {
		return []byte{}, err
	}

	outPFormat, err := toPandocOutputFormat(outFormat)
	if err != nil {
		return []byte{}, err
	}

	cmd := exec.Command("pandoc", "-f", inPFormat, "-t", outPFormat)
	cmd.Stdin = bytes.NewBuffer(input)
	return cmd.Output()
}
