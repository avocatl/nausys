package ns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNausysDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"unmarshal fails with invalid date format",
			args{b: []byte("2020.06.01")},
			true,
		},
		{
			"unmarshal is successful",
			args{b: []byte("01.06.2020")},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &NausysDate{}
			err := d.UnmarshalJSON(tt.args.b)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
