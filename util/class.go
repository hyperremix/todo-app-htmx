package util

import (
	"strings"

	"github.com/a-h/templ"
)

func GetClassAttributeValue(attributes templ.Attributes) string {
	if class, ok := attributes["class"].(string); ok {
		return class
	}

	return ""
}

func GetClassAttributeSlice(attributes templ.Attributes) []string {
	if class, ok := attributes["class"].(string); ok {
		return strings.Split(class, " ")
	}

	return []string{}
}

func DeleteClassAttribute(attributes templ.Attributes) templ.Attributes {
	delete(attributes, "class")
	return attributes
}
