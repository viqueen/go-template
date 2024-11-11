package pet_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viqueen/go-template/pkg/pet"
	"testing"
)

func TestProject(t *testing.T) {
	want := "pet project"
	got := pet.Project()
	assert.Equal(t, want, got)
}
