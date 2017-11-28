package hello

import (
	"context"
	"net/http"
	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type (
	testService struct {
		name string
	}
)

func NewTestService() marvin.JSONService {
	return testService{
		name: "xiaoyu",
	}
}


// no need for CORS
func (s testService) HTTPMiddleware(h http.Handler) http.Handler {
	return h
}

// auth managed by Google and we don't need identity
func (s testService) Middleware(e endpoint.Endpoint) endpoint.Endpoint {
	return e
}

// no custom error handler
func (s testService) Options() []httptransport.ServerOption {
	return []httptransport.ServerOption{
        httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
            // check proto/json by inspecting url
            //path := ctx.Value(httptransport.ContextKeyRequestPath).(string)
            //if strings.HasSuffix(path, ".json") {
            httptransport.EncodeJSONResponse(ctx, w, err)
            return
            //}
            //marvin.EncodeProtoResponse(ctx, w, err)
        }),
    }
}

// gorilla is fine, we need it for recalc endpoint
func (s testService) RouterOptions() []marvin.RouterOption {
	return []marvin.RouterOption{
        marvin.RouterSelect("stdlib"),  // *** Updated this for faster non-static routing
    }
}

func (s testService) JSONEndpoints() map[string]map[string]marvin.HTTPEndpoint {
	return map[string]map[string]marvin.HTTPEndpoint{
		"health": {
			"GET": {
				Endpoint: s.doJob,
				Decoder:  s.decodeGetRequest,
			},
		},
	}
}

type Name struct {
	first string
	last string
}

func (s testService) doJob(ctx context.Context, req interface{}) (interface{}, error) {
	myName := req.(*Name)
	return &myName, nil
}

func (s testService) decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	myName := Name {"xiaoyu", "xu"}
	return &myName, nil
}