package handler

import "net/http"

//like decorator in python
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			if len(username) < 3 || isTokenValid(token) {
				//redirect to the login page
				http.Redirect(w, r, "/static/signin.html", http.StatusFound)
			}
			h(w, r)
		})
}
