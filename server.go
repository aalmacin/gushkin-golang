package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aalmacin/gushkin-golang/auth"
	"github.com/aalmacin/gushkin-golang/dataloaders"
	"github.com/aalmacin/gushkin-golang/graph"
	"github.com/aalmacin/gushkin-golang/graph/generated"
	hooks "github.com/aalmacin/gushkin-golang/pg-hooks"
	"github.com/aalmacin/gushkin-golang/repos"
	"github.com/go-pg/pg"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func getDBOptions() *pg.Options {
	opts, err := pg.ParseURL(os.Getenv("POSTGRES_URL"))

	if err != nil {
		fmt.Errorf("Error getting options from url. %v", err)
		return nil
	}

	return opts
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := pg.Connect(getDBOptions())
	defer db.Close()
	db.AddQueryHook(hooks.DBLogger{})

	activityRepo := repos.ActivityRepo{
		DB: db,
	}

	actionRepo := repos.ActionRepo{
		DB: db,
	}

	wishRepo := repos.WishRepo{
		DB: db,
	}

	resolver := &graph.Resolver{
		ActivityRepo: &activityRepo,
		ActionRepo:   &actionRepo,
		WishRepo:     &wishRepo,
		UserID:       "",
	}

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srvWithLoader := dataloaders.LoaderMiddleware(db, srv)
	srvWithAuth := auth.JwtMiddleware().Handler(srvWithLoader)
	srvWithCurrentUser := auth.CurrentUserMiddleware(resolver, srvWithAuth)
	srvWithCors := c.Handler(srvWithCurrentUser)

	http.Handle("/", c.Handler(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", srvWithCors)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
