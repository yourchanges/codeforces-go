// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`["pes","fifa","gta","pes(2019)"]`, 
			`["pes","fifa","gta","pes(2019)"]`,
		},
		{
			`["gta","gta(1)","gta","avalon"]`, 
			`["gta","gta(1)","gta(2)","avalon"]`,
		},
		{
			`["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"]`, 
			`["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]`,
		},
		{
			`["wano","wano","wano","wano"]`, 
			`["wano","wano(1)","wano(2)","wano(3)"]`,
		},
		{
			`["kaido","kaido(1)","kaido","kaido(1)"]`, 
			`["kaido","kaido(1)","kaido(2)","kaido(1)(1)"]`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getFolderNames, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-194/problems/making-file-names-unique/