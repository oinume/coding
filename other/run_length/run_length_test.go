package run_length

import (
	"reflect"
	"testing"
)

func Test_firstRunLength_Encode(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  []byte
	}{
		"A5B5": {
			input: []byte("AAAAABBBBB"),
			want:  []byte{'A', 5, 'B', 5},
		},
		"A5B5C5": {
			input: []byte("AAAAABBBBBCCCCC"),
			want:  []byte{'A', 5, 'B', 5, 'C', 5},
		},
		"A3B2C1": {
			input: []byte("AAABBC"),
			want:  []byte{'A', 3, 'B', 2, 'C', 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := &firstRunLength{}
			got := s.Encode(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("want=%v, got=%v", test.want, got)
			}

			gotDecode := s.Decode(test.want)
			if !reflect.DeepEqual(test.input, gotDecode) {
				t.Fatalf("want=%v, got=%v", test.input, gotDecode)
			}
		})
	}
}

func Test_secondRunLength_Encode(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  []byte
	}{
		"A5B5": {
			input: []byte("AAAAABBBBB"),
			want:  []byte{'A', 5, 'B', 5},
		},
		"A5B5C5": {
			input: []byte("AAAAABBBBBCCCCC"),
			want:  []byte{'A', 5, 'B', 5, 'C', 5},
		},
		"A3B2C1": {
			input: []byte("AAABBC"),
			want:  []byte{'A', 3, 'B', 2, 'C', 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := &secondRunLength{}
			got := s.Encode(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("want=%v, got=%v", test.want, got)
			}

			gotDecode := s.Decode(test.want)
			if !reflect.DeepEqual(test.input, gotDecode) {
				t.Fatalf("want=%v, got=%v", test.input, gotDecode)
			}
		})
	}
}
