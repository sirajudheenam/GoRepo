package urlshort

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
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
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	
	// 1. Parse the yaml somehow
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	// 2. Convert the yaml array into a map
	pathsToUrls := buildMap(pathUrls)

	// return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil

}

func buildMap(pathUrls []pathUrl) map[string]string {
		// 2. Convert the yaml array into a map
		pathsToUrls := make(map[string]string)
		for _, pu := range pathUrls {
			pathsToUrls[pu.Path] = pu.URL
		}
		return pathsToUrls
}
func parseYaml(data []byte,) ([]pathUrl, error) {

	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}
type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func parseJSON(jsonData []byte) (pathsToURLs []pathUrl, err error) {
	err = json.Unmarshal(jsonData, &pathsToURLs)
	return
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     [
//	     {
//         "path": "/some-path",
//         "url": "https://www.some-url.com/demo"
//       }
//	   ]
//
// The only errors that can be returned all related to having
// invalid JSON data.

func JSONHandler(jsonData []byte, fallback http.Handler) (jsonHandler http.HandlerFunc, err error) {
	parsedJSON, err := parseJSON(jsonData)
	// fmt.Println("Parsed JSON", parsedJSON)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJSON)
	jsonHandler = MapHandler(pathMap, fallback)
	return jsonHandler, nil
}

// DBHandler will use the provided Bolt database and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the DB, then the
// fallback http.Handler will be called instead.
func DBHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var url string
		err := db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("paths"))
			bts := b.Get([]byte(req.URL.Path))
			if bts != nil {
				url = string(bts)
			}
			return nil
		})

		if err == nil && url != "" {
			http.Redirect(res, req, url, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(res, req)
		}
	})
}