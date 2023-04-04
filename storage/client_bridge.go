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
	"io"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

//go:generate counterfeiter -generate

//counterfeiter:generate -o ./fake/client.go . Client

type Client interface {
	Bucket(name string) BucketHandle
	Buckets(ctx context.Context, projectID string) BucketIterator
	Close() error
}

//counterfeiter:generate -o ./fake/object_handle.go . ObjectHandle

type ObjectHandle interface {
	ACL() ACLHandle
	Generation(int64) ObjectHandle
	If(Conditions) ObjectHandle
	Key([]byte) ObjectHandle
	ReadCompressed(bool) ObjectHandle
	Attrs(context.Context) (*ObjectAttrs, error)
	Update(context.Context, ObjectAttrsToUpdate) (*ObjectAttrs, error)
	NewReader(context.Context) (Reader, error)
	NewRangeReader(context.Context, int64, int64) (Reader, error)
	NewWriter(context.Context) Writer
	Delete(context.Context) error
	CopierFrom(ObjectHandle) Copier
	ComposerFrom(...ObjectHandle) Composer
}

//counterfeiter:generate -o ./fake/bucket_handle.go . BucketHandle

type BucketHandle interface {
	Create(context.Context, string, *BucketAttrs) error
	Delete(context.Context) error
	DefaultObjectACL() ACLHandle
	Object(string) ObjectHandle
	Attrs(context.Context) (*BucketAttrs, error)
	Update(context.Context, BucketAttrsToUpdate) (*BucketAttrs, error)
	If(storage.BucketConditions) BucketHandle
	Objects(context.Context, *Query) ObjectIterator
	ACL() ACLHandle
	IAM() *iam.Handle
	UserProject(projectID string) BucketHandle
	Notifications(context.Context) (map[string]*Notification, error)
	AddNotification(context.Context, *Notification) (*Notification, error)
	DeleteNotification(context.Context, string) error
	LockRetentionPolicy(context.Context) error
}

//counterfeiter:generate -o ./fake/object_iterator.go . ObjectIterator

type ObjectIterator interface {
	Next() (*storage.ObjectAttrs, error)
	PageInfo() *iterator.PageInfo
}

//counterfeiter:generate -o ./fake/bucket_iterator.go . BucketIterator

type BucketIterator interface {
	SetPrefix(string)
	Next() (*storage.BucketAttrs, error)
	PageInfo() *iterator.PageInfo
}

//counterfeiter:generate -o ./fake/acl_handle.go . ACLHandle

type ACLHandle interface {
	Delete(context.Context, ACLEntity) error
	Set(context.Context, ACLEntity, ACLRole) error
	List(context.Context) ([]ACLRule, error)
}

//counterfeiter:generate -o ./fake/reader.go . Reader

type Reader interface {
	io.ReadCloser
	Size() int64
	Remain() int64
	ContentType() string
	ContentEncoding() string
	CacheControl() string
}

//counterfeiter:generate -o ./fake/writer.go . Writer

type Writer interface {
	io.WriteCloser
	ObjectAttrs() *ObjectAttrs
	SetChunkSize(int)
	SetProgressFunc(func(int64))
	SetCRC32C(uint32)
	CloseWithError(err error) error
	Attrs() *storage.ObjectAttrs
}

//counterfeiter:generate -o ./fake/copier.go . Copier

type Copier interface {
	ObjectAttrs() *ObjectAttrs
	SetRewriteToken(string)
	SetProgressFunc(func(uint64, uint64))
	SetDestinationKMSKeyName(string)
	Run(context.Context) (*ObjectAttrs, error)
}

//counterfeiter:generate -o ./fake/composer.go . Composer

type Composer interface {
	ObjectAttrs() *ObjectAttrs
	Run(context.Context) (*ObjectAttrs, error)
}
