package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("name", "johnny")

	form := New(postedData)
	v := form.MinLength("name", 2)
	if !v {
		t.Error("expected valid, got invalid")
	}

	isError := form.Errors.Get("name")
	if isError != "" {
		t.Error("expected valid, got invalid")
	}

	postedData = url.Values{}
	postedData.Add("name", "johnny")

	form = New(postedData)
	v = form.MinLength("name", 8)
	if v {
		t.Error("expected invalid, got valid")
	}

	isError = form.Errors.Get("name")
	if isError == "" {
		t.Error("expected invalid, got valid")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "a@b.c")

	form := New(postedData)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("expected valid, got invalid")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")

	form = New(postedData)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("expected invalid, got valid")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)

	v := form.Has("whatever")
	if v {
		t.Error("expected invalid, got valid")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)
	v = form.Has("a")
	if !v {
		t.Error("expected valid, got invalid")
	}
}

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("expected valid, got invalid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("expected invalid, got valid")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("expected valid, got invalid")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
