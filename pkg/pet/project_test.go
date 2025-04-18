package pet_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viqueen/go-template/pkg/pet/petmocks"
	"testing"
)

func TestProject(t *testing.T) {
	project := petmocks.NewProject(t)
	project.EXPECT().Get().Return("project")
	assert.Equal(t, "project", project.Get())
}
