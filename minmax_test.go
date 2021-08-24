package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinMax(t *testing.T) {
	matrixArr1:=[]int{1}
	matrixArr2:=[]int{3,4,6,2,8,1,4,0,4}
	min,max:=GetMinMax(matrixArr1,3)
	assert.Equal(t, 1,min.Value)
	assert.Equal(t, 0,min.Index)
	assert.Equal(t, 1,min.Row)
	assert.Equal(t, 1,min.Column)
	assert.Equal(t, 1,max.Value)
	assert.Equal(t, 0,max.Index)
	assert.Equal(t, 1,max.Row)
	assert.Equal(t, 1,max.Column)
	min,max=GetMinMax(matrixArr2,3)
	assert.Equal(t, 0,min.Value)
	assert.Equal(t, 7,min.Index)
	assert.Equal(t, 3,min.Row)
	assert.Equal(t, 2,min.Column)
	assert.Equal(t, 8,max.Value)
	assert.Equal(t, 4,max.Index)
	assert.Equal(t, 2,max.Row)
	assert.Equal(t, 2,max.Column)

}
