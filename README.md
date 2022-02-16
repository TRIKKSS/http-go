# HTTP-GO

***WHAT IS HTTP-GO***

HTTP-GO is a simple tool made with golang to create http servers and to serve file or directory like the python module http.server.

## ***USAGE***

```
git clone https://github.com/TRIKKSS/http-go.git
cd http-go
go build
```
### *help*

```
  -d string
        dir to serve
  -f string
        file to serve
  -p int
        port (default 80)
  -r string
        redirection
```

<ul>
	<li> -r option will make a redirection of / -> url you put in the parameter. Only the / handler will be redirected. </li>
	<li> -f option will create an handler /filename with the content of your file. </li>
	<li> -d option will serve a directory like the http.server module in python. </li>
</ul>

## ***EXAMPLE***

### *SERVE A FILE :*

![serve file image]("img/serve_file.png")

### *SERVE A DIRECTORY :*

![serve dir image]("img/serve_dir.png")

### *REDIRECTION :*

![redirect image]("img/redirect.png")