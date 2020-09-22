package steamclient

import "testing"

func Test_divideStringFloats(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add more test cases.
		{
			name: "Divide two numbers",
			args: args{
				a: 15,
				b: 3,
			},
			want: "5.00",
		},

		{
			name: "Try to divide by zero",
			args: args{
				a: 15,
				b: 0,
			},
			want: "n/a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideNoZero(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divideStringFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}
