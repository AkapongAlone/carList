package usecases

import (
	"testing"
	"time"
)


func Test_checkTimeToReply(t *testing.T) {
	tests := []struct {
		name        string
		carListTime time.Time
		want        bool
	}{
		{
			name:        "ten years age",
			carListTime: time.Now().AddDate(-10, 0, 0),
			want:        false,
		},
		{
			name:        "Exactly one year ago",
			carListTime: time.Now().AddDate(-1, 0, 0),
			want:        false,
		},
		{
			name:        "More than one year ago",
			carListTime: time.Now().AddDate(-1, -1, 0),
			want:        false,
		},
		{
			name:        "Just one day before one year",
			carListTime: time.Now().AddDate(-1, 0, 1),
			want:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTimeToReply(tt.carListTime); got != tt.want {
				t.Errorf("checkTimeToReply() = %v, want %v", got, tt.want)
			}
		})
	}
}
