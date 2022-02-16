package main

import (
    "fmt"
    "log"
    "net/http"
    "flag"
    "strconv"
    "path/filepath"
)

var (
    servingDir = false
    split_req_output = "      =======================      \n"
    absolutePath string
)

func main() {
    // flags
    redirection := flag.String("r", "", "redirection")
    port_tmp    := flag.Int("p", 80, "port")
    file        := flag.String("f", "", "file to serve")
    dir         := flag.String("d", "", "dir to serve")
    flag.Parse()
    port := strconv.Itoa(*port_tmp)
    
    fmt.Printf("\n[+] Server started at localhost:%s\n", port)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){mainHandler(w, r, *redirection)})

    if *file != "" {
        http.HandleFunc("/" + *file, func(w http.ResponseWriter, r *http.Request){servFile(w, r, file)})
    }

    if (*dir != "" && isDir(*dir)) {
        servingDir   = true
        absolutePath, _ = filepath.Abs(*dir)
    }

    fmt.Println()
    infoStart(port, *redirection, absolutePath, *file) // port string, redirection string, dir string, filename string
    log.Fatal(http.ListenAndServe(":" + port, nil))
}