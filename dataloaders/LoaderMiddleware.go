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
const actionLoaderKey = "actionLoader"

func LoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activityLoader := &ActivityLoader{
			fetch: func(ids []string) ([]*model.Activity, []error) {
				var activities []*model.Activity

				err := db.Model(&activities).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					fmt.Println("Error in ActivityLoaderMiddleware", err)
				}

				// Create buffer
				activitiesBuff := make(map[string]*model.Activity)

				// Populate buffer
				for _, activity := range activities {
					activitiesBuff[activity.ID] = activity
				}

				// Make sorted result slice
				results := make([]*model.Activity, len(ids))

				// Map buffer to result
				for i, id := range ids {
					results[i] = activitiesBuff[id]
				}

				return results, nil
			},
			maxBatch: 100,
			wait:     1 * time.Millisecond,
		}

		actionLoader := &ActionLoader{
			fetch: func(ids []string) ([]*model.Action, []error) {
				var activities []*model.Action

				err := db.Model(&activities).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					fmt.Println("Error in ActionLoaderMiddleware", err)
				}

				// Create buffer
				activitiesBuff := make(map[string]*model.Action)

				// Populate buffer
				for _, action := range activities {
					activitiesBuff[action.ID] = action
				}

				// Make sorted result slice
				results := make([]*model.Action, len(ids))

				// Map buffer to result
				for i, id := range ids {
					results[i] = activitiesBuff[id]
				}

				return results, nil
			},
			maxBatch: 100,
			wait:     1 * time.Millisecond,
		}

		ctx := context.WithValue(r.Context(), activityLoaderKey, activityLoader)
		ctx = context.WithValue(ctx, actionLoaderKey, actionLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetActivityLoader(ctx context.Context) *ActivityLoader {
	return ctx.Value(activityLoaderKey).(*ActivityLoader)
}

func GetActionLoader(ctx context.Context) *ActionLoader {
	return ctx.Value(actionLoaderKey).(*ActionLoader)
}
