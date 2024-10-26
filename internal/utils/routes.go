package utils

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(rr ...Route) Routes {
	routes := make(Routes, 0)

	for _, r := range rr {
		routes = append(routes, r)
	}

	return routes
}

func (rr Routes) Setup() {
	for _, r := range rr {
		r.Setup()
	}
}
