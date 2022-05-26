package Toma

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		summ     int
		currency string
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		{
			name: "Test1",
			args: args{
				summ:     0,
				currency: "TJS",
			},

			wantMin: 0,
			wantMax: 0,
		},
		{
			name: "Test2",
			args: args{
				summ:     1000000,
				currency: "TJS",
			},

			wantMin: 3333,
			wantMax: 5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := Calculate(tt.args.summ, tt.args.currency)
			if gotMin != tt.wantMin {
				t.Errorf("Calculate() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("Calculate() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
