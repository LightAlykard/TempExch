package rest

import "net/http"

func (s *Service) logout(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("logout method called")
	clearCookies(w)
}
