package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	testTable := []struct {
		name string
		data  []string
		rFlag bool
		kFlag int
		nFlag bool
		uFlag bool
		want  []string
	}{
		{
			name: "Сортировка без флагов",
			data:  []string{"abs", "qwe", "bbb"},
			rFlag: false,
			kFlag: 1,
			nFlag: false,
			uFlag: false,
			want:  []string{"abs", "bbb", "qwe"},
		},
		{
			name: "Строки с числами выше",
			data:  []string{"abs", "qwe", "bbb", "1ccc"},
			rFlag: false,
			kFlag: 1,
			nFlag: false,
			uFlag: false,
			want:  []string{"1ccc", "abs", "bbb", "qwe"},
		},
		{
			name: "Сортировка с флагом k",
			data:  []string{"abs bfhb", "qwe fjk", "bbb abb", "1ccc coid"},
			rFlag: false,
			kFlag: 2,
			nFlag: false,
			uFlag: false,
			want:  []string{"bbb abb", "abs bfhb", "1ccc coid", "qwe fjk"},
		},
		{
			name: "Сортировка с флагом r",
			data:  []string{"abs bfhb", "qwe fjk", "bbb abb", "1ccc coid"},
			rFlag: true,
			kFlag: 2,
			nFlag: false,
			uFlag: false,
			want:  []string{"qwe fjk", "1ccc coid", "abs bfhb", "bbb abb"},
		},
		{
			name: "Сортировка с флагом n",
			data:  []string{"1", "11", "2", "22"},
			rFlag: false,
			kFlag: 1,
			nFlag: true,
			uFlag: false,
			want:  []string{"1", "2", "11", "22"},
		},
		{
			name: "Сортировка с флагом u",
			data:  []string{"1", "11", "22", "22"},
			rFlag: false,
			kFlag: 1,
			nFlag: true,
			uFlag: true,
			want:  []string{"1", "11", "22"},
		},
	}

	for _, testCase := range testTable {
		got := Sort(testCase.data, testCase.rFlag, testCase.kFlag, testCase.nFlag, testCase.uFlag)
		assert.Equal(t, testCase.want, got)
	}
}
