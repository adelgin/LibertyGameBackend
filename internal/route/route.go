package route

import (
	"libertyGame/internal/handler"

	"github.com/go-chi/chi"

	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	log *zerolog.Logger
	m   *chi.Mux
	i   *handler.Implementation
}

func NewRoutes(l *zerolog.Logger, m *chi.Mux, i *handler.Implementation) *Routes {
	return &Routes{
		log: l,
		m:   m,
		i:   i,
	}
}

func (route *Routes) Setup() {

	route.m.Group(func(r chi.Router) {
		{
			route.log.Info().Msg("setting up documentation routes")
			r.Mount("/swagger", httpSwagger.WrapHandler)
		}

		{
			route.log.Info().Msg("setting up handlers routes")
			r.Get("/users/all", route.i.CountOfAllUsers())                 // tested
			r.Get("/users/{count}/top", route.i.GetTopOfRefs())            // tested
			r.Get("/user/{id}", route.i.GetUserByID())                     // tested
			r.Get("/user/{id}/refs", route.i.GetRefsOfUserFromID())        // tested
			r.Get("/user/{id}/refscount", route.i.CountRefsOfUserFromID()) // tested
			r.Get("/monthstat", route.i.GetMonthStatisticsHandler())       // tested
			r.Post("/users", route.i.UserRegistration())                   // not tested
		}
	})
}
