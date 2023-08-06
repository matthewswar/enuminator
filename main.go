package main

import (
	_ "embed"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/matthewswar/enuminator/generator"
	"github.com/matthewswar/enuminator/parser"
	"github.com/matthewswar/enuminator/version"
	"github.com/pkg/errors"
)

type VersionCmd struct{}

func (cmd *VersionCmd) Run(ctx *kong.Context) error {
	fmt.Println(version.Version())
	return nil
}

type GenerateCmd struct {
	Name        string `short:"n" required:"" help:"Name of the enumeration type."`
	PackagePath string `short:"p" default:"." help:"Path to the package."`
	FilePrefix  string `short:"f" default:"enum" help:"The prefix for the generated file names."`
	Header      string `short:"h" help:"The header to the output files"`
}

func (cmd *GenerateCmd) Run(ctx *kong.Context) error {
	packageName, enumData, err := parser.ExtractEnumeration(cmd.PackagePath, cmd.Name)
	if err != nil {
		return errors.Wrap(err, "could not extract enum")
	}

	var options []generator.OptionApplier

	if cmd.FilePrefix != "" {
		options = append(options, generator.WithFilePrefix(cmd.FilePrefix))
	}

	if cmd.Header != "" {
		options = append(options, generator.WithHeader(cmd.Header))
	}

	err = generator.GenerateEnum(cmd.PackagePath, packageName, cmd.Name, enumData, options...)
	if err != nil {
		return errors.Wrap(err, "unable to generate enum files")
	}

	return nil
}

var cli struct {
	Generate GenerateCmd `cmd:"" default:"withargs" help:"Generates enums for the specified type."`
	Version  VersionCmd  `cmd:"" help:"Prints the version of this tool."`
}

func main() {
	ctx := kong.Parse(&cli)
	ctx.FatalIfErrorf(ctx.Run())
}
