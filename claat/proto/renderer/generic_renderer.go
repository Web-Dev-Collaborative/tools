package genrenderer

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/googlecodelabs/tools/third_party"
)

// TODO: Flip order of functions

// templateName Maps protos to their type string name
func templateName(el interface{}) string {
	switch el.(type) {
	case *tutorial.StylizedText, tutorial.StylizedText:
		return "StylizedText"
	case *tutorial.InlineCode, tutorial.InlineCode:
		return "InlineCode"
	case *tutorial.InlineContent, tutorial.InlineContent:
		return "InlineContent"
	case *tutorial.Paragraph, tutorial.Paragraph:
		return "Paragraph"
	case *tutorial.Link, tutorial.Link:
		return "Link"
	case *tutorial.Button, tutorial.Button:
		return "Button"
	case *tutorial.List, tutorial.List:
		return "List"
	case *tutorial.Image, tutorial.Image:
		return "Image"
	case *tutorial.ImageBlock, tutorial.ImageBlock:
		return "ImageBlock"
	case *tutorial.YoutubeVideo, tutorial.YoutubeVideo:
		return "YoutubeVideo"
	case *tutorial.CodeBlock, tutorial.CodeBlock:
		return "CodeBlock"
	}
	// This will cause a debug-friendly panic
	return TypeNotSupported("genrenderer.templateName", el)
}

// outputFormatTemplateName concatenates the template name mapping of the passed proto
// with its output package extension
func outputFormatTemplateName(el interface{}, t *template.Template) string {
	return templateName(el) + "." + t.Name()
}

// ExecuteTemplate returns the evaluated template per passed templating
// namespace, based on the passed tutorial proto type string name
func ExecuteTemplate(el interface{}, t *template.Template) string {
	var w bytes.Buffer
	e := t.ExecuteTemplate(&w, outputFormatTemplateName(el, t), el)
	if e != nil {
		// This method outputs directly to templates. Panicking to surfance errors
		// since we should not handle multiple returns in templates.
		// Errors will be more gracefully handled in output-format packages
		panic(fmt.Sprintf("Templating panic: %s\n", e))
	}
	return w.String()
}