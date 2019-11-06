package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Student struct {
	Name   string
	Age    int
	Emails []string
}

const tmpl = `{{$name := .Name}}
The name is {{$name}}.
{{range .Emails}}
	Myname is {{$name}} email id is {{.}}
{{end}}
`

func main() {
	// url := "http://localhost:1234/getTest"

	// proxyReq, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := &http.Client{}
	// proxyRes, err := client.Do(proxyReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer proxyRes.Body.Close()

	// bytes, _ := ioutil.ReadAll(proxyRes.Body)
	// str := string(bytes)
	// fmt.Println(str)

	s := Student{"Dennis", 32, []string{"jazmandorf", "thaeao", "wlqkrdl"}}

	t := template.New("person template")
	//te := template.p

	t, err := t.Parse(tmpl)
	t.ParseFiles()

	if err != nil {
		log.Fatal(err)
	}

	err1 := t.Execute(os.Stdout, s)

	if err1 != nil {
		log.Fatal(err1)
	}

	items := list.New()
	for _, x := range strings.Split("ABCDEFGH", "") {
		fmt.Println(x)
		items.PushBack(x)
	}

	e := items.PushFront(0)
	items.InsertAfter(1, e)

	for element := items.Front(); element != nil; element = element.Next() {
		fmt.Printf("%v ", element.Value)
	}
	filepath.Walk("./src/views/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		return nil
	})
}
