package bloomfilter_test

import (
	"reflect"
	"testing"
	"github.com/axamon/bloomfilter"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *BloomFilter
	}{
		// TODO: Add test cases.
		{name: "first", want: }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
