# colstr

Simple command line utility to colorize strings based on delimiters contained within them

[![ci](https://github.com/jamieweavis/colstr/actions/workflows/ci.yml/badge.svg)](https://github.com/jamieweavis/colstr/actions)
[![downloads](https://img.shields.io/github/downloads/jamieweavis/colstr/total)](https://github.com/jamieweavis/colstr/releases)
[![version](https://img.shields.io/github/v/release/jamieweavis/colstr)](https://github.com/jamieweavis/colstr/releases)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jamieweavis/colstr/blob/main/LICENSE)

> *pronounced "col-ster"*

## Installation

Install via [Homebrew](https://brew.sh) (currently macOS only):

```sh
brew install jamieweavis/tap/colstr
```

## Usage

By default `colstr` will colorize a string based on the following delimiters `. / ? = & # :`

### URL

```sh
colstr 'https://github.com/jamieweavis/colstr?hello=world#test'
```

<img src=".github/url.svg" alt="URL">

### JWT

```sh
colstr 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoZWxsbyI6IndvcmxkIn0.eH9qoMvdv12LsZ3Og_K20no8uiBQFuJg6k6A7O8l06U'
```

<img src=".github/jwt.svg" alt="JWT">

### Custom Delimiters

You can specify custom delimiters by passing a second argument to `colstr`


#### UUIDs

```sh
colstr '123e4567-e89b-12d3-a456-426614174000' '-'
```

<img src=".github/uuid.svg" alt="UUID">

#### Other

```sh
colstr 'foo$bar%baz' '$%'
```

<img src=".github/custom.svg" alt="Custom">

### Help

Print usage and package information:

```sh
colstr
```

<img src=".github/help.svg" alt="Help">

## Building Locally

### Compiling

Build a binary:

```sh
go build
```

Run the binary:

```sh
./colstr
```

### Install

Install the binary to your system:

```sh
mv colstr /usr/local/bin
```

Uninstall the binary from your system:

```sh
rm /usr/local/bin/colstr
```

## Disclaimer

I'm not a Go developer, this is just for fun!

## Built With

- [Go](https://github.com/golang/go)

## Related

- [jamieweavis/homebrew-tap](https://github.com/jamieweavis/homebrew-tap) - Homebrew tap for my brew formulae
- [Carbon](https://carbon.now.sh) - For terminal screenshots (with [Rosé Pine](https://github.com/rose-pine/rose-pine-theme) theme & [JetBrains Mono](https://github.com/JetBrains/JetBrainsMono) font)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
