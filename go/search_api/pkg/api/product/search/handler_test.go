package search_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/product"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/product/search"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type searchTestCase struct {
	query string // query string that represent pkg/api/products/search/request.go
	want  string
}

func searchTestCases() []*searchTestCase {
	return []*searchTestCase{
		{
			query: "q=Jeans", // test query filter
			want: `{"from":0,"query":{"bool":{"must":{"match":{"title":{"query":"Jeans"}}}}},"size":10,` +
				`"sort":[{"price":{"order":"asc"}}]}`,
		}, {
			query: "brand=Nike", // test brand filter
			want: `{"from":0,"query":{"bool":{"filter":{"term":{"brand":"Nike"}}}},"size":10,` +
				`"sort":[{"price":{"order":"asc"}}]}`,
		}, {
			query: "offset=0&limit=1&sort=asc&order_by=price", // test sorting
			want:  `{"from":0,"query":{"bool":{}},"size":1,"sort":[{"price":{"order":"asc"}}]}`,
		}, {
			query: "offset=0&limit=1&sort=desc&order_by=price", // test sorting
			want:  `{"from":0,"query":{"bool":{}},"size":1,"sort":[{"price":{"order":"desc"}}]}`,
		}, {
			query: "offset=1&limit=1&sort=asc&order_by=price", // test offset
			want:  `{"from":1,"query":{"bool":{}},"size":1,"sort":[{"price":{"order":"asc"}}]}`,
		}, {
			query: "offset=0&limit=2&sort=asc&order_by=price", // test limit
			want:  `{"from":0,"query":{"bool":{}},"size":2,"sort":[{"price":{"order":"asc"}}]}`,
		}, {
			query: "offset=-1&limit=0&sort=invalid&order_by=invalid", // test bad request
			want:  ``,
		},
	}
}

func TestHandler(t *testing.T) {
	t.Parallel()

	for _, tc := range searchTestCases() {
		tc := tc

		t.Run(fmt.Sprintf("query:%s", tc.query), func(t *testing.T) {
			t.Parallel()

			var esQuery string

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				body, _ := ioutil.ReadAll(req.Body)
				_ = req.Body.Close()
				esQuery = string(body)
			}))

			defer server.Close()

			handler, err := createHandler(server)
			if err != nil {
				t.Errorf("Can't create search Handler: %v", err)

				return
			}

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/?%s", tc.query), strings.NewReader(""))
			handler.Handle(c)

			if esQuery != tc.want {
				t.Errorf("Got %q elasticsearch query, want %q", esQuery, tc.want)
			}
		})
	}
}

func createHandler(server *httptest.Server) (*search.Handler, error) {
	esClient, err := elastic.NewClient(
		elastic.SetURL(server.URL),
		elastic.SetHttpClient(server.Client()),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, fmt.Errorf("can't create new elasticsearch client: %w", err)
	}

	repository := product.NewRepository(esClient, 300, "test")

	return search.NewHandler(repository), nil
}

func TestHandlerBadRequest(t *testing.T) {
	t.Parallel()

	handler := search.NewHandler(nil)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/?offset=-1&limit=0&sort=invalid&order_by=invalid",
		strings.NewReader(""),
	)

	handler.Handle(c)

	wantCode := http.StatusBadRequest

	if response.Code != wantCode {
		t.Errorf("Got %q response code, want %q", response.Code, wantCode)
	}

	wantContains := []string{
		"Field validation for 'OrderBy' failed",
		"Field validation for 'Sort' failed",
		"Field validation for 'Offset'",
		"Field validation for 'Limit' failed",
	}

	got := response.Body.String()

	for _, text := range wantContains {
		if !strings.Contains(got, text) {
			t.Errorf("Response %q does not contains %q", got, text)
		}
	}
}
