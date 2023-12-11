package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

const HOST_DEFAULT = "http://127.0.0.1"

func getPort() int {
	res, _ := strconv.Atoi(getDotEnvVariable("PORT"))
	return res
}

func getServerInfo(req *http.Request) string {
	clientIp := getClientIpAddr(req)
	// fmt.Printf("DEBUG clientIp %+v\n", clientIp)
	if clientIp == "" || clientIp == "127.0.0.1" {
		return fmt.Sprintf("%s:%d", HOST_DEFAULT, getPort())
	}
	return fmt.Sprintf("%s:%d", getDotEnvVariable("HOST_ROUTER"), getPort())
}

func getClientIpAddr(req *http.Request) string {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	return host
}

func getDotEnvVariable(key string) string {
	err := godotenv.Load("quiz/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Printf("server run in port %d \n", getPort())

	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/", getDashboard)
	mux.HandleFunc("/quiz/static/", staticHandler)
	mux.HandleFunc("/5_57_kotenov", get_5_57_kotenov)
	mux.HandleFunc("/check_5_57_kotenov", check_5_57_kotenov)

	err := http.ListenAndServe(fmt.Sprintf(":%d", getPort()), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	content := ""
	tmpl, err := template.ParseFiles(path.Join("quiz", "templates", "dashboard.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}
	// fmt.Printf("%+v\n", path)
	// fmt.Printf("%+v\n", path[1:])
	data, err := ioutil.ReadFile(path[1:])
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
