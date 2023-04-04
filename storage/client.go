// Copyright 2018 Google LLC
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

package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// NewClient creates a new Google Cloud Storage client.
// The default scope is ScopeFullControl. To use a different scope, like
// ScopeReadOnly, use option.WithScopes.
//
// Clients should be reused instead of created as needed. The methods of Client
// are safe for concurrent use by multiple goroutines.
//
// You may configure the client by passing in options from the [google.golang.org/api/option]
// package. You may also use options defined in this package, such as [WithJSONReads].
func NewClient(ctx context.Context, opts ...option.ClientOption) (Client, error) {
	c, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return client{c}, nil
}

type (
	client         struct{ *storage.Client }
	bucketHandle   struct{ *storage.BucketHandle }
	objectHandle   struct{ *storage.ObjectHandle }
	bucketIterator struct{ *storage.BucketIterator }
	objectIterator struct{ *storage.ObjectIterator }
	reader         struct{ *storage.Reader }
	writer         struct{ *storage.Writer }
	copier         struct{ *storage.Copier }
	composer       struct{ *storage.Composer }
	aclHandle      struct{ *storage.ACLHandle }
)

func (client) embedToIncludeNewMethods()         {}
func (bucketHandle) embedToIncludeNewMethods()   {}
func (objectHandle) embedToIncludeNewMethods()   {}
func (bucketIterator) embedToIncludeNewMethods() {}
func (objectIterator) embedToIncludeNewMethods() {}
func (writer) embedToIncludeNewMethods()         {}
func (reader) embedToIncludeNewMethods()         {}
func (copier) embedToIncludeNewMethods()         {}
func (composer) embedToIncludeNewMethods()       {}
func (aclHandle) embedToIncludeNewMethods()      {}

func (c client) Bucket(name string) BucketHandle {
	return bucketHandle{c.Client.Bucket(name)}
}

func (c client) Buckets(ctx context.Context, projectID string) BucketIterator {
	return bucketIterator{c.Client.Buckets(ctx, projectID)}
}

func (b bucketHandle) Object(name string) ObjectHandle {
	return objectHandle{b.BucketHandle.Object(name)}
}

func (b bucketHandle) If(conds storage.BucketConditions) BucketHandle {
	return bucketHandle{b.BucketHandle.If(conds)}
}

func (b bucketHandle) Objects(ctx context.Context, q *Query) ObjectIterator {
	return objectIterator{b.BucketHandle.Objects(ctx, q)}
}

func (b bucketHandle) DefaultObjectACL() ACLHandle {
	return aclHandle{b.BucketHandle.DefaultObjectACL()}
}

func (b bucketHandle) ACL() ACLHandle {
	return aclHandle{b.BucketHandle.ACL()}
}

func (b bucketHandle) UserProject(projectID string) BucketHandle {
	return bucketHandle{b.BucketHandle.UserProject(projectID)}
}

func (bi bucketIterator) SetPrefix(s string) {
	bi.BucketIterator.Prefix = s
}

func (o objectHandle) ACL() ACLHandle {
	return aclHandle{o.ObjectHandle.ACL()}
}

func (o objectHandle) Generation(gen int64) ObjectHandle {
	return objectHandle{o.ObjectHandle.Generation(gen)}
}

func (o objectHandle) If(conds Conditions) ObjectHandle {
	return objectHandle{o.ObjectHandle.If(conds)}
}

func (o objectHandle) Key(encryptionKey []byte) ObjectHandle {
	return objectHandle{o.ObjectHandle.Key(encryptionKey)}
}

func (o objectHandle) ReadCompressed(compressed bool) ObjectHandle {
	return objectHandle{o.ObjectHandle.ReadCompressed(compressed)}
}

func (o objectHandle) NewReader(ctx context.Context) (Reader, error) {
	r, err := o.ObjectHandle.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return reader{r}, nil
}

func (o objectHandle) NewRangeReader(ctx context.Context, offset, length int64) (Reader, error) {
	r, err := o.ObjectHandle.NewRangeReader(ctx, offset, length)
	if err != nil {
		return nil, err
	}
	return reader{r}, nil
}

func (o objectHandle) NewWriter(ctx context.Context) Writer {
	return writer{o.ObjectHandle.NewWriter(ctx)}
}

func (o objectHandle) CopierFrom(src ObjectHandle) Copier {
	return copier{o.ObjectHandle.CopierFrom(src.(objectHandle).ObjectHandle)}
}

func (o objectHandle) ComposerFrom(srcs ...ObjectHandle) Composer {
	objs := make([]*storage.ObjectHandle, len(srcs))
	for i, s := range srcs {
		objs[i] = s.(objectHandle).ObjectHandle
	}
	return composer{o.ObjectHandle.ComposerFrom(objs...)}
}

func (w writer) ObjectAttrs() *storage.ObjectAttrs {
	return &w.Writer.ObjectAttrs
}

func (w writer) SetChunkSize(s int) {
	w.ChunkSize = s
}

func (w writer) SetProgressFunc(f func(int64)) {
	w.ProgressFunc = f
}

func (w writer) SetCRC32C(c uint32) {
	w.CRC32C = c
	w.SendCRC32C = true
}

func (c copier) ObjectAttrs() *ObjectAttrs {
	return &c.Copier.ObjectAttrs
}

func (c copier) SetRewriteToken(t string) {
	c.RewriteToken = t
}

func (c copier) SetProgressFunc(f func(copiedBytes, totalBytes uint64)) {
	c.ProgressFunc = f
}

func (c copier) SetDestinationKMSKeyName(k string) {
	c.DestinationKMSKeyName = k
}

func (c composer) ObjectAttrs() *ObjectAttrs {
	return &c.Composer.ObjectAttrs
}

type Query = storage.Query
type Conditions = storage.Conditions
type Notification = storage.Notification

type ObjectAttrs = storage.ObjectAttrs
type ObjectAttrsToUpdate = storage.ObjectAttrsToUpdate

type BucketAttrs = storage.BucketAttrs
type BucketAttrsToUpdate = storage.BucketAttrsToUpdate

type ACLRole = storage.ACLRole
type ACLRule = storage.ACLRule
type ACLEntity = storage.ACLEntity
