package utils

import (
	"fmt"
	"strings"

	openAPI "github.com/go-openapi/spec"
)

func Join(enums []interface{}) string {
	if len(enums) == 0 {
		return ""
	}
	strEnums := make([]string, len(enums))
	for i, e := range enums {
		strEnums[i] = fmt.Sprintf("%v", e)
	}
	return strings.Join(strEnums, ", ")
}

func HasEnums(params []openAPI.Parameter) bool {
	for _, param := range params {
		if len(param.Enum) > 0 {
			return true
		}
	}
	return false
}
