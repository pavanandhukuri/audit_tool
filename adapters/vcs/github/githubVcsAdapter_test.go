package github

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuditorService_Audit(t *testing.T) {
	t.Run("Should build schema from provided configuration", func(t *testing.T) {
		// Arrange

		// Track number of times a request is made
		requestCount := 0

		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			var err error
			if r.URL.Path == "/single" {
				requestCount++
				_, err = w.Write([]byte(`{"test_property": "test", "test_property2": {"test_property3": "test3"}}`))
			} else {
				requestCount++
				_, err = w.Write([]byte(`[{"test_array_property": "test1"}, {"test_array_property": "test2"}]`))
			}

			if err != nil {
				t.Fatalf("Failed to write response: %v", err)
			}
		}))
		defer testServer.Close()

		// Write the following content to a file named single.yml
		schema := map[string]interface{}{
			"schema": map[string]interface{}{
				"TestProperty": map[string]interface{}{
					"type":  "string",
					"value": "singeResultEndpoint.test_property",
				},
				"TestProperty2": map[string]interface{}{
					"type":  "string",
					"value": "singeResultEndpoint.test_property2.test_property3",
				},
				"TestArray": map[string]interface{}{
					"type":  "array",
					"value": "arrayEndpoint",
					"fields": map[string]interface{}{
						"TestArrayProperty": map[string]interface{}{
							"type":  "string",
							"value": "item.test_array_property",
						},
					},
				},
			},
			"endpoints": map[string]interface{}{
				"singeResultEndpoint": map[string]interface{}{
					"url":    testServer.URL + "/single",
					"method": "GET",
					"headers": map[string]interface{}{
						"Accept":        "application/vnd.github+json",
						"Authorization": "Bearer testToken",
					},
				},
				"arrayEndpoint": map[string]interface{}{
					"url":    testServer.URL + "/array",
					"method": "GET",
					"headers": map[string]interface{}{
						"Accept":        "application/vnd.github+json",
						"Authorization": "Bearer testToken",
					},
				},
			},
		}

		// Optionally, you can marshal this map back to YAML format to verify the structure
		yamlData, err := yaml.Marshal(&schema)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		content := string(yamlData)

		schemaFile := "schema.yml"
		err = os.WriteFile(schemaFile, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write file: %v", err)
		}
		defer os.Remove(schemaFile)

		versionControlSystem := NewGithubVcsAdapter(schemaFile)

		// Act
		vcsData, err := versionControlSystem.GetData()
		if err != nil {
			t.Fatalf("Failed to get schema: %v", err)
		}

		fmt.Println(vcsData)

		// Assert
		if value, exists := vcsData["TestProperty"]; exists {
			assert.Equal(t, "test", value)
		} else {
			t.Fatalf("Key 'test_property' not found in vcsData")
		}

		if value, exists := vcsData["TestProperty2"]; exists {
			assert.Equal(t, "test3", value)
		} else {
			t.Fatalf("Key 'test_property2' not found in vcsData")
		}

		if value, exists := vcsData["TestArray"]; exists {
			assert.Equal(t, []interface{}{
				map[string]interface{}{"TestArrayProperty": "test1"},
				map[string]interface{}{"TestArrayProperty": "test2"},
			}, value)
		} else {
			t.Fatalf("Key 'test_array_property' not found in vcsData array")
		}

		assert.Equal(t, 2, requestCount)
	})
}
