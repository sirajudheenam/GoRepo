package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/sirajudheenam/goRepo/gophercises/urlshort"
)

var (
	pathsFile = flag.String("pathsFile", "paths.yml", "The file containing shortened paths to URLs")
	initDB    = flag.Bool("initDB", false, "Whether or not to initialize the paths database")
)

func getFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file %s", fileName)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file %s", fileName)
	}

	return buf.Bytes()
}

func initBoltDB(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("paths"))
		if err != nil {
			return err
		}

		err = b.Put([]byte("/yahoo"), []byte("https://yahoo.com"))
		err = b.Put([]byte("/github"), []byte("https://github.com"))

		return err
	})
}

func main() {
	mux := defaultMux()

	flag.Parse()

	ext := filepath.Ext(*pathsFile)

	fmt.Println("EXT is", ext)
	fmt.Println("Flag is", flag.Args())

	var handler, jsonHandler, yamlHandler, mapHandler http.Handler
	var err error

	if ext == ".yml" {
		handler, err = urlshort.YAMLHandler(getFileBytes(*pathsFile), mux)
		if err != nil {
			panic(err)
		}
	} else if ext == ".json" {
		handler, err = urlshort.JSONHandler(getFileBytes(*pathsFile), mux)
		if err != nil {
			panic(err)
		}
	} else if ext == ".db" {
		db, err := bolt.Open(*pathsFile, 0600, nil)
		if err != nil {
			panic(err)
		}

		defer db.Close()

		if *initDB {
			if err = initBoltDB(db); err != nil {
				panic(err)
			}
		}

		handler = urlshort.DBHandler(db, mux)
	} else { // if no file path was given
		// Build the MapHandler using the mux as the fallback
		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}
		mapHandler = urlshort.MapHandler(pathsToUrls, mux)

		// Build the YAMLHandler using the mapHandler as the
		// fallback
		yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
		// yamlHandler will fallback to mapHandler
		yamlHandler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}

		// sample json would be
		json := `[ { "path": "/ttt", "url": "https://technotipstoday.dev"},
					{ "path": "/tttblog", "url": "https://blog.technotipstoday.dev"},
					{ "path": "/medium", "url": "https://medium.com/technotipstoday"},
					{ "path": "/yt", "url": "https://youtube.com" } ]`

		// jsonHandler will fallback to yamlHandler
		jsonHandler, err = urlshort.JSONHandler([]byte(json), yamlHandler)
		handler = jsonHandler
		if err != nil {
			panic(err)
		}
		// log.Fatal("Paths file needs to be either a YAML, a JSON or a bolt DB file")
	}

	fmt.Println("Starting the server on :8080")
	fmt.Println("Access the Desired URLs by visiting http://localhost:8080/{ttt|tttblog|medium|yt}")
	http.ListenAndServe(":8080", handler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	helloWorldHtml := `<html>
	<head>
		<title>Fallback URL</title>
	</head>
	<body>
		<h1>Hello World </h1>
	</body>
	</html>`
	fmt.Fprintln(w, helloWorldHtml)
}
