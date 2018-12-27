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
	for i, stg := range p.Stages {
		start := time.Now()
		stageNum := i + 1
		p.Print(fmt.Sprintf("%d. [RUN] %s", stageNum, stg.Name()))

		if err := stg.Run(); err != nil {
			p.Print(fmt.Sprintf("%d. [ERROR] [TOOK: %s] %s", stageNum, time.Since(start), err))
			return err
		}

		p.Print(fmt.Sprintf("%d. [COMPLETE] [TOOK: %s]\n", stageNum, time.Since(start)))
	}
	return nil
}

func (p *Pipeline) AddStage(name string, f func() error) {
	p.Stages = append(p.Stages, stage{
		name:  name,
		runFn: f,
	})
}

type Stage interface {
	Run() error
	Name() string
}

type stage struct {
	name  string
	runFn func() error
}

func (s stage) Run() error {
	return s.runFn()
}

func (s stage) Name() string {
	return s.name
}
