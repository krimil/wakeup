package main

import (
	"log"
	"net"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/mpolden/wakeup/http"
)

func main() {
	var opts struct {
		CacheFile string `short:"c" long:"cache" description:"Path to cache file" required:"true" value-name:"FILE"`
		SourceIP  string `short:"b" long:"bind" description:"IP address to bind to when sending WOL packets" value-name:"IP"`
		Listen    string `short:"l" long:"listen" description:"Listen address" value-name:"ADDR" default:":8580"`
		StaticDir string `short:"s" long:"static" description:"Path to directory containing static assets" value-name:"DIR"`
	}
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	sourceIP := net.ParseIP(opts.SourceIP)
	if opts.SourceIP != "" && sourceIP == nil {
		log.Fatalf("invalid ip: %s", opts.SourceIP)
	}

	server := http.New(opts.CacheFile)
	server.StaticDir = opts.StaticDir
	server.SourceIP = sourceIP
	if strings.HasPrefix(opts.Listen, ":") {
		log.Printf("Serving at http://0.0.0.0%s", opts.Listen)
	} else {
		log.Printf("Serving at http://%s", opts.Listen)
	}
	if err := server.ListenAndServe(opts.Listen); err != nil {
		log.Fatal(err)
	}
}
