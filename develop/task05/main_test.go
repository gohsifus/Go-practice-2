package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrep(t *testing.T) {
	testTable := []struct {
		name          string
		data          []string
		target        string
		A, B, C       int
		c, i, v, F, n bool
		want          map[int]string
	}{
		{
			name:   "Обычый поиск подстроки",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "qw",
			A:      0,
			B:      0,
			C:      0,
			c:      false,
			i:      false,
			v:      false,
			F:      false,
			n:      false,
			want:   map[int]string{1: "\x1b[31mqw\x1b[0m^\x1b[31mqw\x1b[0mddff"},
		},
		{
			name:   "Регулярное выражение",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "^qw",
			A:      0,
			B:      0,
			C:      0,
			c:      false,
			i:      false,
			v:      false,
			F:      false,
			n:      false,
			want:   map[int]string{1: "\x1b[31mqw\x1b[0m^qwddff"},
		},
		{
			name:   "Точное совпадение (F)",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "^qw",
			A:      0,
			B:      0,
			C:      0,
			c:      false,
			i:      false,
			v:      false,
			F:      true,
			n:      false,
			want:   map[int]string{1: "qw\x1b[31m^qw\x1b[0mddff"},
		},
		{
			name:   "(A)",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "ABO",
			A:      1,
			B:      0,
			C:      0,
			c:      false,
			i:      false,
			v:      false,
			F:      false,
			n:      false,
			want:   map[int]string{0: "\x1b[31mABO\x1b[0mBA", 1: "qw^qwddff"},
		},
		{
			name:   "(B)",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "qw",
			A:      0,
			B:      1,
			C:      0,
			c:      false,
			i:      false,
			v:      false,
			F:      false,
			n:      false,
			want:   map[int]string{0: "ABOBA", 1: "\x1b[31mqw\x1b[0m^\x1b[31mqw\x1b[0mddff"},
		},
		{
			name:   "Игнорировать регистр (i)",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "Abo",
			A:      0,
			B:      0,
			C:      0,
			c:      false,
			i:      true,
			v:      false,
			F:      false,
			n:      false,
			want:   map[int]string{0: "\x1b[31mABO\x1b[0mBA"},
		},
		{
			name:   "Инвертировать (v)",
			data:   []string{"ABOBA", "qw^qwddff", "zxc"},
			target: "Abo",
			A:      0,
			B:      0,
			C:      0,
			c:      false,
			i:      true,
			v:      true,
			F:      false,
			n:      false,
			want:   map[int]string{1: "qw^qwddff", 2: "zxc"},
		},
	}

	for _, testCase := range testTable {
		got := grep(testCase.data, testCase.target, testCase.A, testCase.B, testCase.C, testCase.c, testCase.i, testCase.v, testCase.F, testCase.n)
		assert.Equal(t, testCase.want, got)
	}
}
