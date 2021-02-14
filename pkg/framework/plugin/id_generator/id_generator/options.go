package id_generator

import (
	"framework/class/id_generator"
)

type Options struct{}

type Option func(o *Options)

func New(opts ...Option) id_generator.IdGenerator {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	p := &plugin{
		opts: options,
	}
	if err := p.Init(); err != nil {
		panic(err)
	}

	return p
}
