# Enuminator

This is a tool that generates case insensitive enumerations. Only string enumerations are supported.

## Usage

```
$ ./enuminator --help
Usage: enuminator <command>

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  generate --name=STRING
    Generates enums for the specified type.

  version
    Prints the version of this tool.

Run "enuminator <command> --help" for more information on a command.
```

### Go Generate

This tool is best suited for Go's [generate](https://go.dev/blog/generate) functionality.

```golang
//go:generate enuminator --name Color
type Color string

const (
    ColorRed Color   = "red"
    ColorBlue Color  = "blue"
    ColorGreen Color = "green"
)
```
