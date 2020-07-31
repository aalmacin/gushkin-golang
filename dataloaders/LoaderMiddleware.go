package dataloaders

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

const activityLoaderKey = "activityLoader"

func LoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activityLoader := &ActivityLoader{
			fetch: func(ids []string) ([]*model.Activity, []error) {
				var activities []*model.Activity

				err := db.Model(&activities).Where("id in ?", ids).Select()

				if err != nil {
					fmt.Println("Error in ActivityLoaderMiddleware", err)
				}

				return activities, nil
			},
			maxBatch: 100,
			wait:     1 * time.Millisecond,
		}

		ctx := context.WithValue(r.Context(), activityLoaderKey, activityLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
