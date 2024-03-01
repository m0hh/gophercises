package urlshort

import (
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusPermanentRedirect)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		pathsToUrls, err := parseYaml(yml)
		if err != nil {
			fmt.Println(err)
			return
		}

		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusPermanentRedirect)
			return
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

type urlss struct {
	Short string `yaml:"path"`
	Dest  string `yaml:"url"`
}

func parseYaml(yamlString []byte) (map[string]string, error) {

	output := make(map[string]string)

	var urls []urlss

	err := yaml.Unmarshal(yamlString, &urls)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for _, url := range urls {
		output[url.Short] = url.Dest
	}

	return output, nil
}
