package routers

import "testing"

func TestSignCode(t *testing.T) {
	if code := newSignCode(); len(code) != 6 {
		t.Fatal(code)
	}
}

func BenchmarkSignCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(newSignCode())
	}
}

func TestTrimSignCode(t *testing.T) {
	d := [][]string{
		{"12345", "012345"},
		{"1234567", "123456"},
		{"123456", "123456"},
	}

	for _, a := range d {
		if trimSignCode(a[0], 6) != a[1] {
			t.Fatal(a)
		}
	}
}
