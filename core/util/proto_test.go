package util

import (
	"reflect"
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

type TestingStruct struct {
	V uint64
}

func TestToValue(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want *structpb.Value
	}{
		{
			name: "String Value", args: args{v: "paul aan"}, want: &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "paul aan"}},
		},
		{
			name: "int Value", args: args{v: 1}, want: &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: 1}},
		},
		{
			name: "List int Value",
			args: args{v: []int{1, 200}},
			want: &structpb.Value{
				Kind: &structpb.Value_ListValue{
					ListValue: &structpb.ListValue{
						Values: []*structpb.Value{
							{Kind: &structpb.Value_NumberValue{NumberValue: 1}},
							{Kind: &structpb.Value_NumberValue{NumberValue: 200}},
						},
					},
				},
			},
		},
		{
			name: "unit Value", args: args{v: uint(2)}, want: &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: 2}},
		},
		{
			name: "List uint Value",
			args: args{v: []uint64{1, 256, 65536, 4294967296, 18446744073709551614}},
			want: &structpb.Value{
				Kind: &structpb.Value_ListValue{
					ListValue: &structpb.ListValue{
						Values: []*structpb.Value{
							{Kind: &structpb.Value_NumberValue{NumberValue: 1}},
							{Kind: &structpb.Value_NumberValue{NumberValue: 256}},
							{Kind: &structpb.Value_NumberValue{NumberValue: 65536}},
							{Kind: &structpb.Value_NumberValue{NumberValue: 4294967296}},
							{Kind: &structpb.Value_NumberValue{NumberValue: 18446744073709551614}},
						},
					},
				},
			},
		},
		{
			name: "bool Value", args: args{v: true}, want: &structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: true}},
		},
		{
			name: "List string Value",
			args: args{v: []string{"paul", "aan"}},
			want: &structpb.Value{
				Kind: &structpb.Value_ListValue{
					ListValue: &structpb.ListValue{
						Values: []*structpb.Value{
							{Kind: &structpb.Value_StringValue{StringValue: "paul"}},
							{Kind: &structpb.Value_StringValue{StringValue: "aan"}},
						},
					},
				},
			},
		},
		{
			name: "Struct Value",
			args: args{v: TestingStruct{V: 184467440737095516}},
			want: &structpb.Value{
				Kind: &structpb.Value_StructValue{
					StructValue: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"V": {Kind: &structpb.Value_NumberValue{NumberValue: 184467440737095516}},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueProto(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStruct(t *testing.T) {
	type args struct {
		V map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want *structpb.Struct
	}{
		{
			name: "Metric tagging family",
			args: args{
				V: map[string]interface{}{"name": "Paul", "age": 20},
			},
			want: &structpb.Struct{Fields: map[string]*structpb.Value{
				"name": {Kind: &structpb.Value_StringValue{StringValue: "Paul"}},
				"age":  {Kind: &structpb.Value_NumberValue{NumberValue: 20}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructProto(tt.args.V); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
