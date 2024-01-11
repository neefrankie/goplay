package main

import (
	"context"
	"errors"
	"fmt"
	"goplay/web/model"
	"goplay/web/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/casbin/casbin/v2"
)

var sessionManager *scs.SessionManager

func main() {
	anthEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true

	users := model.CreateUsers()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler(users))
	mux.HandleFunc("/logout", logoutHandler())
	mux.HandleFunc("/member/current", currentMemberHandler())
	mux.HandleFunc("/member/role", memberRoleHandler())
	mux.HandleFunc("/admin/stuff", adminHandler())

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", sessionManager.LoadAndSave(Authorizer(anthEnforcer, users)(mux))))
}

func loginHandler(users model.Users) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		user, err := users.FindByName(name)
		if err != nil {
			render.New(w).BadRequest("WRONG_CREDENTIALS")
			return
		}
		if err := sessionManager.RenewToken(r.Context()); err != nil {
			render.New(w).InternalServerError(err.Error())
			return
		}

		sessionManager.Put(r.Context(), "userId", user.ID)
		sessionManager.Put(r.Context(), "role", user.Role)
		render.New(w).NoContent()
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := sessionManager.RenewToken(r.Context()); err != nil {
			render.New(w).InternalServerError(err.Error())
			return
		}

		render.New(w).NoContent()
	})
}

func getRole(ctx context.Context) (string, error) {
	role := sessionManager.GetString(ctx, "role")
	if role == "" {
		return "", errors.New("NO_ROLE")
	}

	return role, nil
}

func getUserID(ctx context.Context) (int, error) {
	uid := sessionManager.GetInt(ctx, "userId")
	if uid == 0 {
		return 0, errors.New("NO_USER_ID")
	}

	return uid, nil
}

func currentMemberHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid, err := getUserID(r.Context())
		if err != nil {
			render.NewInternalError(err.Error())
			return
		}

		render.New(w).Text(http.StatusOK, fmt.Sprintf("User with ID: %d", uid))
	})
}

func memberRoleHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, err := getRole(r.Context())
		if err != nil {
			render.New(w).InternalServerError(err.Error())
			return
		}

		render.New(w).Text(http.StatusOK, fmt.Sprintf("User with Role: %s", role))
	})
}

func adminHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.New(w).Text(http.StatusOK, "I'm an admin")
	})
}

func Authorizer(e *casbin.Enforcer, users model.Users) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			role := sessionManager.GetString(r.Context(), "role")

			if role == "" {
				role = "anonymous"
			}

			if role == "member" {
				uid, err := getUserID(r.Context())
				if err != nil {
					render.New(w).InternalServerError(err.Error())
					return
				}

				exists := users.Exists(uid)
				if !exists {
					render.New(w).Forbidden("user does not exist")
					return
				}
			}

			res, err := e.Enforce(role, r.URL.Path, r.Method)
			if err != nil {
				render.New(w).InternalServerError(err.Error())
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				render.ErrorForbidden("unauthorized")
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}
