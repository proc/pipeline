package pipeline

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Pipeline struct {
	Stages []Stage
	log    logrus.FieldLogger
}

type Stage interface {
	Run() error
	Name() string
}

type stage struct {
	name  string
	runFn func() error
}

func New(log logrus.FieldLogger) *Pipeline {
	return &Pipeline{
		log: log,
	}

}

func (p *Pipeline) Print(s string) {
	if p.log != nil {
		p.log.Info(s)
	} else {
		fmt.Println(s)
	}
}

func (p *Pipeline) Run() error {
	p.Print(fmt.Sprintf("[begin] %d stage(s)", len(p.Stages)))
	runStart := time.Now()

	for i, stg := range p.Stages {
		stageNum := i + 1
		p.Print(fmt.Sprintf("%d) [%s]", stageNum, stg.Name()))

		start := time.Now()
		if err := stg.Run(); err != nil {
			p.Print(fmt.Sprintf("%d. [error] [%s] %s", stageNum, time.Since(start), err))
			return err
		}

		p.Print(fmt.Sprintf("%d) [%s]", stageNum, time.Since(start)))
	}

	p.Print(fmt.Sprintf("[success] %s", time.Since(runStart)))
	return nil
}

func (p *Pipeline) AddStage(name string, f func() error) {
	p.Stages = append(p.Stages, stage{
		name:  name,
		runFn: f,
	})
}

func (s stage) Run() error {
	return s.runFn()
}

func (s stage) Name() string {
	return s.name
}
