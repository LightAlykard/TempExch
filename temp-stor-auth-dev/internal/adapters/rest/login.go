package rest

import (
	"encoding/json"
	"net/http"
)

type ctxKeyUser struct{}
type credentials struct {
	Login string `json:"username"`
	Pass  string `json:"password"`
}

// CheckAuth reads login/pass from requests body and returns JWT tokenpair
func (s *Service) login(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("Authentication with json in body call")

	// get credentials from request's body
	decoder := json.NewDecoder(r.Body)
	var creds credentials
	err := decoder.Decode(&creds)
	if err != nil {
		s.logger.Sugar().Debugf("Error reading creds from request's body, %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// obtain token if password is correct
	token, err := s.auth.Login(r.Context(), creds.Login, creds.Pass)
	if err != nil {
		s.logger.Sugar().Debugf("wrong login/pass provided, %s: %v", creds.Login, err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	clearCookies(w)
	sendCookies(token.Access, w)

	s.logger.Sugar().Debugf("user %s authenticated successfuly", creds.Login)

}
