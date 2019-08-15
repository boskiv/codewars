package roman

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		roman string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"XXI",args{"XXI"},21},
		{"I",args{"I"},1},
		{"IV",args{"IV"},4},
		{"MMVIII",args{"MMVIII"},2008},
		{"MDCLXVI",args{"MDCLXVI"},1666},
		{"XXI",args{"XXI"},21},
		{"MCMXC",args{"MCMXC"},1900},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.roman); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}