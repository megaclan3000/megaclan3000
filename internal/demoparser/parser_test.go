package demoparser

import (
	"reflect"
	"testing"
)

func TestMyParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    Match
		wantErr bool
	}{
		{
			name: "Parse demo1 file",
			path: "testdata/demo1.dem",
			want: Match{
				Map: "de_mirage",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewMyParser()

			got, err := p.Parse(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
