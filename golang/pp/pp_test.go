package main

import (
	
	"bytes"
	"io"
	
	"testing"
)

func Test_copy(t *testing.T){

	buf1:=bytes.NewBuffer(make([]byte,6))
	t.Log(buf1.Bytes())
	buf1.WriteString("abc")
	t.Log(buf1.Bytes())
	
	buf2:=bytes.NewBuffer(make([]byte,6))


	
	io.Copy(buf2,buf1)
	t.Log("b1",buf1.Bytes())
	t.Log("b2",buf2.Bytes())
	
	buf1.WriteString("efg")
	t.Log("b1",buf1.Bytes())
	t.Log("b2",buf2.Bytes())
	io.Copy(buf2,buf1)
	t.Log("b1",buf1.Bytes())
	t.Log("b2",buf2.Bytes())

	io.Copy(buf2,buf1)
	t.Log(buf2.Bytes())
}