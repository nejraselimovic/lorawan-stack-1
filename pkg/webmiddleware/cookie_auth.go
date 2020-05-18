// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webmiddleware

import (
	"net/http"

	"github.com/gorilla/csrf"
	"go.thethings.network/lorawan-stack/v3/pkg/auth"
)

// CookieAuth extracts the auth cookie and forwards it to the Authorization
// header as SessionToken while securing the authorization via CSRF token.
func CookieAuth(cookieKey string, csrfAuthKey []byte) MiddlewareFunc {
	CSRFMiddleware := CSRF(csrfAuthKey, csrf.CookieName("_csrf"))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") == "" {
				cookieValue, err := r.Cookie(cookieKey)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				encoder, _ := GetEncoder(r.Context())
				authCookie := &auth.CookieShape{}
				err = encoder.Decode(cookieKey, cookieValue.Value, authCookie)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				if authCookie.TokenKey != "" {
					key := auth.JoinToken(auth.SessionToken, authCookie.TokenID, authCookie.TokenKey)
					r.Header.Add("Authorization", "Bearer "+key)
					if w.Header().Get("Access-Control-Allow-Origin") != "" {
						w.Header().Del("Access-Control-Allow-Origin")
					}
					CSRFMiddleware(next).ServeHTTP(w, r)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
