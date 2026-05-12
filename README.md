# Figurine

[![PkgGoDev](https://pkg.go.dev/badge/github.com/arsham/figurine/v2)](https://pkg.go.dev/github.com/arsham/figurine/v2)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/arsham/figurine)
[![Go Report Card](https://goreportcard.com/badge/github.com/arsham/figurine/v2)](https://goreportcard.com/report/github.com/arsham/figurine/v2)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Print your name in style

![Screenshot](/docs/figurine.png?raw=true "Rainbow")

### Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [See Also](#see-also)
4. [License](#license)

## Installation

You can download the latest binary from
[here](https://github.com/arsham/figurine/releases), or you can compile from
source:

```bash
go install github.com/arsham/figurine/v2@latest
```

## Usage

Every time the application is called, it chooses a random font for rendering the
message. Pass the message you want to decorate as arguments.

```bash
figurine Some Text
```

You can print available fonts:

```bash
figurine -l
figurine -l -s
figurine -ls Sample Text
```

To set a font:

```bash
figurine -f "Poison.flf" Some Text
```

To get a list of available arguments:

```bash
figurine -h
```

This application is very light weight, so feel free to add it to your
.zshrc/.bashrc file, so each time you open a new shell it shows you a nice
message.

## Github release workflow

How to use it

1. Commit and push the workflow file.
2. reate and push a tag:

```bash
git tag v2.1.0
git push origin v2.1.0
```

The workflow will automatically:

* Build for Linux (amd64, arm, arm64)
* Build for macOS (amd64)
* Build for Windows (amd64)
* Create .tar.gz / .zip archives
* Generate CHECKSUM file with SHA256 hashes
* Create a GitHub Release with all assets

### Optional improvements

Using Goreleaser (more powerful, recommended for Go projects):
If a more feature-rich solution is preferred, you can replace the Makefile build step with Goreleaser.

## See Also

See also [Rainbow][rainbow], which is the library that colours the output.

## License

Use of this source code is governed by the Apache 2.0 license. License that can
be found in the [LICENSE](./LICENSE) file.

[rainbow]: https://github.com/arsham/rainbow
