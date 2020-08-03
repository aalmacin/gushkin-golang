package graph

import "github.com/aalmacin/gushkin-golang/repos"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	ActivityRepo *repos.ActivityRepo
	ActionRepo   *repos.ActionRepo
	WishRepo     *repos.WishRepo
	UserID       string
}
