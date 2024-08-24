package genreflectionserver_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/pkg/errors"
	"github.com/rerost/giro/runner/genreflectionserver"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	testData = "testdata"
)

func request(testset string) (*pluginpb.CodeGeneratorRequest, error) {
	f, err := os.ReadFile(filepath.Join(testData, testset))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	set := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(f, &set)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req := new(pluginpb.CodeGeneratorRequest)

	for _, f := range set.GetFile() {
		req.ProtoFile = append(req.ProtoFile, f)
		req.FileToGenerate = append(req.FileToGenerate, f.GetName())
	}

	return req, nil
}

func TestRun(t *testing.T) {
	protos, err := os.ReadDir(testData)
	if err != nil {
		t.Error(err)
		return
	}

	for _, protoSetDir := range protos {
		protoSetDir := protoSetDir
		t.Run(protoSetDir.Name(), func(t *testing.T) {
			req, err := request(protoSetDir.Name())
			if err != nil {
				t.Error(err)
				return
			}

			resp, err := genreflectionserver.Run(req)
			if err != nil {
				t.Error(err)
				return
			}

			for _, f := range resp.GetFile() {
				f := f
				t.Run(f.GetName(), func(t *testing.T) {
					cupaloy.SnapshotT(t, f.GetContent())
				})
			}
		})
	}
}
