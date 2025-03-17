# Se**q**uential Co**p**y

A tool to recursively copy files and directories so that timestamps are in order of file names.

## Features

- Copies directories first, then files
- Maintains sequential timestamps
- Displays progress with counters

## Usage

```
qp <source directory> <destination directory>
```

## Build

Make sure you have installed [Go](https://go.dev/) and [Just](https://just.systems/).

Execute:

```
just build
```
