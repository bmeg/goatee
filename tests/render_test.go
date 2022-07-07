package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bmeg/goatee"
	"github.com/google/go-cmp/cmp"
)

func TestRender(t *testing.T) {

	paths, _ := filepath.Glob("*_template.json")

	for _, p := range paths {
		fmt.Printf("Checking %s\n", p)
		template, _ := ioutil.ReadFile(p)
		input, _ := ioutil.ReadFile(strings.TrimSuffix(p, "_template.json") + "_input.json")
		output, _ := ioutil.ReadFile(strings.TrimSuffix(p, "_template.json") + "_output.json")

		templateData := map[string]any{}
		json.Unmarshal(template, &templateData)

		inputData := map[string]any{}
		json.Unmarshal(input, &inputData)

		outputData := map[string]any{}
		json.Unmarshal(output, &outputData)

		renderData, _ := goatee.Render(templateData, inputData)

		if !cmp.Equal(outputData, renderData) {
			//outStr, _ := json.MarshalIndent(outputData, "", "   ")
			//renderStr, _ := json.MarshalIndent(renderData, "", "   ")
			//t.Errorf("%s != %s", outStr, renderStr)
			t.Errorf("%s", cmp.Diff(outputData, renderData))
		}
	}
}
