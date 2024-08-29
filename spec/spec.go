package spec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"

	"gopkg.in/yaml.v2"
)

// API contains all data related to an API
type API struct {
	ID, URL string
	Data    []byte
	Spec    *spec.Swagger
}

// APIs contains all data for all APIs
type APIs []*API

// Load fetches OpenAPI JSON/YAML and parses the spec
func (a *APIs) Load() error {
	fmt.Println("Loading specs...")
	for _, api := range *a {
		if err := api.getJSON(); err != nil {
			return err
		}
		if err := api.loadSpec(); err != nil {
			return err
		}
	}
	fmt.Println("Specs loaded")
	return nil
}

func (a *API) getJSON() error {
	// TODO check the URL is JSON or YAML first and error if not
	var data []byte
	var err error

	if strings.HasPrefix(a.URL, "http://") || strings.HasPrefix(a.URL, "https://") {
		// Remote URL

		// TODO should be using timeouts etc, consider go-ns library instead
		res, err := http.Get(a.URL)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		data, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
	} else {
		// Local file path
		data, err = ioutil.ReadFile(a.URL)
		if err != nil {
			return err
		}
	}

	a.Data = data

	if strings.HasSuffix(strings.ToLower(a.URL), ".json") {
		return nil
	}

	var yamlData interface{}
	if err = yaml.Unmarshal(a.Data, &yamlData); err != nil {
		return err
	}
	if err = transformData(&yamlData); err != nil {
		return err
	}

	a.Data, err = json.Marshal(yamlData)
	if err != nil {
		return err
	}

	return nil
}

func (a *API) loadSpec() error {
	err := json.Unmarshal(a.Data, &a.Spec)
	if err != nil {
		return err
	}

	spec.ExpandSpec(a.Spec, &spec.ExpandOptions{})

	return nil
}

// transformData replaces map[interface{}]interface{} with map[string]interface{}
// so that it's accepted by json.Marshal()
// credit: https://github.com/bronze1man/yaml2json/blob/ee8196e587313e98831c040c26262693d48c1a0c/main.go#L48
func transformData(pIn *interface{}) (err error) {
	switch in := (*pIn).(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(in))
		for k, v := range in {
			if err = transformData(&v); err != nil {
				return err
			}
			var sk string
			switch k.(type) {
			case string:
				sk = k.(string)
			case int:
				sk = strconv.Itoa(k.(int))
			default:
				return fmt.Errorf("type mismatch: expect map key string or int; got: %T", k)
			}
			m[sk] = v
		}
		*pIn = m
	case []interface{}:
		for i := len(in) - 1; i >= 0; i-- {
			if err = transformData(&in[i]); err != nil {
				return err
			}
		}
	}

	return nil
}
