package structutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containerssh/structutils"
)

type mergeTest struct {
	Sub *mergeTestSub
}

type mergeTestSub struct {
	Text string
}

func TestMerge(t *testing.T) {
	data := new(mergeTest)
	newData := new(mergeTest)
	newData.Sub = new(mergeTestSub)
	newData.Sub.Text = "Hello world!"

	err := structutils.Merge(data, newData)
	assert.Nil(t, err, "failed to merge data (%v)", err)

	assert.Equal(t, "Hello world!", data.Sub.Text)
}

func TestMergeExisting(t *testing.T) {
	data := new(mergeTest)
	data.Sub = new(mergeTestSub)
	data.Sub.Text = "Foo"
	newData := new(mergeTest)
	newData.Sub = new(mergeTestSub)
	newData.Sub.Text = "Hello world!"

	err := structutils.Merge(data, newData)
	assert.Nil(t, err, "failed to merge data (%v)", err)

	assert.Equal(t, "Hello world!", data.Sub.Text)
}

func TestMergeExistingEmpty(t *testing.T) {
	data := new(mergeTest)
	data.Sub = new(mergeTestSub)
	data.Sub.Text = "Foo"
	newData := new(mergeTest)

	err := structutils.Merge(data, newData)
	assert.Nil(t, err, "failed to merge data (%v)", err)

	assert.Equal(t, "Foo", data.Sub.Text)
}

func TestMergeExistingEmpty2(t *testing.T) {
	data := new(mergeTest)
	data.Sub = new(mergeTestSub)
	data.Sub.Text = "Foo"
	newData := new(mergeTest)
	newData.Sub = new(mergeTestSub)

	err := structutils.Merge(data, newData)
	assert.Nil(t, err, "failed to merge data (%v)", err)

	assert.Equal(t, "Foo", data.Sub.Text)
}