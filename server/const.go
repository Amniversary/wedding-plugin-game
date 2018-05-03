package server

import (
	"net/http"
)

const (
	Empty      = 0
)

type MethodFunc func(w http.ResponseWriter, r *http.Request)
