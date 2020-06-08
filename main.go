package main

import (
	"log"
	"net/http"
	"flag"
	"strings"
)

var (
	NetPort string	= ":80"
	FileDir string	= "."
)

func init() {
	flag.StringVar(&NetPort, "Port", ":80", "The network port that you want the webserver to serve out of.")
	flag.StringVar(&FileDir, "Dir", ".", "File Path that you want the webserver to serve. E.g. '.' or 'files'")
}

func main() {
	// Parse Flags
	flag.Parse()

	// Check NetPort String has Colon.
	if !strings.HasPrefix(NetPort, ":") {
		// Apply Colon if Missing as a prefix from NetPort.
		NetPort = ":" + NetPort
	}

	// Static File Server.
	log.Fatal(http.ListenAndServe(NetPort, http.FileServer(http.Dir(FileDir))))
}