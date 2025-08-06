# gh-coauthor

A GitHub CLI extension for adding co-authors to your commits.

Interactively shows repository authors to choose from if no argument is provided.

## Development

```shell
go build && gh coauthor
```

## Installation

```shell
gh extension install .; gh coauthor
```

## Usage

```shell
gh coauthor [username] [flags]

Flags:
  -h, --help   help for gh coauthor
```

## Examples

```shell
# Interactive
gh coauthor

# Specify a username
gh coauthor johndoe
```
