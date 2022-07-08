package repo

import (
	"reflect"
	"testing"
)

func TestBahRam(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name    string
		b       BahRamAdapter
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "BahRam",
			b:    BahRamAdapter{},
			args: args{
				n: "5",
			},
			want:    []string{"BahRam", "BahBahRamRam", "Zzz", "BahRamRamRamRam", "Zzz"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.BahRam(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("BahRamAdapter.BahRam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BahRamAdapter.BahRam() = %v, want %v", got, tt.want)
			}
		})
	}
}
