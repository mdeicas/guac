//
// Copyright 2024 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent"
	gen "github.com/guacsec/guac/pkg/guacrest/generated"
)

// EntConnectedServer implements the REST API interface, using by default the
// GrapQL API Server as a backend, but also allows overriding the default
// handlers to ones that directly use the ENT backend.
type EntConnectedServer struct {
	ent *ent.Client
	*DefaultServer
}

func NewEntConnectedServer(ent *ent.Client, gqlClient graphql.Client) *EntConnectedServer {
	return &EntConnectedServer{
		DefaultServer: NewDefaultServer(gqlClient),
	}
}

func (s *EntConnectedServer) AnalysisDependencies(ctx context.Context, request gen.AnalysisDependenciesRequestObject) (gen.AnalysisDependenciesResponseObject, error) {
	return nil, fmt.Errorf("AnalysisDependencies unimplemented for Ent")
}
