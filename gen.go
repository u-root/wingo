// +build ignore

// This program generates datafiles.go
package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

func DataFile(n string) []byte {
	d, err := ioutil.ReadFile(n)
	die(err)
	return d
}

func main() {
	f, err := os.Create("misc/datafiles.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp     time.Time
		DejavusansTTF []byte
		WingoWav []byte
		WingoPng []byte
		ClosePng []byte
		MinimizePng []byte
		MaximizePng []byte
	}{
		Timestamp:     time.Now(),
		DejavusansTTF: DataFile("data/DejaVuSans.ttf"),
		WingoWav:      DataFile("data/wingo.wav"),
		WingoPng:      DataFile("data/wingo.png"),
		ClosePng:      DataFile("data/close.png"),
		MinimizePng:   DataFile("data/minimize.png"),
		MaximizePng:   DataFile("data/maximize.png"),
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// These are data files built in to wingo.
// {{ .Timestamp }}
// 

package misc

var (
	DejavusansTTF = []byte( {{with $x := .DejavusansTTF}}{{printf "%q" $x}}{{end}})
	WingoWav = []byte( {{with $x := .WingoWav}}{{printf "%q" $x}}{{end}})
	WingoPng = []byte( {{with $x := .WingoPng}}{{printf "%q" $x}}{{end}})
	ClosePng = []byte( {{with $x := .ClosePng}}{{printf "%q" $x}}{{end}})
	MinimizePng = []byte( {{with $x := .MinimizePng}}{{printf "%q" $x}}{{end}})
	MaximizePng = []byte( {{with $x := .MaximizePng}}{{printf "%q" $x}}{{end}})
)
`))