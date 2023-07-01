
# gopandoc

Go API for using the [pandoc](https://pandoc.org/) document converter.

## Install

Follow th installation instructions on [pandoc.org/installing.html](https://pandoc.org/installing.html).
Then, in your project folder, run

```sh
go get github.com/alexcoder04/gopandoc
```

## Usage

Use pandoc in your go code as follows:

```go
package main

import (
  "github.com/alexcoder04/gopandoc"
)

func main() {
  // convert markdown to docx
  var inBytes []byte = ...
  outBytes, err := gopandoc.BytesToBytes(inBytes, "markdown", "docx")
  if err != nil {
    ...
  }
  ...
}
```

There is also a `BytesToFile(input []byte, inFormat string, outFormat string, outputFile string)` function for directly writing the results into a file.
