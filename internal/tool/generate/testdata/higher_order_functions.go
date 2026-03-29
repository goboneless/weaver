// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// EXPECTED
// func (s foo_client_stub) A(ctx context.Context, a0 func()) (err error) {
// panic("method foo.A has function-type arguments and can only be called on a local component")
// func (s foo_client_stub) B(ctx context.Context) (r0 func(), err error) {
// panic("method foo.B has function-type arguments and can only be called on a local component")
// func (s foo_server_stub) a(ctx context.Context, args []byte) (res []byte, err error) {
// panic("method foo.a has function-type arguments and can only be called on a local component")
// func (s foo_server_stub) b(ctx context.Context, args []byte) (res []byte, err error) {
// panic("method foo.b has function-type arguments and can only be called on a local component")

// Higher-order function support: component methods with function-type arguments
// or results. These methods can only be called locally (in-process).
package foo

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type foo interface {
	A(context.Context, func()) error
	B(context.Context) (func(), error)
}

type impl struct {
	weaver.Implements[foo]
}

func (impl) A(_ context.Context, fn func()) error { fn(); return nil }
func (impl) B(_ context.Context) (func(), error)  { return func() {}, nil }
