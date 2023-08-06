package generator

import (
	"fmt"
	"strings"
)

type Options struct {
	FilePrefix   string
	Header       string
	OutputFolder string
	ImportPath   string
}

type OptionApplier interface {
	Apply(options *Options)
}

type filePrefixOption string

func (opt filePrefixOption) Apply(options *Options) {
	options.FilePrefix = string(opt)
}

func WithFilePrefix(prefix string) OptionApplier {
	return filePrefixOption(prefix)
}

type headerOption string

func (opt headerOption) Apply(options *Options) {
	options.Header = string(opt)
}

func WithHeader(header string) OptionApplier {
	if !strings.HasPrefix(header, "//") {
		header = fmt.Sprintf("//%s", header)
	}
	return headerOption(header)
}
