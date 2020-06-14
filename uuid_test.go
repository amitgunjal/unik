package uuid

import "testing"

func TestNewUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUUID(); got != tt.want {
				t.Errorf("NewUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
