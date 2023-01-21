# Code Name Generator

A Go package to generate a three-word code name for multi-purpose use. The package relies on
codenamegenerator.com for generation as opposed to lengthy slices within the package.

The output is as follows: `<string>-<string>-<string>`.

## Badges

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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

import "github.com/jmvbxx/codenamegenerator"

func main() {
	codenamegenerator.NameGenerate()
}
```

## Example output
`brussels-ankara-yowie`