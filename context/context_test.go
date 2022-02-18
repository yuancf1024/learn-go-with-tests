package context

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder() // 返回一个初始化了的ResponseRecorder

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf('got "%s", want "%s"', response.Body.String(), data)
	}
}