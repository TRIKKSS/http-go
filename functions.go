package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
    "path/filepath"
    "github.com/olekukonko/tablewriter"
)


func isDir(dirname string) bool {
    /*
    check if given name is a directory or file
    directory         -> true
    file or not exist -> false
    */
    dir, err := os.Stat(dirname)
    if os.IsNotExist(err) {
        return false
    }
    if dir.IsDir() {
        return true
    } else {
        return false
    }
}


func printFile(w http.ResponseWriter, r *http.Request, filename string) error {
    /*
    print a file.
    */
    file_bytes, err := os.ReadFile(filename)
    if err != nil {
        fmt.Printf("[-] can't read file %s\n", filename)
        page404(w, r)
        return err
    }
    fmt.Printf("[+] %s --> 200 OK\n", r.URL.Path)
    fmt.Fprintf(w, string(file_bytes))
    return nil
}


func printParams(r *http.Request) {
    /*
    print post and get parameters of a request
    */
    if len(r.URL.Query()) != 0 {
        fmt.Println("\n[*] GET params were:")
        for key, _ := range r.URL.Query(){
            fmt.Printf("  - %s=%s\n", key, r.URL.Query().Get(key))
        }
    } else {
        fmt.Println("\n[!] no GET parameters")
    }
    err := r.ParseForm()
    if err != nil {
        fmt.Println("[-] cant obtain post parameters, error somewhere")
        return
    }
    
    if len(r.PostForm) != 0{
        fmt.Println("[*] POST params were:")
        for key, _ := range r.PostForm{
            fmt.Printf("  - %s=%s\n", key, r.PostForm.Get(key))
        }
    } else {
        fmt.Println("[!] no POST parameters")
    }
    fmt.Println()
}


func servDir(w http.ResponseWriter, r *http.Request) {
    /*
    serve a directory, inspired from the source code of python module http.server 
    */
    var path = strings.Split(r.URL.Path, "?")[0]
    path = strings.Split(path, "#")[0]
    path = filepath.Join(absolutePath, path)
    printFile(w, r, path)
    printParams(r)
    fmt.Println(split_req_output)
}


func page404(w http.ResponseWriter, r *http.Request) {
    /*
    page not found : status code -> 404 
    + ascii 404 image into the 404.txt file
    */
    var page_404 string
    file_bytes, err := os.ReadFile("404.txt")
    if err != nil {
        fmt.Println(err)
        page_404 = "404 page not found"
    } else {
        page_404 = string(file_bytes)
    }

    w.WriteHeader(404)
    fmt.Printf("[-] %s  -->  404 Not Found\n", r.URL.Path)
    fmt.Fprintf(w, page_404)
}


func infoStart(port string, redirection string, dir string, filename string) {
    /*
    create and print table
    */    
    data := [][]string{
        []string{"port", port},
    }
    table := tablewriter.NewWriter(os.Stdout)
    if redirection != "" {
        data = append(data, []string{"redirection", redirection})
    }
    if dir != "" {
        data = append(data, []string{"directory to serve", dir})
    }
    if filename != "" {
        data = append(data, []string{"file to serve", filename})
    }
    for _, v := range data {
        table.Append(v)
    }
    table.Render()
    fmt.Println()
}