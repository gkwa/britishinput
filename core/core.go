package core

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Exiter interface {
	ShouldSucceed() bool
}

type FiftyFifty struct {
	rand *rand.Rand
}

func NewFiftyFifty() *FiftyFifty {
	return &FiftyFifty{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (f *FiftyFifty) ShouldSucceed() bool {
	return f.rand.Float64() < 0.5
}

type Passer struct{}

func NewPasser() *Passer {
	return &Passer{}
}

func (p *Passer) ShouldSucceed() bool {
	return true
}

type Failer struct{}

func NewFailer() *Failer {
	return &Failer{}
}

func (f *Failer) ShouldSucceed() bool {
	return false
}

func RunFiftyFifty(cmd *cobra.Command, args []string) {
	if !NewFiftyFifty().ShouldSucceed() {
		os.Exit(1)
	}
}

func RunPasser(cmd *cobra.Command, args []string) {
	if !NewPasser().ShouldSucceed() {
		os.Exit(1)
	}
}

func RunFailer(cmd *cobra.Command, args []string) {
	if !NewFailer().ShouldSucceed() {
		os.Exit(1)
	}
}
