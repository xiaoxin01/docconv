package docconv

import (
	"os"
	"strings"
	"testing"
)

func TestConvertTrimsSpace(t *testing.T) {
	resp, err := Convert(
		strings.NewReader(" \n\n\nthe \n file\n\n"),
		"text/plain",
		false,
	)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	if want := "the \n file"; resp.Body != want {
		t.Errorf("body = %v, want %v", resp.Body, want)
	}
}

func TestMaxWord(t *testing.T) {
	t.Run("max word not set", func(t *testing.T) {
		checkMaxWord := checkMaxWord()
		if checkMaxWord != false {
			t.Fatalf("got %v, want false", checkMaxWord)
		}
	})
	t.Run("test checkMaxWord", func(t *testing.T) {
		SetConfig(Config{MaxWord: 10})
		checkMaxWord := checkMaxWord()
		if checkMaxWord != true {
			t.Fatalf("got %v, want true", checkMaxWord)
		}
	})
	t.Run("test maxWordExceed", func(t *testing.T) {
		SetConfig(Config{MaxWord: 10})
		exceed := maxWordExceed(10)
		if exceed != false {
			t.Fatalf("got %v, want false", exceed)
		}
		exceed = maxWordExceed(11)
		if exceed != true {
			t.Fatalf("got %v, want true", exceed)
		}
	})
	t.Run("test parse pptx with maxword", func(t *testing.T) {
		SetConfig(Config{MaxWord: 2})
		f, err := os.Open("./docx_test/testdata/sample_3.docx")
		if err != nil {
			t.Fatalf("got error = %v, want nil", err)
		}

		resp, _, err := ConvertDocx(f)
		if err != nil {
			t.Fatalf("got error = %v, want nil", err)
		}
		if want := "Content from docx file"; !strings.Contains(resp, want) {
			t.Errorf("expected %v to contains %v", resp, want)
		}
		if want := "second"; strings.Contains(resp, want) {
			t.Errorf("expected %v to not contains %v", resp, want)
		}
	})

}
