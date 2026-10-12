package server

import "testing"

func TestPaginate(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	tests := []struct {
		name          string
		offset, limit int
		want          []int
	}{
		{"first page", 0, 2, []int{1, 2}},
		{"middle page", 2, 2, []int{3, 4}},
		{"last partial page", 4, 2, []int{5}},
		{"offset past end", 5, 2, nil},
		{"negative offset is clamped to zero", -1, 2, []int{1, 2}},
		{"negative limit yields empty page", 0, -3, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := paginate(items, tt.offset, tt.limit)
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
