package settings

import (
	"testing"
)

func TestSettings_Flimit(t *testing.T) {
	s := Settings{}
	s.Config.Flimit = 1
	if err := s.Flimit(); err != nil {
		t.Fatal("Unable to increase limit", err)
	}
}