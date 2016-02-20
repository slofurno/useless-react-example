package main

import (
	"net/http"
	"time"
)

const html string = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div id="root"></div>
<script>var subs = [];
function push (n) {
	subs.forEach(x => x(n));	
}
</script>
<script src="http://192.168.1.104:8000/static/bundle.js"></script>
`

func stream(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte(html))
	f, _ := res.(http.Flusher)

	for {
		<-time.After(time.Second)
		res.Write([]byte(`<script>push("encoding completed successfully in 654ms");</script>`))

		f.Flush()
		<-time.After(time.Second)
		res.Write([]byte(`<script>push("encoding completed successfully in 555ms");</script>`))

		f.Flush()
		<-time.After(time.Second)
		res.Write([]byte(`<script>push("encoding completed successfully in 789ms");</script>`))

		f.Flush()
		<-time.After(time.Second)
		res.Write([]byte(`<script>push("error incorrect format specified or codec not installed");</script>`))

		f.Flush()
	}

}

func main() {

	http.HandleFunc("/", stream)
	http.ListenAndServe(":1234", nil)

}
