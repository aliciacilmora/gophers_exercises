package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc which also implements http.Handler
// that will attempt to map any paths (keys in the map) to their corresponding
// URL aflues that each key in the map points to, in string format

// If the path is not provided in the map, then the fallback http.Handler will
// be called instead.

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// if we match a path redirect to it
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// else
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the YAML
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	// 2. Convert  YAML array into map
	pathsToUrls := buildMap(pathUrls)
	// 3. Return a map hamdler using the mapping
	return MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
