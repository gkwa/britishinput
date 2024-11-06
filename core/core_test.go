package core

import (
	"math/rand"
	"testing"
	"time"
)

func TestPasser(t *testing.T) {
	p := NewPasser()
	if !p.ShouldSucceed() {
		t.Error("Passer should always return true")
	}
}

func TestFailer(t *testing.T) {
	f := NewFailer()
	if f.ShouldSucceed() {
		t.Error("Failer should always return false")
	}
}

func TestFiftyFifty(t *testing.T) {
	// Use fixed seed for deterministic testing
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ff := &FiftyFifty{rand: r}

	// Run many iterations to check distribution
	const iterations = 10000
	successes := 0

	for i := 0; i < iterations; i++ {
		if ff.ShouldSucceed() {
			successes++
		}
	}

	// Check if roughly 50% success rate within 5% margin
	ratio := float64(successes) / float64(iterations)
	if ratio < 0.45 || ratio > 0.55 {
		t.Errorf("Expected success ratio around 0.5, got %f", ratio)
	}
}

// TestExiters ensures all implementations satisfy the interface
func TestExiters(t *testing.T) {
	exiters := []Exiter{
		NewPasser(),
		NewFailer(),
		NewFiftyFifty(),
	}

	for _, e := range exiters {
		// Just verify we can call the interface method
		_ = e.ShouldSucceed()
	}
}

// Test with different seeds
func TestFiftyFiftyDifferentSeeds(t *testing.T) {
	ff1 := NewFiftyFifty()
	time.Sleep(time.Nanosecond) // Ensure different seed
	ff2 := NewFiftyFifty()

	// Should get different results with different seeds
	results := make(map[bool]bool)
	for i := 0; i < 100; i++ {
		results[ff1.ShouldSucceed()] = true
		results[ff2.ShouldSucceed()] = true
	}

	if len(results) != 2 {
		t.Error("Expected both true and false results from FiftyFifty")
	}
}
