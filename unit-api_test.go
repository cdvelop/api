package api_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cdvelop/api"
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/fetchserver"
	"github.com/cdvelop/logserver"
	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

type auth struct{}

func (auth) GetLoginUser(params any) (*model.User, error) {

	user := model.User{
		Token:          "123",
		Id:             "123456789101112",
		Ip:             "172.0.0.1",
		Name:           "don Juanito dev test",
		Area:           "s",
		AccessLevel:    "1",
		LastConnection: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &user, nil
}

func (auth) UserAuthNumber() (string, error) {
	return "1", nil
}

func Test_Api(t *testing.T) {
	h := &model.Handlers{
		AuthAdapter: auth{},
		Logger:      logserver.Add(),
		FileApi:     module{},
	}

	objects := ModuleProduct().Objects
	h.AddObjects(objects...)

	cutkey.AddDataConverter(h)

	fetchserver.AddFetchAdapter(h)

	conf, err := api.Add(h)
	if err != nil {
		t.Fatal(err)
		return
	}

	mux := conf.ServeMuxAndRoutes()
	srv := httptest.NewServer(mux)
	defer srv.Close()

	for prueba, r := range testData {
		t.Run((prueba), func(t *testing.T) {

			r.Server = srv

			endpoint := srv.URL

			if r.Endpoint != "" {
				endpoint += "/" + r.Endpoint
			}
			if r.Object != "" {
				endpoint += "/" + r.Object
			}
			// fmt.Println("ENDPOINT:", endpoint)
			// var err error
			h.SendOneRequest(r.Method, endpoint, r.Object, r.Data, func(resp []map[string]string, err error) {

				if err != nil {
					t.Fatal(err)
					return
				}

				testools.CheckTest(prueba, r.Expected, resp, t)
			})

		})
	}
}
