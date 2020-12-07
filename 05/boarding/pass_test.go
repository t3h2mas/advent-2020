package boarding

import (
	"reflect"
	"testing"
)

func TestNewPass(t *testing.T) {
	type args struct {
		seat string
	}
	tests := []struct {
		name      string
		args      args
		want      *Pass
		wantError bool
	}{
		{
			name: "example one",
			args: args{
				seat: "BFFFBBFRRR",
			},
			want: &Pass{
				row: 70, col: 7,
			},
			wantError: false,
		},
		{
			name: "example two",
			args: args{
				seat: "FFFBBBFRRR",
			},
			want: &Pass{
				row: 14, col: 7,
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPass(tt.args.seat)

			haveError := err != nil

			if haveError != tt.wantError {
				t.Errorf("NewPass() unexpected error: %s", err.Error())
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
