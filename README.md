# Code Name Generator

A Go package to generate a three-word code name for multi-purpose use. The package relies on
[codenamegenerator.com](https://codenamegenerator.com) for generation as opposed to lengthy slices within the package.

The output is as follows: `<string>-<string>-<string>`.

## Badges

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/jmvbxx/codenamegenerator)](https://goreportcard.com/report/github.com/jmvbxx/codenamegenerator)

## Install

```bash
go get github.com/jmvbxx/codenamegenerator
```

## Update

```bash
go get -u github.com/jmvbxx/codenamegenerator
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/jmvbxx/codenamegenerator"
)

func main() {
	cn, err := codenamegenerator.NameGenerate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cn)
}
```

## Example output
`brussels-ankara-yowie`
