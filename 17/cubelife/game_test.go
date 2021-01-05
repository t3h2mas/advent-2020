package cubelife

import (
	"reflect"
	"testing"
)

func TestGameFromSeed(t *testing.T) {
	type args struct {
		seed string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "create a game with active cubes seeded from a string",
			args: args{
				seed: ".#.\n..#\n###",
			},
			want: Game{
				active: cubeSet{
					{0, 0, 0}: true,
					{1, 0, 0}: true,
					{2, 0, 0}: true,
					{2, 1, 0}: true,
					{1, 2, 0}: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GameFromSeed(tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameFromSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
