package main

import "testing"

func TestCut(t *testing.T){
	testTable := []struct{
		input string
		sFlag bool
		dFlag string
		fFlag []int
		want string
	} {
		{
			input: "qwe\tasd\tzxc",
			sFlag: false,
			dFlag: "",
			fFlag: []int{1, 2},
			want:  "qwe asd",
		},
		{
			input: "qwe,asd,zxc",
			sFlag: false,
			dFlag: "",
			fFlag: []int{1, 2},
			want:  "qwe,asd,zxc",
		},
		{
			input: "qwe,asd,zxc",
			sFlag: false,
			dFlag: ",",
			fFlag: []int{1, 2},
			want:  "qwe asd",
		},
		{
			input: "Я строка для тестирования кодировки",
			sFlag: false,
			dFlag: " ",
			fFlag: []int{1, 3, 4},
			want:  "Я для тестирования",
		},
		{
			input: "qwe,asd,asd",
			sFlag: true,
			dFlag: "non-existent delimiter",
			fFlag: []int{1, 4},
			want:  "",
		},
		{
			input: "qwe,asd,asd",
			sFlag: true,
			dFlag: ",",
			fFlag: []int{1, 4},
			want:  "qwe",
		},
	}

	for _, testCase := range testTable{
		got := Cut(testCase.input, testCase.dFlag, testCase.fFlag, testCase.sFlag)
		if got != testCase.want{
			t.Errorf("want: %v, got: %v", testCase.want, got)
		}
	}
}
