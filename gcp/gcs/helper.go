package gcs

import (
	"context"
	"encoding/base64"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iam/v1"

	"github.com/yusuke0701/goutils/gcp"
)

// SignedURLType is the http method to use the signed urls.
type SignedURLType string

// SignedURLType list
const (
	SignedURLTypeGET SignedURLType = "GET"
	SignedURLTypePUT SignedURLType = "PUT"
)

// GetSignedURL get the signed urls by default service accounts.
// ref: https://cloud.google.com/storage/docs/access-control/signed-urls?hl=ja
func GetSignedURL(ctx context.Context, bucketName, fileName, mimeType string, method SignedURLType, expires time.Time) (string, error) {
	return storage.SignedURL(bucketName, fileName, &storage.SignedURLOptions{
		GoogleAccessID: gcp.DefaultServiceAccountName,
		Method:         string(method),
		Expires:        expires,
		ContentType:    mimeType,
		SignBytes: func(b []byte) ([]byte, error) {
			resp, err := gcp.IAMService.Projects.ServiceAccounts.SignBlob(
				gcp.DefaultServiceAccountID,
				&iam.SignBlobRequest{BytesToSign: base64.StdEncoding.EncodeToString(b)},
			).Context(ctx).Do()
			if err != nil {
				return nil, err
			}
			return base64.StdEncoding.DecodeString(resp.Signature)
		},
	})
}
