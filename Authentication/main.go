package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var hmacSecret = []byte("replace-with-strong-secret")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

// createToken builds a simple HS256 JWT (no external libs).
func createToken(username string, ttl time.Duration) (string, error) {
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	hb, _ := json.Marshal(header)
	headerPart := base64.RawURLEncoding.EncodeToString(hb)

	now := time.Now().Unix()
	claims := Claims{
		Sub: username,
		Iat: now,
		Exp: now + int64(ttl.Seconds()),
	}
	cb, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payloadPart := base64.RawURLEncoding.EncodeToString(cb)

	signingInput := headerPart + "." + payloadPart
	mac := hmac.New(sha256.New, hmacSecret)
	mac.Write([]byte(signingInput))
	signature := mac.Sum(nil)
	sigPart := base64.RawURLEncoding.EncodeToString(signature)

	return signingInput + "." + sigPart, nil
}

// validateToken verifies signature and expiration, returns claims on success.
func validateToken(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("token must have 3 parts")
	}
	signingInput := parts[0] + "." + parts[1]
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, errors.New("invalid signature encoding")
	}

	mac := hmac.New(sha256.New, hmacSecret)
	mac.Write([]byte(signingInput))
	expected := mac.Sum(nil)
	if !hmac.Equal(sig, expected) {
		return nil, errors.New("invalid token signature")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, errors.New("invalid payload encoding")
	}

	var claims Claims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, errors.New("invalid claims")
	}

	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("token expired")
	}
	return &claims, nil
}

// loginHandler issues a token for valid credentials (demo: username=user password=pass)
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// Demo credential check. Replace with real authentication.
	if creds.Username != "user" || creds.Password != "pass" {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := createToken(creds.Username, 15*time.Minute)
	if err != nil {
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// authMiddleware verifies Bearer token and sets the username in the request context via header.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := validateToken(token)
		if err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}
		// For simplicity, pass username in a header for the downstream handler.
		r.Header.Set("X-User", claims.Sub)
		next.ServeHTTP(w, r)
	})
}

// protectedHandler demonstrates an authenticated endpoint.
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "hello, " + user})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))

	addr := ":8080"
	fmt.Println("server listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}
