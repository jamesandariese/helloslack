package helloslack

import (
	"appengine"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if textToPrint, err := json.Marshal("Hello " + r.Form.Get("user_name")); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{'text': 'unexpected error marshalling stuff'}`)
		return
	} else {
		c := appengine.NewContext(r)
		var jsonLogBuf bytes.Buffer
		logWriter := io.MultiWriter(&jsonLogBuf, w)
		fmt.Fprintf(logWriter, `{"username":"Hello Bot","text": %s}`, textToPrint)
		c.Infof("%v", jsonLogBuf.String())
	}
}
