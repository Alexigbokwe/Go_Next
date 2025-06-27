package app

func RegisterModuleComponents(container *Container, components ...any) {
	for _, c := range components {
		container.Register(c)
		container.MustAutowire(c)
	}
}
