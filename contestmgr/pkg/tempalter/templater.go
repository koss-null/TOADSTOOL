package templater

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	TemplateOpenBr  = "<<<"
	TemplateCloseBr = ">>>"
)

type (
	templater struct {
		templatePath string
		templateCode string
		resultPath   string

		templateList          map[string]struct{}
		templateToHandlerFunc map[string]func() string
		templateToHandler     map[string]string

		isTemplateFilled bool
		templateRegexp   *regexp.Regexp
	}

	Templater interface {
		ReadTemplate() error
		AddHandler(handlerName string, handler func() string)
		// Creates the file with all templates filled
		Build() error
	}
)

func NewTemplater(
	templatePath string,
	resultPath string,
) Templater {
	return &templater{
		templatePath,
		"",
		resultPath,
		make(map[string]struct{}),
		make(map[string]func() string),
		make(map[string]string),
		false,
		regexp.MustCompile(
			fmt.Sprintf("%s[[:alpha:]_]*%s", TemplateOpenBr, TemplateCloseBr),
		),
	}
}

func (t *templater) ReadTemplate() error {
	content, err := os.ReadFile(t.templatePath)
	if err != nil {
		return err
	}

	t.templateCode = string(content)

	// selecting all the templates from a file
	templates := t.templateRegexp.FindAllString(t.templateCode, -1)

	for i := range templates {
		tstrt, tend := len(TemplateOpenBr), len(templates[i])-len(TemplateCloseBr)
		t.templateList[templates[i][tstrt:tend]] = struct{}{}
	}

	t.isTemplateFilled = true
	return nil
}

func (t *templater) AddHandler(handlerName string, handler func() string) {
	t.templateToHandlerFunc[handlerName] = handler
}

func (t *templater) writeResult(result string) error {
	file, err := os.Create(fmt.Sprintf("%s%cmain.go", t.resultPath, os.PathSeparator))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, result)
	if err != nil {
		return err
	}
	return file.Sync()
}

func (t *templater) evalTemplateCode(template string) string {
	if val, ok := t.templateToHandler[template]; ok {
		return val
	}
	t.templateToHandler[template] = t.templateToHandlerFunc[template]()
	return t.templateToHandler[template]
}

func (t *templater) Build() error {
	// check if the template was read
	if !t.isTemplateFilled {
		return errors.New("template file was not read")
	}
	// check if all file templates have its handlers
	for k := range t.templateList {
		if _, ok := t.templateToHandlerFunc[k]; !ok {
			return fmt.Errorf(
				"there is no handler for the %s%s%s template",
				TemplateOpenBr, k, TemplateCloseBr,
			)
		}
	}

	resultCode := strings.Builder{}
	lastSplit := 0
	for _, i := range t.templateRegexp.FindAllStringIndex(t.templateCode, -1) {
		resultCode.WriteString(t.templateCode[lastSplit:i[0]])

		start, end := i[0]+len(TemplateOpenBr), i[1]-len(TemplateCloseBr)
		resultCode.WriteString(t.evalTemplateCode(t.templateCode[start:end]))

		lastSplit = i[1]
	}
	if lastSplit != len(t.templateCode)-1 {
		resultCode.WriteString(t.templateCode[lastSplit:len(t.templateCode)])
	}

	return t.writeResult(resultCode.String())
}
