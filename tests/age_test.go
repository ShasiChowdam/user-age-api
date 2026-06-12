package tests

import (
	"testing"
	"time"

	"github.com/ShasiChowdam/user-age-api/internal/service"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	age := service.CalculateAge(dob)

	if age <= 0 {
		t.Errorf("expected age > 0, got %d", age)
	}
}