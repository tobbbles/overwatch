package path

import (
	"net/http"

	"service/server/context/id/ability"
	"service/server/context/id/hero"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Middleware for populating contexts from path variables. This function does a fair bit of magic and ought to be
// updated to include ANY new path variables introduced to endpoints. It's common patten that these are then parsed and
// put into a context from their own context package.
func Middleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if len(mux.Vars(r)["hero_id"]) != 0 {
				ctx, err := hero.NewContext(r.Context(), mux.Vars(r)["hero_id"])
				if err != nil {
					logger.Error(
						"failed setting hero id from request path to context",
						zap.String("hero_id", mux.Vars(r)["hero_id"]),
						zap.Error(err),
					)
					return
				}

				r = r.WithContext(ctx)
			}

			if len(mux.Vars(r)["ability_id"]) != 0 {
				ctx, err := ability.NewContext(r.Context(), mux.Vars(r)["ability_id"])
				if err != nil {

					logger.Error(
						"failed setting ability id from request path to context",
						zap.String("ability_id", mux.Vars(r)["ability_id"]),
						zap.Error(err),
					)
				}

				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
