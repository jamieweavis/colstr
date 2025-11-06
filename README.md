# colstr

Simple command line utility to colorize strings based on delimiters contained within them

[![ci](https://github.com/jamieweavis/colstr/actions/workflows/ci.yml/badge.svg)](https://github.com/jamieweavis/colstr/actions)
[![downloads](https://img.shields.io/github/downloads/jamieweavis/colstr/total)](https://github.com/jamieweavis/colstr/releases)
[![version](https://img.shields.io/github/v/release/jamieweavis/colstr)](https://github.com/jamieweavis/colstr/releases)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jamieweavis/colstr/blob/main/LICENSE)

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

## Development

### Prerequisites

- [Go](https://github.com/golang/go) (>=1.24.x)

### Getting Started

Clone the repository:

```sh
git clone https://github.com/jamieweavis/colstr.git

cd colstr
```

Run the application:
```bash
make run
```

Compile the application binary:
```bash
make build
```

Package the binary into a tar.gz archive:
```bash
make package
```

Print the check-sum of the packaged application:
```bash
make checksum
```

Delete build artifacts:
```bash
make clean
```
