
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// a concurrent safe map type by embedding sync.Mutex
type cancelMap struct {
	sync.Mutex
	internal map[string]context.CancelFunc
}

func newCancelMap() *cancelMap {
	return &cancelMap{
		internal: make(map[string]context.CancelFunc),
	}
}

func (c *cancelMap) Get(key string) (value context.CancelFunc, ok bool)  {
	c.Lock()
	result, ok := c.internal[key]
	c.Unlock()
	return result, ok
}

func (c *cancelMap) Set(key string, value context.CancelFunc) {
	c.Lock()
	c.internal[key] = value
	c.Unlock()
}

func (c *cancelMap) Delete(key string) {
	c.Lock()
	delete(c.internal, key)
	c.Unlock()
}

// create global jobs map with cancel function
var jobs = newCancelMap()

// the pretend worker will be wrapped here
func work(ctx context.Context, id string) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Cancelling job id %s\n", id)
			return
		case <-time.After(time.Second):
			fmt.Printf("Doing job id %s\n", id)
		}
	}
}


func startHandler(w http.ResponseWriter, r *http.Request) {

	// get job id and name from query parameters
	id := r.URL.Query().Get("id")

	// check if job already exists in jobs map
	if _, ok := jobs.Get(id); ok {
		fmt.Fprintf(w, "Already started job id: %s\n", id)
		return
	}

	// create new context with cancel for the job
	ctx, cancel := context.WithCancel(context.Background())

	// save it in the global map of jobs
	jobs.Set(id, cancel)

	// actually start running the job
	go work(ctx, id)

	// return 200 with message
	fmt.Fprintf(w, "Job id: %s has been started\n", id)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {

	// get job id and name from query parameters
	id := r.URL.Query().Get("id")

	// check for cancel func from jobs map
	cancel, found := jobs.Get(id)
	if !found {
		fmt.Fprintf(w, "Job id: %s is not running\n", id)
		return
	}

	// cancel the jobs
	cancel()

	// delete job from jobs map
	jobs.Delete(id)

	// return 200 with message
	fmt.Fprintf(w, "Job id: %s has been canceled\n", id)
}

func main() {
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/stop", stopHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
