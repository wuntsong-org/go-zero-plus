package httpx

import (
	"fmt"
	"github.com/wuntsong-org/go-zero-plus/rest/internal/header"
	"net/http"
	"strings"
)

const xForwardedFor = "X-Forwarded-For"

// GetFormValues returns the form values.
func GetFormValues(r *http.Request) (map[string]any, error) {
	if strings.Contains(r.Header.Get(header.ContentType), header.FormDataType) {
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			return nil, err
		}

		if r.MultipartForm == nil {
			return nil, fmt.Errorf("not form-data")
		}

		params := make(map[string]any, len(r.MultipartForm.Value))
		for name, value := range r.MultipartForm.Value {
			if len(value) != 1 {
				continue
			}

			formValue := value[0]
			if len(formValue) == 0 {
				continue
			}

			params[name] = formValue
		}

		return params, nil
	} else if strings.Contains(r.Header.Get(header.ContentType), header.FormUrlEncodedType) {
		err := r.ParseForm()
		if err != nil {
			return nil, err
		}

		params := make(map[string]any, len(r.Form))
		for name := range r.Form {
			formValue := r.Form.Get(name)
			if len(formValue) == 0 {
				continue
			}
			params[name] = formValue
		}

		return params, nil
	} else {
		return make(map[string]any, 0), nil
	}
}

// GetRemoteAddr returns the peer address, supports X-Forward-For.
func GetRemoteAddr(r *http.Request) string {
	v := r.Header.Get(xForwardedFor)
	if len(v) > 0 {
		return v
	}

	return r.RemoteAddr
}
