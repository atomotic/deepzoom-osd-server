package main

import (
	"archive/zip"
	"bitbucket.org/tebeka/nrsc"
	"fmt"
	"github.com/remyoudompheng/go-misc/zipfs"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func checkinputdir() {
	if _, err := os.Stat("./dzi"); os.IsNotExist(err) {
		fmt.Println("-- missing dzi directory, creating.")
		os.Mkdir("./dzi", 0775)
		return
	}
}

func dzi(res http.ResponseWriter, req *http.Request) {
	dzi := "dzi/" + strings.Split(html.EscapeString(req.URL.Path), "/")[2] + ".zip"

	z, err := zip.OpenReader(dzi)
	if err != nil {
		http.Error(res, err.Error(), 404)
		return
	}
	defer z.Close()

	http.StripPrefix("/dzi/", http.FileServer(zipfs.NewZipFS(&z.Reader))).ServeHTTP(res, req)
}

var templ = template.Must(template.New("index").Parse(indexTemplate))

func view(res http.ResponseWriter, req *http.Request) {
	dziID := strings.Split(html.EscapeString(req.URL.Path), "/")[2]
	dzi := "/dzi/" + dziID + "/" + dziID + ".dzi"
	templ.Execute(res, dzi)
}

func main() {
	checkinputdir()

	nrsc.Handle("/openseadragon/")
	http.HandleFunc("/dzi/", dzi)
	http.HandleFunc("/view/", view)

	fmt.Println("-- running on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const indexTemplate = `
<html>
<head><title>deepzoom-server {{.}}</title></head>
<body>
<div id="openseadragon1" style="width: 1024px; height: 768px;"></div>
<script src="/openseadragon/openseadragon.min.js"></script>
<script type="text/javascript">
    var viewer = OpenSeadragon({
        id: "openseadragon1",
        prefixUrl: "/openseadragon/images/",
        tileSources: "{{.}}"
    });
</script>
</body>
</html>
`
