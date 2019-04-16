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