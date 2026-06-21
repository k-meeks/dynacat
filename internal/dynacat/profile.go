package dynacat

import (
	"net/http"
	"time"
)

const PROFILE_COOKIE_NAME = "active_profile"

func (a *application) ProfilesEnabled() bool {
	return len(a.Config.Auth.Profiles) > 0
}

func (a *application) isValidProfile(name string) bool {
	if name == "" {
		return false
	}
	_, exists := a.profileSet[name]
	return exists
}

func (a *application) handleSetProfileRequest(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if !a.isValidProfile(name) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     PROFILE_COOKIE_NAME,
		Value:    name,
		Path:     a.Config.Server.BaseURL + "/",
		Secure:   a.isRequestHTTPS(r),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Expires:  time.Now().Add(2 * 365 * 24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
}
