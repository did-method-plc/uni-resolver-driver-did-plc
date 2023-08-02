package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// single handler for all routes
func allRoutes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		io.WriteString(w, "did:plc universal resolver driver\n")
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/1.0/identifiers/did:plc:") {
		http.NotFound(w, r)
		return
	}
	did := strings.SplitN(r.URL.Path, "/", 4)[3]

	resp, err := http.Get(fmt.Sprintf("https://plc.directory/%s", did))
	if err != nil {
		log.Fatal("plc.directory fetch failed: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("plc.directory fetch failed: ", err)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {
	var bind = os.Getenv("DID_PLC_DRIVER_BIND")
	if bind == "" {
		bind = ":8000"
	}

	http.HandleFunc("/", allRoutes)

	log.Println("Starting did:plc driver proxy: ", bind)
	if err := http.ListenAndServe(bind, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
