package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gossip-phantom-hosting/auth"
	"github.com/gossip-phantom-hosting/config"
	"github.com/gossip-phantom-hosting/utils"
)

var mnString string
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("website/*"))
}

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/deploy", DeployMasternodeHandler).Methods("GET")
	r.HandleFunc("/configuration", ViewConfigurationHandler).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generateconfigfile", GenerateConfigFile).Methods("POST")
	api.HandleFunc("/generatemasternodestring", GenerateMasternodeString).Methods("POST")
	api.HandleFunc("/addmasternode", AddMasternode).Methods("POST")

	api.HandleFunc("/register", auth.Register).Methods("POST")
	api.HandleFunc("/login", auth.Login).Methods("POST")

	api.HandleFunc("/viewconfiguration", ViewConfigFile).Methods("GET")

	http.Handle("/", r)

	r.PathPrefix("/www/").Handler(http.StripPrefix("/www/", http.FileServer(http.Dir("www"))))

	fmt.Println("Running on http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", r))
}

// Frontend Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		fmt.Println("error template")
	}
}

func DeployMasternodeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "deploy.gohtml", nil)
	if err != nil {
		fmt.Println("error template")
	}
}

func ViewConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "configuration.gohtml", nil)
	if err != nil {
		fmt.Println("error template")
	}
}

// API Handlers
func GenerateConfigFile(w http.ResponseWriter, r *http.Request) {
	err := config.GenerateConfigurationFile("conf/masternode.txt")

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, "Configuration file created", nil)
}

func GenerateMasternodeString(w http.ResponseWriter, r *http.Request) {
	var mnConfig config.MasternodeString

	err := json.NewDecoder(r.Body).Decode(&mnConfig)

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	mnString, err = config.GenerateNodeDetails(mnConfig)
	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, mnString, nil)
}

func AddMasternode(w http.ResponseWriter, r *http.Request) {
	err := config.AddMasternodeToConfigFile("conf/masternode.txt", mnString)

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, mnString, nil)
}

func ViewConfigFile(w http.ResponseWriter, r *http.Request) {
	data, err := config.ViewConfiguration("conf/masternode.txt")

	if err != nil {
		utils.Respond(w, nil, err)
		return
	}

	utils.Respond(w, data, nil)
}
