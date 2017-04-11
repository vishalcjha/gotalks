package adaptersetup

import (
	"errors"
	"math/rand"
	"net/http"
	"time"
)

//HelloWithError this generate error with alternate request
func HelloWithError(w http.ResponseWriter, req *http.Request) error {
	if rand.Int63n(int64(time.Now().Second()))%2 == 0 {
		return errors.New("Error In Handler")
	}
	return nil
}
