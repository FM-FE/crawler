package main

import "testing"

func TestFindmax(t *testing.T)  {
	tests := []struct{
		s string
		i int}{
		{"aaa",1},
		{"abcdfe",6},
		{"aabbccd",2},
		{"abbbcde",4},
		{"一二三二一",3},
		{"旋律缓慢而忧伤",7},
	}

	for i, tt := range tests{
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.i {
			t.Errorf("> %d : got %d, actul %d",i , tt.i, actual)
		}
	}
}

func BenchmarkFindmax(b *testing.B)  {
	s := "旋律缓慢而忧伤"
	ans := 7

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("> %d : got %d, actul %d",i , ans, actual)
		}
	}



}