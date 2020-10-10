package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	mainTemplate = template.Must(template.New("main.go").Funcs(map[string]interface{}{"PackageName": PackageName}).Parse(`package main

import (
  "context"
  "fmt"
  "log"
  "os"
  "net"

  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "google.golang.org/grpc/status"
  "google.golang.org/grpc/codes"
  health "google.golang.org/grpc/health"
  healthpb "google.golang.org/grpc/health/grpc_health_v1"

  {{- range $index, $goImportPath := .GoImportPath }}
  {{ PackageName $goImportPath }} "{{ $goImportPath }}"
  {{- end }}
)

{{- range $index, $service := .Services}}
func New{{ $service.PackageName }}{{ $service.GoName }}() {{ $service.PackageName }}.{{ $service.GoName }}Server {
        return &{{ $service.StructName }}{}
}

type {{ $service.StructName }} struct {
}

{{ $structName := $service.StructName }}
{{- range $index, $method := $service.Methods}}
func (s *{{ $structName }}) {{ $method.GoName }}(ctx context.Context, req *{{ PackageName $method.RequestTypeGoImportPath }}.{{ $method.RequestTypeName }}) (*{{ PackageName $method.ResponseTypeGoImportPath }}.{{ $method.ResponseTypeName }}, error) {
  // TODO: Not yet implemented.
  return nil, status.Error(codes.Unimplemented, "Dummy")
}
{{- end}}
{{- end}}

func main() {
  port := os.Getenv("APP_PORT")
  if port == "" {
    fmt.Println("Please set APP_PORT")
    port = "5000"
  }
  lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  fmt.Printf("server listening at %v\n", lis.Addr())

  server := grpc.NewServer()
  healthpb.RegisterHealthServer(server, health.NewServer())
  {{- range $index, $service := .Services }}
  {{ $service.PackageName }}.Register{{$service.GoName}}Server(server, New{{ $service.PackageName }}{{ $service.GoName }}())
  {{- end }}
  reflection.Register(server)

  if err := server.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
`))
)

type Method struct {
	GoName                   string
	RequestTypeName          string
	ResponseTypeName         string
	RequestTypeGoImportPath  string
	ResponseTypeGoImportPath string
}

type Service struct {
	GoName       string
	Methods      []Method
	GoImportPath string
}

func (s Service) PackageName() string {
	return PackageName(s.GoImportPath)
}

func (s Service) StructName() string {
	return PackageName(s.GoImportPath) + string(s.GoName) + "Impl"
}

type ReflectionServerFile struct {
	ServiceRegistry []Service
	GoImportPath    []string
}

func PackageName(goImportPath string) string {
	uniqName := strings.Replace(strings.Replace(strings.Replace(goImportPath, "/", "_", -1), ".", "_", -1), "-", "_", -1)
	return uniqName
}

func (r *ReflectionServerFile) Content() (string, error) {
	type ReflectionServerFileData struct {
		Services     []Service
		GoImportPath []string
	}

	goImportPath := []string{}
	uniq := map[string]bool{}
	for _, v := range r.GoImportPath {
		_, ok := uniq[v]
		if ok {
			continue
		}
		uniq[v] = true
		goImportPath = append(goImportPath, v)
	}

	buf := bytes.NewBuffer([]byte(""))
	err := mainTemplate.Execute(buf, ReflectionServerFileData{Services: r.ServiceRegistry, GoImportPath: goImportPath})
	if err != nil {
		return "", errors.WithStack(err)
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		return "", errors.WithStack(err)
	}

	out, err = imports.Process("main.go", out, nil)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(out), nil
}

func main() {
	flag.Parse()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		Error(err)
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		Error(err)
	}

	gen, err := protogen.Options{}.New(req)
	if err != nil {
		Error(err)
	}

	rsf := ReflectionServerFile{}
	for _, f := range gen.Files {
		rsf.GoImportPath = append(rsf.GoImportPath, string(f.GoImportPath))
		for _, s := range f.Services {
			methods := make([]Method, len(s.Methods))
			for i, m := range s.Methods {
				methods[i] = Method{
					GoName:                   m.GoName,
					RequestTypeName:          m.Input.GoIdent.GoName,
					RequestTypeGoImportPath:  string(m.Input.GoIdent.GoImportPath),
					ResponseTypeName:         m.Output.GoIdent.GoName,
					ResponseTypeGoImportPath: string(m.Output.GoIdent.GoImportPath),
				}
			}
			service := Service{
				GoName:       s.GoName,
				GoImportPath: string(f.GoImportPath),
				Methods:      methods,
			}
			rsf.ServiceRegistry = append(rsf.ServiceRegistry, service)
		}
	}

	resp, err := RenderResponse(rsf)
	if err != nil {
		Error(err)
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		Error(err)
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		Error(err)
	}
}

func RenderResponse(rsf ReflectionServerFile) (*pluginpb.CodeGeneratorResponse, error) {
	var resp pluginpb.CodeGeneratorResponse

	mainContent, err := rsf.Content()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String("main.go"),
		Content: proto.String(mainContent),
	})

	return &resp, nil
}

func Debug(vs ...interface{}) {
	fmt.Fprintln(os.Stderr, vs)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
