package yq2md

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
)

func returnFileContents(filename *string) *[]byte {
	tempData, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal("Cannot read given file:", err)
	}

	return &tempData
}

func returnYamlStruct(yamlData *[]byte) *yamlDefs {
	structInstance := yamlDefs{}

	err := yaml.Unmarshal(*yamlData, &structInstance)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &structInstance
}

func (y *yamlDefs) returnAllItemsDumpedStruct() *yamlDefs {
	var tempYamlStruct yamlDef

	for _, categ := range *y {
		for _, item := range categ.Items {
			tempYamlStruct.Items = append(tempYamlStruct.Items, item)
		}
	}

	return &yamlDefs{tempYamlStruct}
}

func (y *yamlDefs) randomizeItems() {
	for i := 0; i < len(*y); i++ {
		items := (*y)[i].Items

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(
			len(items),
			func(i, j int) { items[i], items[j] = items[j], items[i] },
		)
	}
}

func (m *mdBuffer) fillBufFromYqInstance(y *YqInstance) {
	if !y.options.HTMLOptimization {
		m.addString(&y.options.MDTemplate)
		m.addNewLines(2)
	}

	for _, category := range *(y.YamlStruct) {
		categLen := len(category.Items)

		m.addCategName(&category.Name)
		m.addNewLines(2)
		for i, item := range category.Items {
			if y.options.ToNumerize {
				m.addQNum(&i, &categLen)
				m.addNewLines(2)
			}
			m.addQuesTitle(&item.Question)
			m.addNewLines(1)
			if y.options.ShowAnswers {
				m.addQuesAns(&item.Answer)
				m.addNewLines(1)
			}
			m.addHLine()
			m.addNewLines(2)
		}
	}
}

func (b *mdBuffer) addString(mdTemplate *string) {
	b.buf.WriteString(*mdTemplate)
}

func (b *mdBuffer) addNewLines(n int) {
	for i := 0; i < n; i++ {
		b.buf.WriteString("\n")
	}
}

func (b *mdBuffer) addCategName(categName *string) {
	b.buf.WriteString("# ")
	b.buf.WriteString(*categName)
	if b.isHTMLOptimized {
		b.buf.WriteString("{.categ_name}")
	}
}

func (b *mdBuffer) addQNum(currentIndex *int, categLen *int) {
	if b.isHTMLOptimized {
		b.buf.WriteString(":::{.ques_num}\n")
	}
	b.buf.WriteString("*Question #")
	b.buf.WriteString(strconv.Itoa(*currentIndex + 1))
	b.buf.WriteString(" out of ")
	b.buf.WriteString(strconv.Itoa(*categLen))
	b.buf.WriteString("*")
	if b.isHTMLOptimized {
		b.buf.WriteString("\n")
		b.buf.WriteString(":::")
	}
}

func (b *mdBuffer) addQuesTitle(title *string) {
	if b.isHTMLOptimized {
		b.buf.WriteString(":::{.ques_title}\n")
	}
	b.buf.WriteString(*title)
	if b.isHTMLOptimized {
		b.buf.WriteString("\n")
		b.buf.WriteString(":::")
	}
}

func (b *mdBuffer) addQuesAns(ans *string) {
	if b.isHTMLOptimized {
		b.buf.WriteString(":::{.ques_ans}\n")
	}
	b.buf.WriteString(*ans)
	if b.isHTMLOptimized {
		b.buf.WriteString("\n")
		b.buf.WriteString(":::")
	}
}

func (b *mdBuffer) addHLine() {
	if b.isHTMLOptimized {
		b.buf.WriteString("<hr>")
		return
	}
	b.buf.WriteString("$\\hrulefill$")
}