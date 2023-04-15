package gopandoc

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

func BytesToFile(input []byte, inFormat string, outFormat string, outputFile string) error {
	output, err := BytesToBytes(input, inFormat, outFormat)
	if err != nil {
		return err
	}
	return os.WriteFile(outputFile, output, 0600)
}

func BytesToBytes(input []byte, inFormat string, outFormat string) ([]byte, error) {
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
