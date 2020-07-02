package docconv

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		maxword int
		file    string
		want    string
	}{
		{
			desc:    "test limit word for doc",
			maxword: 3,
			file:    "./test_data/simple.doc",
			want:    "1",
		},
		{
			desc:    "test limit word with chinese for doc",
			maxword: 14,
			file:    "./test_data/simple.doc",
			want:    "123456789壹",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			SetConfig(Config{Limitation: LenthLimitation{XMLMaxWord: tC.maxword}})
			res, err := ConvertPath(tC.file)
			if err != nil {
				t.Fatalf("got error = %v, want nil", err)
			}
			if want := tC.want; res.Body != want {
				t.Errorf("expected %v to eq %v", res.Body, want)
			}
		})
	}
}
