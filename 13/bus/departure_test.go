package bus

import (
	"reflect"
	"testing"
)

func TestClosestDeparture(t *testing.T) {
	buses := []Bus{
		{7},
		{13},
		{59},
		{31},
		{19},
	}
	time := 939
	want := Departure{
		busID: 59,
		time:  944,
	}
	if got := ClosestDeparture(buses, time); !reflect.DeepEqual(got, want) {
		t.Errorf("ClosestDeparture() = %v, want %v", got, want)
	}
}
