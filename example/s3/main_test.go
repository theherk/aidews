package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_s3 "github.com/cleardataeng/aidews/example/s3/mocks"
)

//go:generate mockgen -destination=mocks/mock.go github.com/cleardataeng/aidews/s3/s3iface Service

func TestDocInS3_Get(t *testing.T) {
	tt := []struct {
		name  string
		color string
		key   string
		err   error
	}{
		{"item found", "blue", "colordocs/blue", nil},
		{"item not found", "", "colordocs/red", fmt.Errorf("not found")},
	}
	for _, c := range tt {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockS3 := mock_s3.NewMockService(ctrl)
			mockS3.EXPECT().ReadUnmarshal(c.key, gomock.Any()).Do(func(k string, out *DocInS3) {
				switch k {
				case "colordocs/blue":
					out.Color = "blue"
				}
			}).Return(c.err)
			d := DocInS3{}
			d.Svc = mockS3
			err := d.Get(c.key)
			if err != c.err {
				t.Errorf("got: %v, want: %v", err, c.err)
			}
			if d.Color != c.color {
				t.Errorf("wrong color; got: %s, want: %s", d.Color, c.color)
			}
		})
	}
}
