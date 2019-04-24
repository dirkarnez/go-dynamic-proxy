package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

const pogo = "pogo"
const entity = "entity"

func NewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fmt.Sprintf("./%s", pogo))
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	//NewWatcher()
	//fmt.Println(os.Getenv("GOFILE"))
	//fmt.Println(os.Getenv("GOPACKAGE"))

	//wd, _ := os.Getwd()
	//
	//fmt.Println()
	//pkgInfo, err := build.ImportDir(wd, build.IgnoreVendor)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(pkgInfo.GoFiles)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Getenv("GOFILE"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		//ret, ok := n.(*ast.TypeSpec)
		//if ok {
		//	printer.Fprint(os.Stdout, fset, ret)
		//	return true
		//}

		d, ok := n.(*ast.StructType)
		fmt.Print(d.Struct)
		if ok {
			printer.Fprint(os.Stdout, fset, d)
			return true
		}
		return true
	})
}