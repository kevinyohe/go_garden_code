package main

import (
	"os"
	"text/template"
)

func main(){
type Inventory struct {
	Material string
	Count    uint
}
type Thing struct {
	Name string
	Noun string
}
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }


thing:= Thing{"Kevin", "cool man"}
net, err := template.New("adfa").Parse("{{ .Name }} is a {{ .Noun }}")
	if err != nil { panic(err) }
	err = net.Execute(os.Stdout, thing)
	if err != nil { panic(err) }
}