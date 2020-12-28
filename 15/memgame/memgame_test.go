package memgame

import "testing"

func TestSpokenAt(t *testing.T) {
	type args struct {
		startingNumbers []int
		turn            int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tiny example",
			args: args{
				startingNumbers: []int{0, 3, 6},
				turn:            5,
			},
			want: 3,
		},
		{
			name: "tiny example 2",
			args: args{
				startingNumbers: []int{0, 3, 6},
				turn:            7,
			},
			want: 1,
		},
		{
			name: "small example",
			args: args{
				startingNumbers: []int{0, 3, 6},
				turn:            10,
			},
			want: 0,
		},
		{
			name: "example 1",
			args: args{
				startingNumbers: []int{0, 3, 6},
				turn:            2020,
			},
			want: 436,
		},
		{
			name: "example 2",
			args: args{
				startingNumbers: []int{1, 3, 2},
				turn:            2020,
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				startingNumbers: []int{2, 1, 3},
				turn:            2020,
			},
			want: 10,
		},
		{
			name: "example 4",
			args: args{
				startingNumbers: []int{1, 2, 3},
				turn:            2020,
			},
			want: 27,
		},
		{
			name: "example 5",
			args: args{
				startingNumbers: []int{2, 3, 1},
				turn:            2020,
			},
			want: 78,
		},
		{
			name: "example 6",
			args: args{
				startingNumbers: []int{3, 2, 1},
				turn:            2020,
			},
			want: 438,
		},
		{
			name: "example 7",
			args: args{
				startingNumbers: []int{3, 1, 2},
				turn:            2020,
			},
			want: 1836,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpokenAt(tt.args.startingNumbers, tt.args.turn); got != tt.want {
				t.Errorf("SpokenAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
