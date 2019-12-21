package stringutil

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimLeadingHeadf(t *testing.T) {
	type args struct {
		sa []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "#1: simeple",
			args: args{
				sa: []string{"", "1", "2"},
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimLeadingHead(tt.args.sa); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlValidateString(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name        string
		args        args
		wantMatched bool
	}{
		{
			name: "#1 Test validate url correct and not error with schema is http",
			args: args{
				pattern: PatternUrl,
				s:       "http://wwe.com/abc/xyz",
			},
			wantMatched: true,
		},
		{
			name: "#2 Test validate url correct and not error with schema is https",
			args: args{
				pattern: PatternUrl,
				s:       "https://wwe.com/abc/xyz",
			},
			wantMatched: true,
		},
		{
			name: "#3 Test validate url incorrect  ",
			args: args{
				pattern: PatternUrl,
				s:       "https://wwe.com/ abc/xyz",
			},
			wantMatched: false,
		},
		{
			name: "#4 Test validate url incorrect with use FTP",
			args: args{
				pattern: PatternUrl,
				s:       "ftp://wwe.com/abc/xyz",
			},
			wantMatched: false,
		},
		{
			name: "#5 Test validate url incorrect with miss //",
			args: args{
				pattern: PatternUrl,
				s:       "http:wwe.com/abc/xyz",
			},
			wantMatched: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched, _ := UrlValidateString(tt.args.pattern, tt.args.s)
			assert.Equal(t, tt.wantMatched, matched)

		})
	}
}

func TestDifferDomainAndPath(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "#1: Equal",
			args: args{
				left:  "https://google.com/left",
				right: "https://google.com/left",
			},
			want: false,
		},
		{
			name: "#1: Equal",
			args: args{
				left:  "https://google.com/left",
				right: "https://google.com/left?search=innete",
			},
			want: false,
		},
		{
			name: "#2: Diff because of HTTPS scheme",
			args: args{
				left:  "https://google.com/left",
				right: "http://google.com/left",
			},
			want: true,
		},
		{
			name: "#3: Diff because of path",
			args: args{
				left:  "https://google.com/left",
				right: "https://google.com/right",
			},
			want: true,
		},
		{
			name: "#4: Diff because of path",
			args: args{
				left:  "https://google.com/left",
				right: "https://google.com/right?search=google",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferDomainAndPath(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("DifferDomainAndPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutualString(t *testing.T) {
	type args struct {
		source      string
		destination string
		length      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1: Non exist",
			args: args{
				source:      "XXXX",
				destination: "YYYY",
				length:      2,
			},
			want: "",
		},
		{
			name: "#2: Non exist",
			args: args{
				source:      "XYXY",
				destination: "YYYY",
				length:      2,
			},
			want: "",
		},
		{
			name: "#3: GA ID Detect",
			args: args{
				source:      "<!-- Global site tag (gtag.js) - Google Ads: 783408815 -->\n<script async src=\"https://www.googletagmanager.com/gtag/js?id=AW-783408815\"></script>\n<script>\n  window.dataLayer = window.dataLayer || [];\n  function gtag(){dataLayer.push(arguments);}\n  gtag('js', new Date());\n\n  gtag('config', 'AW-783408815');\n</script>\n",
				destination: "<!-- Event snippet for BabyStep Landing conversion page -->\n<script>\n  gtag('event', 'conversion', {\n      'send_to': 'AW-783408815/kTtwCN60h7MBEK-9x_UC',\n      'value': 11.11,\n      'currency': 'PLN'\n  });\n</script>\n",
				length:      9,
			},
			want: "783408815",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MutualString(tt.args.source, tt.args.destination, tt.args.length); got != tt.want {
				t.Errorf("MutualString() = %v, want %v", got, tt.want)
			}
		})
	}
}
