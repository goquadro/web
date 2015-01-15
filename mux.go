package main

import (
	"log"
	"net/http"
	"os"

	"git.qrawl.net/qdoc/core"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/gorelic"
	"github.com/martini-contrib/render"
)

const (
	API_ENDPOINT = "/v1" // https://api.goquadro.com/v1
)

var (
	// config is the struct in which are stored all the credentials
	// used throughout the package.
	gqConfig config
)

type config struct {
	serveAddress       string
	resources          string
	newRelicAppName    string
	newRelicKey        string
	jwtSignKey         []byte // openssl genrsa -out goquadro_jwt.rsa 2048
	jwtVerifyKey       []byte // openssl rsa -in goquadro_jwt.rsa -pubout > goquadro_jwt.rsa.pub
	googleOauth2Id     string
	googleOauth2Secret string
}

func main() {
	m := martini.Classic()
	if gqConfig.newRelicKey != "" {
		gorelic.InitNewrelicAgent(gqConfig.newRelicKey, gqConfig.newRelicAppName, true)
		m.Use(gorelic.Handler)
	}
	m.Use(render.Renderer(render.Options{
		//Directory:  RESOURCES + "/templates",   // Specify what path to load the templates from.
		//Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Layout:     "base",                     // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		IndentJSON: true, // Output human readable JSON
		//IndentXML:  true,                       // Output human readable XML
	}))

	//m.Use(JwtCheck)

	m.Post(API_ENDPOINT+"/login", binding.Bind(core.User{}), ApiUserLogin)
	m.Post(API_ENDPOINT+"/signup", binding.Bind(core.User{}), ApiUserSignup)

	m.Get(API_ENDPOINT+"/me", ApiUserGetMe)

	m.Get(API_ENDPOINT+"/me/documents", ApiDocumentsGetAll)
	m.Post(API_ENDPOINT+"/me/documents", binding.Bind(core.Document{}), ApiDocumentsPost)
	m.Get(API_ENDPOINT+"/me/documents/:id", ApiDocumentsGetOne)
	m.Put(API_ENDPOINT+"/me/documents/:id", binding.Bind(core.Document{}), ApiDocumentsPut)
	m.Delete(API_ENDPOINT+"/me/documents/:id", ApiDocumentsDelete)

	//m.Get(API_ENDPOINT+"/me/topics", ApiTopicsGetAll)

	log.Println("Serving on", gqConfig.serveAddress)
	log.Fatal(http.ListenAndServe(gqConfig.serveAddress, m))
}

// Gets the variable from the environment. `def` is the default value
// that gets used if no env is found with that name.
func getenv(varName, def string) string {
	if newVar := os.Getenv(varName); newVar != "" {
		return newVar
	}
	return def
}

func init() {
	gqConfig = config{
		serveAddress:       getenv("QDOC_WEB_SERVE_ADDRESS", "localhost:8001"),
		resources:          getenv("QDOC_RESOURCES_DIR", "/var/www/goquadro"),
		newRelicAppName:    "goquadro",
		newRelicKey:        getenv("QDOC_NEWRELIC_KEY", ""),
		googleOauth2Id:     getenv("QDOC_GOAUTH_ID", ""),
		googleOauth2Secret: getenv("QDOC_GOAUTH_SECRET", ""),
		jwtSignKey:         getenv("QDOC_JWT_PRIVATE_KEY", ""),
		jwtVerifyKey:       getenv("QDOC_JWT_PUBLIC_KEY", ""),
	}
}
