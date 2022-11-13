package test

import (
	"IM/helper"
	"testing"
)

func TestRandName(t *testing.T){
	name:=helper.RandName()
	t.Log(name)

}
func TestGenerateRandInt(t *testing.T){
	i:=helper.GenerateRandInt(0,10)
	t.Log(i)
}