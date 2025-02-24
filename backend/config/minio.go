package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewMinioClient(viper *viper.Viper) (*minio.Client, error) {
	accessKey := viper.GetString("minio.access_key")
	secretKey := viper.GetString("minio.secret_key")
	endpoint := viper.GetString("minio.endpoint")
	useSSL := viper.GetBool("minio.use_ssl")

	opts := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	}
	return minio.New(endpoint, opts)
}
