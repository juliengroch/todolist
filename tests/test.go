package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/juliengroch/todolist/application"
	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/server"
	"github.com/juliengroch/todolist/store"
	"github.com/juliengroch/todolist/tests/fixtures"
	"github.com/juliengroch/todolist/views"
)

// ContentTypeJSON set json
const ContentTypeJSON = "application/json"

// Request make a test request
type Request struct {
	URI     string
	Auth    *Auth
	Payload map[string]interface{}
}

// Auth credentials
type Auth struct {
	Name string
	Key  string
}

// Runner is test runner with context
func Runner(t *testing.T, runTest func(router *gin.Engine, ctx context.Context)) {
	gin.SetMode(gin.TestMode)

	// read config test file
	cfg, err := config.LoadConfigFile(os.Getenv("TODOLIST_CONF_TEST"))
	assert.Nil(t, err)

	// load Context
	ctx, err := application.Load(cfg)
	assert.Nil(t, err)

	st := store.FromContext(ctx)

	// init Bdd
	require.NoError(t, st.Migrate(ctx))

	// init data
	err = fixtures.InitTestData(ctx)
	if err != nil {
		cleanBddIfErr(ctx)
	}

	// start gin router
	router := server.Router(ctx)
	views.Routes(router)

	// execute Test
	runTest(router, ctx)

	// end test clean bdd
	require.NoError(t, st.ResetDB(ctx))
	require.NoError(t, st.Close(ctx))
}

// GET make HTTP GET request with authentificated user on the uri
func GET(router *gin.Engine, req *Request) *httptest.ResponseRecorder {
	return request(router, "GET", req)
}

// POST make HTTP POST request with authentificated user on the uri
func POST(router *gin.Engine, req *Request) *httptest.ResponseRecorder {
	return request(router, "POST", req)
}

// PATCH make HTTP PATCH request with authentificated user on the uri
func PATCH(router *gin.Engine, req *Request) *httptest.ResponseRecorder {
	return request(router, "PATCH", req)
}

func request(router *gin.Engine, method string, req *Request) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	var hreq *http.Request

	if req.Payload != nil {
		body, _ := json.Marshal(req.Payload)
		hreq, _ = http.NewRequest(method, req.URI, bytes.NewBuffer(body))
	} else {
		hreq, _ = http.NewRequest(method, req.URI, nil)
	}

	hreq.Header.Set("Content-Type", ContentTypeJSON)

	auth := fmt.Sprintf("ApiKey %s:%s", req.Auth.Name, req.Auth.Key)
	hreq.Header.Set("Authorization", auth)

	router.ServeHTTP(resp, hreq)

	return resp
}

func cleanBddIfErr(ctx context.Context) {
	st := store.FromContext(ctx)

	err := st.ResetDB(ctx)
	if err != nil {
		panic(err)
	}

	err = st.Migrate(ctx)
	if err != nil {
		panic(err)
	}

	err = fixtures.InitTestData(ctx)
	if err != nil {
		panic(err)
	}
}
