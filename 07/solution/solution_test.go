package solution

import (
	"reflect"
	"testing"
)

func TestLineParts(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "it parses a line with one contains clause",
			args: args{
				line: "bright white bags contain 1 shiny gold bag.",
			},
			want: [][]string{
				{"bright", "white", "bags"},
				{"1", "shiny", "gold", "bag."},
			},
		},
		{
			name: "it parses a line with two contains clauses",
			args: args{
				line: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			},
			want: [][]string{
				{"light", "red", "bags"},
				{"1", "bright", "white", "bag"},
				{"2", "muted", "yellow", "bags."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineParts(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBagClauses(t *testing.T) {
	type args struct {
		parts [][]string
	}
	tests := []struct {
		name string
		args args
		want *Bag
	}{
		{
			name: "gets clauses for a bag",
			args: args{
				[][]string{
					{"light", "red", "bags"},
					{"1", "bright", "white", "bag"},
					{"2", "muted", "yellow", "bags."},
				},
			},
			want: &Bag{
				Color: "light red",
				Contains: []string{
					"bright white",
					"muted yellow",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BagClauses(tt.args.parts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BagClauses() = %v, want %v", got, tt.want)
			}
		})
	}
}
