package pipeline

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	p := Pipeline{}
	logger, hook := test.NewNullLogger()
	p.log = logger

	s := "base string"
	x := "another string"
	z := "string 3"

	p.Stages = append(p.Stages, stage{
		name: "Do something",
		runFn: func() error {
			s = "stage 1"
			return nil
		},
	})

	p.Stages = append(p.Stages, stage{
		name: "and then",
		runFn: func() error {
			s = "the next stage"
			x = "tee"
			return nil
		},
	})

	p.Stages = append(p.Stages, stage{
		name: "finally",
		runFn: func() error {
			s = "the next stage"
			z = "zee"
			return nil
		},
	})

	err := p.Run()

	assert.NoError(t, err)
	assert.Equal(t, "the next stage", s)
	assert.Equal(t, "tee", x)
	assert.Equal(t, "zee", z)
	assert.Equal(t, 6, len(hook.Entries))
}
