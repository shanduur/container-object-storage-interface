# Developing Client Apps in Go

## Configuration Structure

The `Config` struct is the primary configuration object for the storage package. It encapsulates all necessary settings for interacting with different storage providers.

```go
// import "example.com/pkg/storage"
package storage

type Config struct {
	Spec Spec `json:"spec"`
}

type Spec struct {
	BucketName         string             `json:"bucketName"`
	AuthenticationType string             `json:"authenticationType"`
	Protocols          []string           `json:"protocols"`
	SecretS3           *s3.SecretS3       `json:"secretS3,omitempty"`
	SecretAzure        *azure.SecretAzure `json:"secretAzure,omitempty"`
}
```

## Azure Secret Structure

The `SecretAzure` struct holds authentication credentials for accessing Azure-based storage services.

```go
// import "example.com/pkg/storage/azure"
package azure

type SecretAzure struct {
	AccessToken     string    `json:"accessToken"`
	ExpiryTimestamp time.Time `json:"expiryTimeStamp"`
}
```

## S3 Secret Structure

The `SecretS3` struct holds authentication credentials for accessing S3-compatible storage services.

```go
// import "example.com/pkg/storage/s3"
package s3

type SecretS3 struct {
	Endpoint        string `json:"endpoint"`
	Region          string `json:"region"`
	AccessKeyID     string `json:"accessKeyID"`
	AccessSecretKey string `json:"accessSecretKey"`
}
```

## Factory

The [factory pattern](https://en.wikipedia.org/wiki/Factory_method_pattern) is used to instantiate the appropriate storage backend based on the provided configuration. We will hide the implementation behind the interface.

Here is a minimal interface that supports only basic `Delete`/`Get`/`Put` operations:

```go
type Storage interface {
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string, wr io.Writer) error
	Put(ctx context.Context, key string, data io.Reader, size int64) error
}
```

Our implementation of factory method can be defined as following:

```go
// import "example.com/pkg/storage"
package storage

import (
	"fmt"
	"slices"
	"strings"

	"example.com/pkg/storage/azure"
	"example.com/pkg/storage/s3"
)

func New(config Config, ssl bool) (Storage, error) {
	if slices.ContainsFunc(config.Spec.Protocols, func(s string) bool { return strings.EqualFold(s, "s3") }) {
		if !strings.EqualFold(config.Spec.AuthenticationType, "key") {
			return nil, fmt.Errorf("%w: invalid authentication type for s3", ErrInvalidConfig)
		}

		s3secret := config.Spec.SecretS3
		if s3secret == nil {
			return nil, fmt.Errorf("%w: s3 secret missing", ErrInvalidConfig)
		}

		return s3.New(config.Spec.BucketName, *s3secret, ssl)
	}

	if slices.ContainsFunc(config.Spec.Protocols, func(s string) bool { return strings.EqualFold(s, "azure") }) {
		if !strings.EqualFold(config.Spec.AuthenticationType, "key") {
			return nil, fmt.Errorf("%w: invalid authentication type for azure", ErrInvalidConfig)
		}

		azureSecret := config.Spec.SecretAzure
		if azureSecret == nil {
			return nil, fmt.Errorf("%w: azure secret missing", ErrInvalidConfig)
		}

		return azure.New(config.Spec.BucketName, *azureSecret)
	}

	return nil, fmt.Errorf("%w: invalid protocol (%v)", ErrInvalidConfig, config.Spec.Protocols)
}
```

## Clients

As we alredy defined the factory and uppermost configuration, let's get into the details of the clients, that will implement the `Storage` interface.

### S3

In the implementation of S3 client, we will use [MinIO](https://github.com/minio/minio-go) client library, as it's more lightweight than AWS SDK.

```go
// import "example.com/pkg/storage/s3"
package s3

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	s3cli      *minio.Client
	bucketName string
}

func New(bucketName string, s3secret SecretS3, ssl bool) (*Client, error) {
	s3cli, err := minio.New(s3secret.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3secret.AccessKeyID, s3secret.AccessSecretKey, ""),
		Region: s3secret.Region,
		Secure: ssl,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	return &Client{
		s3cli:      s3cli,
		bucketName: bucketName,
	}, nil
}

func (c *Client) Delete(ctx context.Context, key string) error {
	return c.s3cli.RemoveObject(ctx, c.bucketName, key, minio.RemoveObjectOptions{})
}

func (c *Client) Get(ctx context.Context, key string, wr io.Writer) error {
	obj, err := c.s3cli.GetObject(ctx, c.bucketName, key, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, obj)
	return err
}

func (c *Client) Put(ctx context.Context, key string, data io.Reader, size int64) error {
	_, err := c.s3cli.PutObject(ctx, c.bucketName, key, data, size, minio.PutObjectOptions{})
	return err
}
```

### Azure Blob

In the implementation of Azure client, we will use [Azure SDK](https://github.com/Azure/azure-sdk-for-go) client library. Note, that the configuration is done with `NoCredentials` client, as the Azure secret contains shared access signatures (SAS).

```go
// import "example.com/pkg/storage/azure"
package azure

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type Client struct {
	azCli         *azblob.Client
	containerName string
}

func New(containerName string, azureSecret SecretAzure) (*Client, error) {
	azCli, err := azblob.NewClientWithNoCredential(azureSecret.AccessToken, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	return &Client{
		azCli:         azCli,
		containerName: containerName,
	}, nil
}

func (c *Client) Delete(ctx context.Context, blobName string) error {
	_, err := c.azCli.DeleteBlob(ctx, c.containerName, blobName, nil)
	return err
}

func (c *Client) Get(ctx context.Context, blobName string, wr io.Writer) error {
	stream, err := c.azCli.DownloadStream(ctx, c.containerName, blobName, nil)
	if err != nil {
		return fmt.Errorf("unable to get download stream: %w", err)
	}
	_, err = io.Copy(wr, stream.Body)
	return err
}

func (c *Client) Put(ctx context.Context, blobName string, data io.Reader, size int64) error {
	_, err := c.azCli.UploadStream(ctx, c.containerName, blobName, data, nil)
	return err
}
```

## Summing up

Once all our code is in place, we can start using it in our app. Reading configuration is as simple as opening file and decoding it using standard `encoding/json` package.

```go
import (
	"encoding/json"
	"os"

	"example.com/pkg/storage"
)

func example() {
	f, err := os.Open("/opt/cosi/BucketInfo.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg storage.Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}

	client, err := storage.New(cfg, true)
	if err != nil {
		panic(err)
	}

	// use client Put/Get/Delete
	// ...
}
```
