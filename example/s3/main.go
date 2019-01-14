package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cleardataeng/aidews/s3"
	"github.com/cleardataeng/aidews/s3/s3iface"
)

// DocInS3 is an example of any document you store in S3.
// In this case it looks like:
// {"Color": "blue"}
type DocInS3 struct {
	Color string

	Svc s3iface.Service `json:"-"`
}

// Get the contents from S3.
func (d *DocInS3) Get(k string) error {
	d.Init()
	return d.Svc.ReadUnmarshal(k, d)
}

// Init the service with defaults if needed.
func (d *DocInS3) Init() {
	if d.Svc == nil {
		d.Svc = s3.New("some-bucket", aws.String("us-west-2"), nil)
	}
}

func main() {
	d := DocInS3{}
	if err := d.Get("colordocs/blue"); err != nil {
		panic(err)
	}
}
