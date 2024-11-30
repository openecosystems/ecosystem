package main

import (
	"strings"

	surface "github.com/googleapis/gnostic/surface"
)

const parameters = "parameters"

// ParameterList returns a string representation of a method's parameters
func ParameterList(parametersType *surface.Type) string {
	result := "ctx context.Context,\n"
	if parametersType != nil {
		result += parameters + " " + parametersType.Name
	}
	return result
}

func (renderer *Renderer) RenderClient() ([]byte, error) {
	f := NewLineWriter()

	f.WriteLine("// GENERATED FILE: DO NOT EDIT!")
	f.WriteLine(``)
	f.WriteLine("package " + renderer.Package)

	// imports will be automatically added by imports pkg

	f.WriteLine(`// Client represents an API client.`)
	f.WriteLine(`type Client struct {`)
	f.WriteLine(`  service string`)
	f.WriteLine(`  APIKey string`)
	f.WriteLine(`  client *http.Client`)
	f.WriteLine(`}`)

	f.WriteLine(`// NewClient creates an API client.`)
	f.WriteLine(`func NewClient(service string, c *http.Client) *Client {`)
	f.WriteLine(`	client := &Client{}`)
	f.WriteLine(`	client.service = service`)
	f.WriteLine(`  if c != nil {`)
	f.WriteLine(`    client.client = c`)
	f.WriteLine(`  } else {`)
	f.WriteLine(`    client.client = http.DefaultClient`)
	f.WriteLine(`  }`)
	f.WriteLine(`	return client`)
	f.WriteLine(`}`)

	for _, method := range renderer.Model.Methods {
		parametersType := renderer.Model.TypeWithTypeName(method.ParametersTypeName)
		responsesType := renderer.Model.TypeWithTypeName(method.ResponsesTypeName)

		f.WriteLine(commentForText(method.Description))
		f.WriteLine(`func (client *Client) ` + method.ClientName + `(`)
		f.WriteLine(ParameterList(parametersType) + `) (`)
		if method.ResponsesTypeName == "" {
			f.WriteLine(`err error,`)
		} else {
			f.WriteLine(`response *` + method.ResponsesTypeName + `,`)
			f.WriteLine(`err error,`)
		}
		f.WriteLine(` ) {`)

		path := method.Path
		path = strings.Replace(path, "{+", "{", -1)
		f.WriteLine(`path := client.service + "` + path + `"`)

		if parametersType != nil {
			if parametersType.HasFieldWithPosition(surface.Position_PATH) {
				for _, field := range parametersType.Fields {
					if field.Position == surface.Position_PATH {
						f.WriteLine(`path = strings.Replace(path, "{` + field.Name + `}", fmt.Sprintf("%v", ` +
							parameters + "." + strings.Title(field.ParameterName) + `), 1)`)
					}
				}
			}
			if parametersType.HasFieldWithPosition(surface.Position_QUERY) {
				f.WriteLine(`v := url.Values{}`)
				for _, field := range parametersType.Fields {
					if field.Position == surface.Position_QUERY {
						if field.NativeType == "string" {
							f.WriteLine(`if (` + parameters + "." + strings.Title(field.ParameterName) + ` != "") {`)
							f.WriteLine(`  v.Set("` + field.Name + `", ` + parameters + "." + strings.Title(field.ParameterName) + `)`)
							f.WriteLine(`}`)
						}
					}
				}
				f.WriteLine(`if client.APIKey != "" {`)
				f.WriteLine(`  v.Set("key", client.APIKey)`)
				f.WriteLine(`}`)
				f.WriteLine(`if len(v) > 0 {`)
				f.WriteLine(`  path = path + "?" + v.Encode()`)
				f.WriteLine(`}`)
			}
		}

		if method.Method == "POST" {
			f.WriteLine(`payload := new(bytes.Buffer)`)
			if parametersType != nil && parametersType.FieldWithPosition(surface.Position_BODY) != nil {
				f.WriteLine(`json.NewEncoder(payload).Encode(` + parameters + "." + strings.Title(parametersType.FieldWithPosition(surface.Position_BODY).FieldName) + `)`)
			}
			f.WriteLine(`req, err := http.NewRequest("` + method.Method + `", path, payload)`)
			f.WriteLine(`reqHeaders := make(http.Header)`)
			f.WriteLine(`reqHeaders.Set("Content-Type", "application/json")`)
			f.WriteLine(`req.Header = reqHeaders`)
		} else {
			f.WriteLine(`req, err := http.NewRequest("` + method.Method + `", path, nil)`)
		}
		f.WriteLine(`if err != nil {return}`)
		f.WriteLine(`req = req.WithContext(ctx)`)
		f.WriteLine(`resp, err := client.client.Do(req)`)
		f.WriteLine(`if err != nil {return}`)
		f.WriteLine(`defer resp.Body.Close()`)
		f.WriteLine(`if resp.StatusCode != 200 {`)

		if responsesType != nil {
			f.WriteLine(`	return nil, errors.New(resp.Status)`)
		} else {
			f.WriteLine(`	return errors.New(resp.Status)`)
		}
		f.WriteLine(`}`)

		if responsesType != nil {
			f.WriteLine(`response = &` + responsesType.TypeName + `{}`)
			f.WriteLine(`body, err := ioutil.ReadAll(resp.Body)`)
			f.WriteLine(`if err != nil {return nil, err}`)
			f.WriteLine(`err = json.Unmarshal(body, response)`)
			f.WriteLine(`if err != nil {return nil, err}`)
			f.WriteLine("return response, nil")
		} else {
			f.WriteLine("return nil")
		}
		f.WriteLine("}")
	}

	return f.Bytes(), nil
}
