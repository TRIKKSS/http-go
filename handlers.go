package main

import (
    "fmt"
    "net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request, newUrl string) {
    /*
    main handler
    */
    fmt.Printf("[+] request received from %s\n", r.RemoteAddr)

    // if the url path is not / check if the -d flag is set
    // if it is set check if we can access to the file in the url
    // else return 404 not found
    if r.URL.Path != "/" {
        if servingDir == true {
            servDir(w, r)
            return
        } else {
            page404(w, r)
            fmt.Println(split_req_output) // split each request
            return
        }
    }

    // if the -r flag is set the / url path redirect to the value set into the -r flag
    // example : http-go -r https://google.com/ -p 8080
    //           http://localhost:8080  redirect to https://google.com
    if newUrl != ""{
        http.Redirect(w, r, newUrl, http.StatusSeeOther)
        fmt.Printf("[+] / --> 303 See Other\n[+] redirected to %s\n", newUrl)
    }
    
    printParams(r)
    
    fmt.Println(split_req_output)
}

func servFile(w http.ResponseWriter, r *http.Request, filename *string){
    /*
    handler create if the -f flag is set to serve a file.
    */
    fmt.Printf("[+] request received from %s\n", r.RemoteAddr)
    printFile(w, r, *filename)
    printParams(r)
    fmt.Println(split_req_output)
}