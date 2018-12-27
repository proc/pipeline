package pipeline

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Pipeline struct {
	Stages []Stage
	log    *logrus.Logger
}

type Stage interface {
	Run() error
	Name() string
}

type stage struct {
	name  string
	runFn func() error
}

func New(log *logrus.Logger) *Pipeline {
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
	p.Print(fmt.Sprintf("[BEGIN PIPELINE] %d stage(s)", len(p.Stages)))

	for i, stg := range p.Stages {
		stageNum := i + 1
		p.Print(fmt.Sprintf("%d. [RUN STAGE] %s", stageNum, stg.Name()))

		start := time.Now()
		if err := stg.Run(); err != nil {
			p.Print(fmt.Sprintf("%d. [ERROR] [TOOK: %s] %s", stageNum, time.Since(start), err))
			return err
		}

		p.Print(fmt.Sprintf("%d. [STAGE COMPLETE] [TOOK: %s]\n", stageNum, time.Since(start)))
	}

	p.Print(fmt.Sprintf("[END PIPELINE] %d stage(s)", len(p.Stages)))
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
