package table

import (
	"testing"
)

func TestDrawHeaders(t *testing.T) {
	testCase := []string{"hello world", "is", "journeying"}
	w := 3
	t1 := NewTable(testCase).SetWidth(w)
	got := t1.drawHeader(testCase)
	expected := "| hel | is  | jou |\n| lo  |     | rne |\n| wor |     | yin |\n| ld  |     | g   |\n"
	
	if got != expected {
		t.Errorf("expected\n\n%s\n\n got\n\n%s\n", expected, got)
	}
}

func TestDrawHeadersWithLessFirst(t *testing.T) {
	testCase := []string{"is", "journeying", "hello world"}
	w := 3
	t1 := NewTable(testCase).SetWidth(w)
	got := t1.drawHeader(testCase)
	expected := "| is  | jou | hel |\n|     | rne | lo  |\n|     | yin | wor |\n|     | g   | ld  |\n"
	
	if got != expected {
		t.Errorf("expected\n\n%s\n\n got\n\n%s\n", expected, got)
	}
}
