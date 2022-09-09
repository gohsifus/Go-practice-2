package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	type res struct {
		unpackedString string
		err            bool
	}
	var testTable = []struct {
		data string
		want res
	}{
		{
			data: "a4bc2d5e",
			want: res{
				unpackedString: "aaaabccddddde",
				err:            false,
			},
		},
		{
			data: "abcd",
			want: res{
				unpackedString: "abcd",
				err:            false,
			},
		},
		{
			data: "45",
			want: res{
				unpackedString: "",
				err:            true,
			},
		},
		{
			data: "",
			want: res{
				unpackedString: "",
				err:            false,
			},
		},
		{
			data: "a2b2c3",
			want: res{
				unpackedString: "aabbccc",
				err:            false,
			},
		},
		{
			data: "a1b2c3",
			want: res{
				unpackedString: "abbccc",
				err:            false,
			},
		},
		{
			data: `a1b2\34c3`,
			want: res{
				unpackedString: "abb3333ccc",
				err:            false,
			},
		},
		{
			data: `qwe\4\5`,
			want: res{
				unpackedString: "qwe45",
				err:            false,
			},
		},
		{
			data: `qwe\\5`,
			want: res{
				unpackedString: `qwe\\\\\`,
				err:            false,
			},
		},
		{
			data: `qwe\45`,
			want: res{
				unpackedString: "qwe44444",
				err: false,
			},
		},
	}
	for _, testCase := range testTable {
		got, err := unpackWithEscape(testCase.data)
		if !assert.Equal(t, got, testCase.want.unpackedString) {
			t.Errorf("ожидалось %v %v, получено %v %v", testCase.want.unpackedString, testCase.want.err, got, err)
		}

		if testCase.want.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
