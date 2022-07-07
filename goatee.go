package goatee

import (
	"fmt"

	"github.com/aymerick/raymond"
	"github.com/aymerick/raymond/lexer"
)

func IsLoop(template string) (string, bool) {
	l := lexer.Scan(template)
	a := l.NextToken()
	if a.Kind == lexer.TokenOpenBlock {
		n := l.NextToken()
		if n.Kind == lexer.TokenID && n.Val == "each" {
			v := l.NextToken()
			return v.Val, true
		}
	}
	return "", false
}

func IsMerge(template string) bool {
	l := lexer.Scan(template)
	a := l.NextToken()
	if a.Kind == lexer.TokenOpenBlock {
		n := l.NextToken()
		if n.Kind == lexer.TokenID && n.Val == "merge" {
			return true
		}
	}
	return false
}

func GetField(input any, field string) (any, bool) {
	switch tv := input.(type) {
	case map[string]any:
		return tv[field], true
	case map[string]string:
		return tv[field], true
	}
	return nil, false
}

func Render(template any, input any) (any, error) {
	switch tv := template.(type) {
	case map[string]any:
		out := map[string]any{}
		for keyT, valueT := range tv {
			keyV, err := raymond.Render(keyT, input)
			if err != nil {
				if field, ok := IsLoop(keyT); ok {
					if fieldData, ok := GetField(input, field); ok {
						if fArray, ok := fieldData.([]any); ok {
							subA := []any{}
							for _, x := range fArray {
								sub, err := Render(valueT, x)
								if err == nil {
									subA = append(subA, sub)
								}
							}
							return subA, nil
						}
						if fArray, ok := fieldData.([]map[string]any); ok {
							subA := []any{}
							for _, x := range fArray {
								sub, err := Render(valueT, x)
								if err == nil {
									subA = append(subA, sub)
								}
							}
							return subA, nil
						}
						if fArray, ok := fieldData.([]string); ok {
							subA := []any{}
							for _, x := range fArray {
								sub, err := Render(valueT, x)
								if err == nil {
									subA = append(subA, sub)
								}
							}
							return subA, nil
						}
						fmt.Printf("Sub type unknown: %#T\n", fieldData)
					} else {
						fmt.Printf("Field %s not found\n", field)
					}
				} else if ok := IsMerge(keyT); ok {
					if f, err := Render(valueT, input); err == nil {
						out := map[string]any{}
						if xa, ok := f.([]any); ok {
							for _, a := range xa {
								if xm, ok := a.(map[string]any); ok {
									for k, v := range xm {
										out[k] = v
									}
								}
							}
						}
						return out, nil
					}
				}
			} else {
				valueV, _ := Render(valueT, input)
				out[keyV] = valueV
			}
		}
		return out, nil
	case []any:
		out := []any{}
		for _, o := range tv {
			c, _ := Render(o, input)
			out = append(out, c)
		}
		return out, nil
	case string:
		return raymond.Render(tv, input)
	}
	return template, nil
}
