package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULTBGCOLOR = "white"

	webTemplate = `
<!DOCTYPE html>
<html>
<body bgcolor="{{.BgColor}}" text="{{.TextColor}}">
<h1>Hostname: {{.Hostname}} </h1>
<h1>Namespace: {{.Namespace}}</h1>
</body>
</html>	
`
)

var bgColorFontColorMapping = map[string]string{
	"white":"black",
	"green":"cyan",
	"blue":"white",
	"red":"yellow"}

func main(){

	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").Parse(webTemplate)
	if err != nil{
		log.Fatalf("Unable to create template: %v",err)
	}
	settings := loadConfig()
	t.Execute(w, settings)
}

func loadConfig() *WebServerSettings{
	bgColor, ok := os.LookupEnv("BGCOLOR"); if ! ok {
		bgColor = DEFAULTBGCOLOR
	}
	textColor, ok := bgColorFontColorMapping[bgColor];if !ok {
		bgColor = DEFAULTBGCOLOR
		textColor = bgColorFontColorMapping[bgColor]
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Unable to get hostname: %v", err)
	}

	ns, ok := os.LookupEnv("NAMESPACE"); if ! ok {
		log.Fatalf("You have to specify NAMESPACE env variable")
	}

	return &WebServerSettings{
		Hostname:  hostname,
		BgColor:   bgColor,
		TextColor: textColor,
		Namespace: ns,
	}
}

type WebServerSettings struct {
	Hostname string
	BgColor string
	TextColor string
	Namespace string
}