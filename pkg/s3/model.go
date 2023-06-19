package s3

type (
	Options struct {
		EndPoint         string
		Region           string
		CustomHost       string
		RemoveBucket     bool
		S3ForcePathStyle *bool
	}
	Object struct {
		Key  string
		Size int64
	}
)
