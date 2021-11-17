package yq2md

import (
	"bytes"
)

type yamlDef struct {
	Name  string `yaml:"name"`
	Items []struct {
		Question string `yaml:"question"`
		Answer   string `yaml:"answer"`
	} `yaml:"items"`
}

type yamlDefs []yamlDef

type YqOptions struct {
	YamlFileName     string
	ToRandomize      bool
	ToNumerize       bool
	AllIntoOne       bool
	ShowAnswers      bool
	ToggleAnswer     bool
	HTMLOptimization bool
	MDHeader         string
}

type YqInstance struct {
	options    *YqOptions
	YamlStruct *yamlDefs
}

type mdBuffer struct {
	buf             bytes.Buffer
	isHTMLOptimized bool
	isAnswerToggle  bool
}

func NewYqInstance(options *YqOptions) *YqInstance {
	fileContents := returnFileContents(&options.YamlFileName)
	yamlStruct := returnYamlStruct(fileContents)

	if options.AllIntoOne {
		yamlStruct = yamlStruct.returnAllItemsDumpedStruct()
	}

	if options.ToRandomize {
		yamlStruct.randomizeItems()
	}

	return &YqInstance{
		options:    options,
		YamlStruct: yamlStruct,
	}
}

func (y *YqInstance) ReturnMD() *bytes.Buffer {
	mdBuf := &mdBuffer{
		isHTMLOptimized: y.options.HTMLOptimization,
		isAnswerToggle:  y.options.ToggleAnswer,
	}
	mdBuf.fillBufFromYqInstance(y)

	return &mdBuf.buf
}
