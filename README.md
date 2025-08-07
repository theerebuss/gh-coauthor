# gh-coauthor

A GitHub CLI extension for adding co-authors to your commits.

Interactively shows repository authors to choose from if no argument is provided.

![example of using the coauthor extension in interactive mode. when passing no username, a list of the repository's collaborators is shown and the user interactively selects a random collaborator. we then see that the commit message has been edited to include the user's name, email and the "Co-authored-by" tag.](./assets/interactive.gif)

## Installation

```shell
gh extension install theerebuss/gh-coauthor
```

## Usage

```shell
gh coauthor [username] [flags]

Flags:
  -h, --help   help for gh coauthor
```

## Examples

### Specific co-author

```shell
gh coauthor ashtom
```

![example of using the coauthor extension with a specific username. when passed a username, we see that the commit message has been edited to include the user's name, email and the "Co-authored-by" tag.](./assets/specific.gif)

### Interactive selection

```shell
gh coauthor
```

![example of using the coauthor extension in interactive mode. when passing no username, a list of the repository's collaborators is shown and the user interactively selects a random collaborator. we then see that the commit message has been edited to include the user's name, email and the "Co-authored-by" tag.](./assets/interactive.gif)

## Development

```shell
go build && gh coauthor
```
