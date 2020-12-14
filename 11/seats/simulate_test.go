package seats

import (
	"reflect"
	"testing"
)

func Test_calculateOccupied(t *testing.T) {
	type args struct {
		layout Layout
	}
	tests := []struct {
		name string
		args args
		want Layout
	}{
		{
			name: "should calculate the next occupied seat state starting empty",
			args: args{
				layout: LayoutFrom(
					[]string{
						"L.LL",
						"LLLL",
						"L.L.",
					},
				),
			},
			want: LayoutFrom(
				[]string{
					"#.##",
					"####",
					"#.#.",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateOccupied(tt.args.layout); !reflect.DeepEqual(got, tt.want.occupied) {
				t.Errorf("calculateOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}
