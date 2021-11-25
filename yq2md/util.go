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
		m.addString(y.options.MDHeader)
		m.addNewLines(2)
	}

	for _, category := range *(y.YamlStruct) {
		categLen := len(category.Items)

		m.addCategName(&category.Name)
		m.addNewLines(2)
		for i, item := range category.Items {
			if y.options.HTMLOptimization {
				m.addString("<div class=\"ques_instance\">")
				m.addNewLines(1)
			}

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

			if y.options.HTMLOptimization {
				m.addNewLines(1)
				m.addString("</div>")
			}
		}
	}
}

func (b *mdBuffer) addString(mdHeader string) {
	b.buf.WriteString(mdHeader)
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
	/* Sometimes the answer is empty */
	if *ans == "" || *ans == "\n" {
		return
	}

	if b.isHTMLOptimized {
		b.buf.WriteString(":::{.ques_ans}\n")
		if b.isAnswerToggle {
			b.buf.WriteString("<details><summary>Answer</summary>")
		}
	}
	b.buf.WriteString(*ans)
	if b.isHTMLOptimized {
		if b.isAnswerToggle {
			b.buf.WriteString("</details>")
		}
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
