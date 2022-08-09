package routers

import (
	"github.com/ribeirosaimon/go_flight_api/src/controllers"
	"net/http"
)

var FlighRouter = []Router{
	{
		URI:                "/flight",
		Method:             http.MethodGet,
		Function:           controllers.CreateFlight,
		NeedAuthentication: false,
	},
}
