package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/KaiserWerk/sessionstore"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	version = "0.0.0" // inject at compile time
	commitHash = "xxx"
	versionDate = time.Now() // inject at compile time
	listenPort string
	configFile string
	centralConfig configuration
	templates map[string]*template.Template
	sessMgr *sessionstore.SessionManager

)

func main() {
	writeToConsole("Tiny Build Server, Version " + version + " build from commit " + commitHash + " at " + versionDate.String())
	flag.StringVar(&listenPort, "p", "8271", "The port which the build server should listen on")
	flag.StringVar(&configFile, "c", "app.yaml", "The location of the configuration file")
	flag.Parse()

	centralConfig = getConfiguration()
	sessMgr = sessionstore.NewManager("tbs_sessid")
	funcMap := template.FuncMap{
		"getBuildDefCaption": getBuildDefCaption,
		"getUsernameById": getUsernameById,
		"getFlashbag": getFlashbag(sessMgr),
	}
	templates = populateTemplates(funcMap)

	listenAddr := fmt.Sprintf(":%s", listenPort)
	writeToConsole("server will be handling requests at port " + listenPort)

	//if _, err := loadSysConfig(); err != nil {
	//	log.Fatal("could not handle config/app.yaml file; something went wrong")
	//}

	router := mux.NewRouter()
	router.Use(limit)
	router.NotFoundHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		t := templates["404.html"]
		if t != nil {
			_ = t.Execute(w, r.URL.Path)
		}
	})

	// asset file handlers
	router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("public"))).Methods("GET")
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("public"))).Methods("GET")
	router.PathPrefix("/assets/").Handler(http.FileServer(http.Dir("public"))).Methods("GET")

	// site handlers
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	router.HandleFunc("/logout", logoutHandler).Methods("GET", "POST")
	router.HandleFunc("/password/request", requestNewPasswordHandler).Methods("GET", "POST")
	router.HandleFunc("/password/reset", resetPasswordHandler).Methods("GET", "POST")
	router.HandleFunc("/register", registrationHandler).Methods("GET", "POST")

	router.HandleFunc("/admin/settings", adminSettingsHandler).Methods("GET", "POST")
	router.HandleFunc("/admin/buildtarget/list", adminBuildTargetListHandler).Methods("GET")
	router.HandleFunc("/admin/buildtarget/add", adminBuildTargetAddHandler).Methods("GET", "POST")
	router.HandleFunc("/admin/buildtarget/{id}/edit", adminBuildTargetEditHandler).Methods("GET", "POST")
	router.HandleFunc("/admin/buildtarget/{id}/remove", adminBuildTargetRemoveHandler).Methods("GET")
	router.HandleFunc("/admin/buildstep/list", adminBuildStepListHandler).Methods("GET")
	router.HandleFunc("/admin/buildstep/add", adminBuildStepAddHandler).Methods("GET", "POST")
	router.HandleFunc("/admin/buildstep/{id}/edit", adminBuildStepEditHandler).Methods("GET", "POST")
	router.HandleFunc("/admin/buildstep/{id}/remove", adminBuildStepRemoveHandler).Methods("GET")
	// show?

	router.HandleFunc("/builddefinition/list", buildDefinitionListHandler).Methods("GET")
	router.HandleFunc("/builddefinition/add", buildDefinitionAddHandler).Methods("GET", "POST")
	router.HandleFunc("/builddefinition/{id}/show", buildDefinitionShowHandler).Methods("GET")
	router.HandleFunc("/builddefinition/{id}/edit", buildDefinitionEditHandler).Methods("GET", "POST")
	router.HandleFunc("/builddefinition/{id}/remove", buildDefinitionRemoveHandler).Methods("GET")
	router.HandleFunc("/builddefinition/{id}/listexecutions", buildDefinitionListExecutionsHandler).Methods("GET")

	router.HandleFunc("/buildexecution/list", buildExecutionListHandler).Methods("GET")
	router.HandleFunc("/buildexecution/{id}/show", buildExecutionShowHandler).Methods("GET")

	// API handlers
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/bitbucket-receive", bitBucketReceiveHandler).Methods("POST")
	apiRouter.HandleFunc("/github-receive", gitHubReceiveHandler).Methods("POST")
	apiRouter.HandleFunc("/gitlab-receive", gitLabReceiveHandler).Methods("POST")
	apiRouter.HandleFunc("/gitea-receive", giteaReceiveHandler).Methods("POST")

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go readConsoleInput(quit)

	go func() {
		<-quit
		writeToConsole("server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("could not gracefully shut down the server: %v\n", err)
		}
		close(done)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s: %v\n", listenAddr, err)
	}
	<-done
	writeToConsole("server shutdown complete. Have a nice day!")
}

func populateTemplates(fm template.FuncMap) map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html")).Funcs(fm)
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("failed to open template block directory: " + err.Error())
	}
	defer dir.Close()
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("failed to contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("failed to open template '"+fi.Name()+"': " + err.Error())
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("failed to read content from file '"+fi.Name()+"': " + err.Error())
		}
		_ = f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("failed to parse contents of file '"+fi.Name()+"': " + err.Error())
		}
		result[fi.Name()] = tmpl
	}
	return result
}