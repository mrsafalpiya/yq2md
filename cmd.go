package main

import (
	"fmt"
	"log"

	"github.com/mrsafalpiya/yq2md/yq2md"
	"github.com/spf13/pflag"
)

var passedArgs yq2md.YqOptions

func init() {
	parseArgs()
}

func main() {
	yq2mdInstance := yq2md.NewYqInstance(&passedArgs)

	mdBuf := yq2mdInstance.ReturnMD()
	fmt.Printf("%s", mdBuf)
}

func parseArgs() {
	var err error

	pflag.BoolVarP(
		&passedArgs.ToRandomize,
		"randomize",
		"r",
		false,
		"Enable this flag to randomize items within categories",
	)
	pflag.BoolVarP(
		&passedArgs.ToNumerize,
		"numerize",
		"n",
		false,
		"Enable this flag to numerize items within categories",
	)
	pflag.BoolVarP(
		&passedArgs.AllIntoOne,
		"all",
		"a",
		false,
		"Enable this flag to dump all questions into a single root category (disable categorization)",
	)
	pflag.BoolVarP(
		&passedArgs.ShowAnswers,
		"show-answer",
		"s",
		false,
		"Enable this flag to answers of the questions too",
	)
	pflag.BoolVarP(
		&passedArgs.ToggleAnswer,
		"toggle-answer",
		"t",
		false,
		"Enable this flag to make the answer toggle (Available only when --html flag is enabled)",
	)
	pflag.BoolVar(
		&passedArgs.HTMLOptimization,
		"html",
		false,
		"Produce markdown optimized for HTML conversion",
	)

	pflag.Parse()

	passedArgs.YamlFileName, err = getFileNameFromArgs()
	if err != nil {
		log.Fatal("[Error] Cannot read file: ", err)
	}

	passedArgs.MDHeader = MDHeader
}
