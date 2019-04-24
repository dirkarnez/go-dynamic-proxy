# go-dynamic-proxy

### Reference
https://github.com/gogo/letmegrpc/blob/master/main.go

```
var mainStr = `package main
import tmpprotos "tmpprotos"
import "google.golang.org/grpc"
func main() {
	tmpprotos.Serve("` + *httpAddr + `", "` + *grpcAddr + `",
		tmpprotos.DefaultHtmlStringer,
		grpc.WithInsecure(), grpc.WithDecompressor(grpc.NewGZIPDecompressor()),
	)
}
`
	if err := ioutil.WriteFile(filepath.Join(cmdDir, "/main.go"), []byte(mainStr), 0777); err != nil {
		log.Fatalf("%s\n", err)
	}
	gorun := exec.Command("go", "run", "main.go")
	envs := os.Environ()
	for i, e := range envs {
		if strings.HasPrefix(e, "GOPATH") {
			envs[i] = envs[i] + ":" + tmpDir
		}
	}
	gorun.Env = envs
	gorun.Dir = cmdDir

    out, err := gorun.CombinedOutput()
    if err != nil {
        log.Fatalf("%s %s\n", string(out), err)
    }
```

```
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"testing"
)

func main() {
	tests := []testing.InternalTest{{"TestAst", TestAst}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}

func TestAst(t *testing.T) {

	source := `package a

// B comment
type B struct {
	// C comment
	C string
}`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", []byte(source), parser.ParseComments)
	if err != nil {
		t.Error(err)
	}

	v := &visitor{
		file: file,
	}
	ast.Walk(v, file)

	var output []byte
	buf := bytes.NewBuffer(output)
	if err := printer.Fprint(buf, fset, file); err != nil {
		t.Error(err)
	}

	expected := `package a

// B comment
type B struct {
	// C comment
	C string
	// D comment
	D int
	// E comment
	E float64
}
`
	if buf.String() != expected {
		t.Error(fmt.Sprintf("Test failed. Expected:\n%s\nGot:\n%s", expected, buf.String()))
	}

}

type visitor struct {
	file *ast.File
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return v
	}

	switch n := node.(type) {
	case *ast.GenDecl:
		if n.Tok != token.TYPE {
			break
		}
		ts := n.Specs[0].(*ast.TypeSpec)
		if ts.Name.Name == "B" {
			fields := ts.Type.(*ast.StructType).Fields
			addStructField(fields, v.file, "int", "D", "D comment")
			addStructField(fields, v.file, "float64", "E", "E comment")
		}
	}

	return v
}

func addStructField(fields *ast.FieldList, file *ast.File, typ string, name string, comment string) {
	c := &ast.Comment{Text: fmt.Sprint("// ", comment)}
	cg := &ast.CommentGroup{List: []*ast.Comment{c}}
	f := &ast.Field{
		Doc:   cg,
		Names: []*ast.Ident{ast.NewIdent(name)},
		Type:  ast.NewIdent(typ),
	}
	fields.List = append(fields.List, f)
	file.Comments = append(file.Comments, cg)
}
```
https://play.golang.org/p/RID4N30FZK