package strunpack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testCase1  = `a4bc2d5e2`
	testCase2  = `abcd`
	testCase3  = `45`
	testCase4  = `qwe\45`
	testCase5  = `qwe\4\5`
	testCase6  = `qwe\\5`
	testCase7  = `qwe11`
	testCase8  = `qwe\\\\\5`
	testCase9  = `qwe\511`
	testCase10 = `\\`
	testCase11 = `qwe\511s`
)

func TestStringParse(t *testing.T) {
	assert.Equal(t, "aaaabccdddddee", StringParse(testCase1))
	assert.Equal(t, "abcd", StringParse(testCase2))
	assert.Equal(t, "", StringParse(testCase3))
	assert.Equal(t, "qwe44444", StringParse(testCase4))
	assert.Equal(t, "qwe45", StringParse(testCase5))
	assert.Equal(t, `qwe\\\\\`, StringParse(testCase6))
	assert.Equal(t, `qweeeeeeeeeee`, StringParse(testCase7))
	assert.Equal(t, "qwe\\\\5", StringParse(testCase8))
	assert.Equal(t, "qwe55555555555", StringParse(testCase9))
	assert.Equal(t, `\`, StringParse(testCase10))
	assert.Equal(t, "qwe55555555555s", StringParse(testCase11))
}
