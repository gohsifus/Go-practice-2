package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAnagramSet(t *testing.T) {
	testTable := []struct {
		data *[]string
		want *map[string]*[]string
	}{
		{
			data: &[]string{"Привет", "Пятка", "Пятак", "Тяпка", "Листок", "Слиток", "Столик", "А"},
			want: &map[string]*[]string{
				"листок": {"листок", "слиток", "столик"},
				"пятка":  {"пятак", "пятка", "тяпка"},
			},
		},
	}

	for _, testCase := range testTable {
		got := *getAnagramSet(testCase.data)

		if len(got) != len(*testCase.want) {
			t.Error()
		}

		for k, v := range *testCase.want{
			assert.EqualValues(t, v, got[k])
		}

	}
}
