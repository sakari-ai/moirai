package bootstrap

import (
	"github.com/sakari-ai/moirai/config/loader"
	"github.com/facebookgo/inject"
)

func populate(opts ...Object) error {
	//Default loading database and config.loader
	opts = append(opts,
		ByName("config_loader", cfgLoader),
	)
	if dbEngine != nil {
		opts = append(opts, ByName("database", dbEngine))
	}
	i := &ioc{
		ObjectFuncs: opts,
	}
	return i.Populate()
}

func Populate(f func(l loader.Loader) []Object) error {
	opts := f(cfgLoader)
	return populate(opts...)
}

type Object func() *inject.Object

func ByValue(value interface{}) func() *inject.Object {
	return func() *inject.Object {
		return &inject.Object{
			Value: value,
		}
	}
}

func ByName(name string, value interface{}) func() *inject.Object {
	return func() *inject.Object {
		return &inject.Object{
			Name:  name,
			Value: value,
		}
	}
}

type ioc struct {
	ObjectFuncs []Object
}

func (i *ioc) Populate() error {
	var g inject.Graph
	var objects []*inject.Object
	for _, f := range i.ObjectFuncs {
		objects = append(objects, f())
	}

	if err := g.Provide(objects...); err != nil {
		return err
	}
	return g.Populate()
}
