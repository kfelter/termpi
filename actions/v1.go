package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/kfelter/termpi/models"
)

type PingRequestV1 struct {
	ThingID string   `json:"user_id"`
	Secret  string   `json:"secret"`
	Status  string   `json:"status"`
	Tags    []string `json:"tags"`
}

// V1Ping default implementation.
func V1Ping(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	pr := PingRequestV1{}
	err = json.Unmarshal(b, &pr)
	if err != nil {
		return err
	}

	// Allocate an empty Thing
	thing := &models.Thing{}
	if err := tx.Find(thing, pr.ThingID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if thing.Secret != pr.Secret {
		return c.Error(http.StatusForbidden, fmt.Errorf("secrets did not match"))
	}

	thing.Tags = joinTags(thing.Tags, pr.Tags)
	thing.Status = pr.Status

	_, err = tx.ValidateAndUpdate(thing)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(map[string]string{"status": "accepted"}))
}

func joinTags(existing, new []string) []string {
	r := []string{}
	eMap := map[string]string{}
	for _, t := range existing {
		if ss := strings.SplitN(t, ":", 2); len(ss) == 2 {
			eMap[ss[0]] = ss[1]
		} else {
			eMap[ss[0]] = ""
		}
	}
	for _, t := range new {
		if ss := strings.SplitN(t, ":", 2); len(ss) == 2 {
			eMap[ss[0]] = ss[1]
		} else {
			eMap[ss[0]] = ""
		}
	}
	for k, v := range eMap {
		if v != "" {
			r = append(r, k+":"+v)
		} else {
			r = append(r, k)
		}
	}
	sort.Strings(r)
	return r
}

// V1Client default implementation.
func V1Client(c buffalo.Context) error {
	f, err := os.Open("/bin/things")
	if err != nil {
		return c.Error(404, err)
	}
	return c.Render(http.StatusOK, r.Download(c, "things", f))
}
