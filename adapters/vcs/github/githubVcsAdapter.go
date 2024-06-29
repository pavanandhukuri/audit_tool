package github

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"security_audit_tool/util"
	"strings"
)

type Configuration struct {
	Schema    map[string]SchemaField    `yaml:"schema"`
	Endpoints map[string]EndPointConfig `yaml:"endpoints"`
}

type SchemaField struct {
	Type   string                 `yaml:"type"`
	Value  string                 `yaml:"value"`
	Fields map[string]SchemaField `yaml:"fields"`
}

type EndPointConfig struct {
	Url     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
}

type EndpointData struct {
	Single      interface{}
	Array       []interface{}
	ContextData map[string]EndpointData
}

// SetSingle sets a single value
func (e *EndpointData) SetSingle(value interface{}) {
	e.Single = value
	e.Array = nil
}

// SetArray sets an array of values
func (e *EndpointData) SetArray(values []interface{}) {
	e.Single = nil
	e.Array = values
}

// IsSingle returns true if it's holding a single value
func (e *EndpointData) IsSingle() bool {
	return e.Single != nil && e.Array == nil
}

// IsArray returns true if it's holding an array of values
func (e *EndpointData) IsArray() bool {
	return e.Array != nil
}

func (e *EndpointData) Get() interface{} {
	if e.IsSingle() {
		return e.Single
	}
	if e.IsArray() {
		return e.Array
	}
	return nil
}

type VcsAdapter struct {
	EndpointData map[string]EndpointData
}

func (adapter *VcsAdapter) GetData() (map[string]interface{}, error) {
	config := getConfig()
	finalObj := make(map[string]interface{})

	for key, value := range config.Schema {
		adapter.resolveValue(value, config, finalObj, key, nil)
	}

	return finalObj, nil
}

func (adapter *VcsAdapter) resolveValue(value SchemaField, config Configuration, schemaObj map[string]interface{}, key string, context interface{}) {
	valueExpression := value.Value
	valuePrefix := getValuePrefix(valueExpression)

	if value.Type == "array" {
		adapter.handleArrayType(value, config, schemaObj, key, context, valuePrefix)
	} else {
		adapter.handleSingleType(value, config, schemaObj, key, context, valuePrefix)
	}
}

func (adapter *VcsAdapter) handleArrayType(value SchemaField, config Configuration, schemaObj map[string]interface{}, key string, context interface{}, valuePrefix string) {
	schemaObj[key] = make([]interface{}, 0)

	cachedEndpointData, exists := adapter.EndpointData[valuePrefix]
	var endpointData []interface{}
	if exists {
		endpointData = cachedEndpointData.Array
	} else {
		endpointData = getDataFromEndpoint[[]interface{}](config.Endpoints[valuePrefix], context)
		adapter.EndpointData[valuePrefix] = EndpointData{Array: endpointData}
	}

	for _, data := range endpointData {
		nestedObj := make(map[string]interface{})
		for nestedKey, nestedValue := range value.Fields {
			adapter.resolveValue(nestedValue, config, nestedObj, nestedKey, data)
		}
		schemaObj[key] = append(schemaObj[key].([]interface{}), nestedObj)
	}
}

func (adapter *VcsAdapter) handleSingleType(value SchemaField, config Configuration, schemaObj map[string]interface{}, key string, context interface{}, valuePrefix string) {
	var data interface{}

	if valuePrefix == "item" {
		data = context
	} else {
		data = adapter.getCachedData(config, valuePrefix, context)
	}

	valueData := getValueFromData(value.Value, data)
	schemaObj[key] = valueData
}

func (adapter *VcsAdapter) getCachedData(config Configuration, valuePrefix string, context interface{}) interface{} {
	contextKey, _ := serializeContext(context)
	cachedData, exists := adapter.EndpointData[valuePrefix]

	if exists {
		if context != nil {
			contextualCachedData, exists := cachedData.ContextData[contextKey]
			if exists {
				return contextualCachedData.Single
			}
			return adapter.fetchAndCacheData(config, valuePrefix, context, contextKey)
		}
		return cachedData.Single
	}
	return adapter.fetchAndCacheData(config, valuePrefix, context, contextKey)
}

func (adapter *VcsAdapter) fetchAndCacheData(config Configuration, valuePrefix string, context interface{}, contextKey string) interface{} {
	data := getDataFromEndpoint[interface{}](config.Endpoints[valuePrefix], context)
	if context != nil {
		adapter.EndpointData[valuePrefix] = EndpointData{ContextData: map[string]EndpointData{contextKey: {Single: data}}}
	} else {
		adapter.EndpointData[valuePrefix] = EndpointData{Single: data}
	}
	return data
}

func getValuePrefix(valueExpression string) string {
	if strings.Contains(valueExpression, ".") {
		return valueExpression[:strings.Index(valueExpression, ".")]
	}
	return valueExpression
}

func serializeContext(context interface{}) (string, error) {
	bytes, err := json.Marshal(context)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getValueFromData(valueExpression string, data interface{}) interface{} {
	valueSuffix := valueExpression[strings.Index(valueExpression, ".")+1:]

	if strings.Contains(valueSuffix, ".") {
		nestedValue := strings.Split(valueSuffix, ".")
		for _, key := range nestedValue {
			data = data.(map[string]interface{})[key]
		}
	} else {
		data = data.(map[string]interface{})[valueSuffix]
	}
	return data
}

func getDataFromEndpoint[T any](endpointConfig EndPointConfig, context interface{}) T {
	if endpointConfig.Method == "GET" {
		url := interpolate(endpointConfig.Url, context)
		headers := prepareHeaders(endpointConfig.Headers, context)
		data, err := util.GetAsType[T](url, headers)
		if err != nil {
			panic(err)
		}
		return data
	}
	return *new(T)
}

func prepareHeaders(headers map[string]string, context interface{}) map[string]string {
	interpolatedHeaders := make(map[string]string)
	for key, value := range headers {
		interpolatedHeaders[key] = interpolate(value, context)
	}
	return interpolatedHeaders
}

func getConfig() Configuration {
	currentWorkingDirectory, _ := os.Getwd()
	yamlFile, err := os.ReadFile(currentWorkingDirectory + "/resources/config.yml")
	if err != nil {
		panic(err)
	}

	var config Configuration
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func interpolate(input string, context interface{}) string {
	if input == "" {
		return input
	}

	result := interpolateContextVariables(input, context)
	return interpolateEnvVariables(result)
}

func interpolateContextVariables(input string, context interface{}) string {
	re := regexp.MustCompile(`\{\{item\.([A-Za-z_]+)\}\}`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		varName := re.FindStringSubmatch(match)[1]
		envValue := getValueFromData(varName, context)
		if envValue == "" {
			return match
		}
		return envValue.(string)
	})
}

func interpolateEnvVariables(input string) string {
	re := regexp.MustCompile(`\{\{env\.([A-Z_]+)\}\}`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		varName := re.FindStringSubmatch(match)[1]
		envValue := os.Getenv(varName)
		if envValue == "" {
			return match
		}
		return envValue
	})
}

func NewGithubVcsAdapter() *VcsAdapter {
	return &VcsAdapter{EndpointData: make(map[string]EndpointData)}
}
