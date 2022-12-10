package orderedform

import "testing"

func TestPlaintextEncoder(t *testing.T) {
	form := new(Form)
	form.Add("b", "a")
	form.Add("a", "a")

	if form.Encode(PlaintextEncoder) != "b=a&a=a" {
		t.Fail()
	}
}

func TestURLEncoder(t *testing.T) {
	form := new(Form)
	form.Add("b", "https://test.com")
	form.Add("a", "a")

	t.Log(form.Encode(URLEncoder))
	if form.Encode(URLEncoder) != "b=https%3A%2F%2Ftest.com&a=a" {
		t.Fail()
	}
}
