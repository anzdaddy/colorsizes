package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitIntoThrees(t *testing.T) {
	assert.Equal(t, []string{}, splitIntoThrees(""))
	assert.Equal(t, []string{"1"}, splitIntoThrees("1"))
	assert.Equal(t, []string{"12"}, splitIntoThrees("12"))
	assert.Equal(t, []string{"123"}, splitIntoThrees("123"))
	assert.Equal(t, []string{"1", "234"}, splitIntoThrees("1234"))
	assert.Equal(t, []string{"12", "345"}, splitIntoThrees("12345"))
	assert.Equal(t, []string{"123", "456"}, splitIntoThrees("123456"))
	assert.Equal(t, []string{"1", "234", "567"}, splitIntoThrees("1234567"))
}
