package hello

import (
	"net/http"
	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
)

func NewStatsService() marvin.JSONService {
	return nil
}

