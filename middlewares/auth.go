package middlewares

import (
	"context"
	"github.com/Matemateg/blog/entities"
	"github.com/Matemateg/blog/service"
	"net/http"
)

func Auth(next http.Handler, userService *service.UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("SSID")
		if err == nil && cookie != nil {
			currentUser, err := userService.GetBySSID(cookie.Value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ctx := context.WithValue(r.Context(), "currentUser", currentUser)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func GetCurrentUser(ctx context.Context) *entities.User {
	currentUser, _ := ctx.Value("currentUser").(*entities.User)
	return currentUser
}
