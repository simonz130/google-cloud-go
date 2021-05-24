// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dataqna

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dataqnapb "google.golang.org/genproto/googleapis/cloud/dataqna/v1alpha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newAutoSuggestionClientHook clientHook

// AutoSuggestionCallOptions contains the retry settings for each method of AutoSuggestionClient.
type AutoSuggestionCallOptions struct {
	SuggestQueries []gax.CallOption
}

func defaultAutoSuggestionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataqna.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataqna.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataqna.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAutoSuggestionCallOptions() *AutoSuggestionCallOptions {
	return &AutoSuggestionCallOptions{
		SuggestQueries: []gax.CallOption{},
	}
}

// internalAutoSuggestionClient is an interface that defines the methods availaible from Data QnA API.
type internalAutoSuggestionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SuggestQueries(context.Context, *dataqnapb.SuggestQueriesRequest, ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error)
}

// AutoSuggestionClient is a client for interacting with Data QnA API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This stateless API provides automatic suggestions for natural language
// queries for the data sources in the provided project and location.
//
// The service provides a resourceless operation suggestQueries that can be
// called to get a list of suggestions for a given incomplete query and scope
// (or list of scopes) under which the query is to be interpreted.
//
// There are two types of suggestions, ENTITY for single entity suggestions
// and TEMPLATE for full sentences. By default, both types are returned.
//
// Example Request:
//
// The service will retrieve information based on the given scope(s) and give
// suggestions based on that (e.g. “top item” for “top it” if “item” is a known
// dimension for the provided scope).
type AutoSuggestionClient struct {
	// The internal transport-dependent client.
	internalClient internalAutoSuggestionClient

	// The call options for this service.
	CallOptions *AutoSuggestionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AutoSuggestionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AutoSuggestionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AutoSuggestionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SuggestQueries gets a list of suggestions based on a prefix string.
// AutoSuggestion tolerance should be less than 1 second.
func (c *AutoSuggestionClient) SuggestQueries(ctx context.Context, req *dataqnapb.SuggestQueriesRequest, opts ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error) {
	return c.internalClient.SuggestQueries(ctx, req, opts...)
}

// autoSuggestionGRPCClient is a client for interacting with Data QnA API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type autoSuggestionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AutoSuggestionClient
	CallOptions **AutoSuggestionCallOptions

	// The gRPC API client.
	autoSuggestionClient dataqnapb.AutoSuggestionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAutoSuggestionClient creates a new auto suggestion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This stateless API provides automatic suggestions for natural language
// queries for the data sources in the provided project and location.
//
// The service provides a resourceless operation suggestQueries that can be
// called to get a list of suggestions for a given incomplete query and scope
// (or list of scopes) under which the query is to be interpreted.
//
// There are two types of suggestions, ENTITY for single entity suggestions
// and TEMPLATE for full sentences. By default, both types are returned.
//
// Example Request:
//
// The service will retrieve information based on the given scope(s) and give
// suggestions based on that (e.g. “top item” for “top it” if “item” is a known
// dimension for the provided scope).
func NewAutoSuggestionClient(ctx context.Context, opts ...option.ClientOption) (*AutoSuggestionClient, error) {
	clientOpts := defaultAutoSuggestionGRPCClientOptions()
	if newAutoSuggestionClientHook != nil {
		hookOpts, err := newAutoSuggestionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AutoSuggestionClient{CallOptions: defaultAutoSuggestionCallOptions()}

	c := &autoSuggestionGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		autoSuggestionClient: dataqnapb.NewAutoSuggestionServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *autoSuggestionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *autoSuggestionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *autoSuggestionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *autoSuggestionGRPCClient) SuggestQueries(ctx context.Context, req *dataqnapb.SuggestQueriesRequest, opts ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SuggestQueries[0:len((*c.CallOptions).SuggestQueries):len((*c.CallOptions).SuggestQueries)], opts...)
	var resp *dataqnapb.SuggestQueriesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.autoSuggestionClient.SuggestQueries(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
