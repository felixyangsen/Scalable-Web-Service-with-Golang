// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// const port = ":8080"

// type ResponseData struct {
// 	Name    string
// 	Address string
// 	Age     int
// }

// const view string = `
// <html>
// 	<title>template</title>
// 	<body>
// 		<p>hello</p>
// 	</body>
// </html>
// `

// func main() {
// 	http.HandleFunc("/index", index)

// 	fmt.Println("running at http://localhost", port)
// 	http.ListenAndServe(port, nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	var tmpl = template.Must(template.New("main-template").Parse(view))
// 	if err := tmpl.Execute(w, nil); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
