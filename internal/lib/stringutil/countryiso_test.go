package stringutil

import "testing"

func TestGetCountryCodeByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1: Vietnam",
			args: args{
				name: "Viet Nam",
			},
			want: "VN",
		},
		{
			name: "#1: Vietnam",
			args: args{
				name: "Vietnam",
			},
			want: "VN",
		},
		{
			name: "#1: Burma Myanmar",
			args: args{
				name: "Burma Myanmar",
			},
			want: "MM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCountryCodeByName(tt.args.name); got != tt.want {
				t.Errorf("GetCountryCodeByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
