package main

import (
	"strings"
	"testing"
	"log"
)

type TestStruct struct {
	TemplateImpl
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T){
		s := &TestStruct{}
		log.Println(s)

		res := s.ExecuteAlgorithm(s)
		expected := "world"

		if !strings.Contains(res, expected) {
		  t.Errorf("Expected string '%s' wasn't found on returned string\n", expected)
		}
	})

	t.Run("Using anonymous functions", func(t *testing.T){
		 m := new(AnonymousTemplate)
		 res := m.ExecuteAlgorithm(func() string {
		   return "world"
		 })
		 expectedOrError(res, " world ", t)
	})

	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T){
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("Can not continue with a nil MessageRetriever")
		}

		template := TestStruct{}
		res := template.ExecuteAlgorithm(messageRetriever)

		expectedOrError(res, " world ", t)
	})
}

func expectedOrError(res string, expected string, t *testing.T){
	if !strings.Contains(res, expected) {
		t.Errorf("Expected string '%s' was not found on returned string\n", expected)
	}
}