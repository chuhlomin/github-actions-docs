package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessDir(t *testing.T) {
	app, err := newApp("template.md")
	assert.NoError(t, err, "newApp should not return error")

	dir := "testdata/action"
	err = app.ProcessDir(dir, "README.md")
	assert.NoError(t, err, "ProcessDir should not return error")

	expected, err := ioutil.ReadFile("testdata/action/README_expected.md")
	assert.NoError(t, err, "ReadFile should not return error (expected)")

	actual, err := ioutil.ReadFile(dir + "/README.md")
	assert.NoError(t, err, "ReadFile should not return error (actial)")

	assert.Equal(t, string(expected), string(actual))
}
