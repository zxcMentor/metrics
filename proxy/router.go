package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"metricsProm/proxy/internal/controller"
	"metricsProm/proxy/internal/metrics"
	"metricsProm/proxy/internal/service"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var token *jwtauth.JWTAuth

func newApiRouter() *chi.Mux {
	r := chi.NewRouter()
	token = jwtauth.New("HS256", []byte("secret"), nil)

	r.Mount("/debug", middleware.Profiler())

	r.Handle("/metrics", promhttp.Handler())

	r.Group(func(r chi.Router) {
		fmt.Println("protected route is entered")

		//if token == nil {
		//	fmt.Println("token is empty")
		//	return
		//}

		r.Use(jwtauth.Verifier(token))
		fmt.Println("token verified")
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search*", func(w http.ResponseWriter, r *http.Request) {

			startTime := time.Now()
			m := metrics.NewCounterHandler()
			m.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path}).Inc()

			//r.ParseForm()
			//
			//addrReq := r.FormValue("query") //Header.Get("query")

			request, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("request reading error:", err)
				return
			}

			var addrReq service.Addresser

			addrReq = &service.AddressDadata{
				Query: "",
				Cache: repo,
			}

			err = json.Unmarshal(request, addrReq)

			if err != nil {
				fmt.Println("request parsing error:", err)
				return

			}

			//fmt.Println("query parsed:", addrReq)

			//addrReq = `{"query":"нижний советс 3"}`

			resp := addrReq.GetAddress()

			newRespond := &controller.Response{}

			logger := zap.NewExample()

			if err == nil {
				newRespond.Data = resp
				newRespond.Success = true

				responder := controller.NewResponder(logger, newRespond)
				responder.OutputJSON(w, err)
			}

			if strings.Contains(err.Error(), "forb") {
				newRespond.Success = false

				responder := controller.NewResponder(logger, newRespond)
				responder.ErrorForbidden(w, err)
			}

			if strings.Contains(err.Error(), "400") {
				responder := controller.NewResponder(logger, newRespond)
				responder.ErrorBadRequest(w, err)
			}

			elapsedTime := time.Since(startTime)
			metrics.NewTimeHandler().
				With(prometheus.Labels{"method": r.Method, "path": r.URL.Path}).
				Observe(elapsedTime.Seconds())

		})

		r.Post("/api/address/geocode*", func(w http.ResponseWriter, r *http.Request) {

			request, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("request reading error:", err)
				return
			}

			var geoSuggReq service.Geocoder

			geoSuggReq = &service.GeocodeDadata{}

			err = json.Unmarshal(request, geoSuggReq)

			if err != nil {

				fmt.Println("request parsing error:", err)
				return

			}

			//geoSuggReq := `{ "lat": "55.878", "lng": "37.653" }`

			resp, err := geoSuggReq.GetGeocode()

			newRespond := &controller.Response{}

			logger := zap.NewExample()

			if err == nil {
				newRespond.Data = resp
				newRespond.Success = true

				responder := controller.NewResponder(logger, newRespond)
				responder.OutputJSON(w, err)
			}

			if strings.Contains(err.Error(), "forb") {
				newRespond.Success = false

				responder := controller.NewResponder(logger, newRespond)
				responder.ErrorForbidden(w, err)
			}

			if strings.Contains(err.Error(), "400") {
				responder := controller.NewResponder(logger, newRespond)
				responder.ErrorBadRequest(w, err)
			}

		})

	})

	r.Get("/swagger*", func(w http.ResponseWriter, r *http.Request) {

		//http.ServeFile(w, r, "./swagger.json")

		//fmt.Println("/swagger entered")

		w.Header().Set("Content-Type", "text/html	; charset=utf-8")
		tmpl, err := template.New("swagger").Parse(swaggerTemplate)

		//fmt.Println("template prepared")
		if err != nil {
			fmt.Println("template error:", err)
			return
		}
		err = tmpl.Execute(w, struct {
			Time int64
		}{
			Time: time.Now().Unix(),
		})
		if err != nil {
			fmt.Println("template execution error:", err)
			return
		}
	})

	r.Get("/docs/*", func(w http.ResponseWriter, r *http.Request) {
		//r.Get("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "/docs/swagger.json")

		//fmt.Println("/docs entered")

		//http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))).ServeHTTP(w, r)

	})

	r.Get("/api/register/{user}/{password}", func(w http.ResponseWriter, r *http.Request) {

		user := chi.URLParam(r, "user")
		password := chi.URLParam(r, "user")

		regRespond := Register(user, password)

		w.Write([]byte(regRespond))

	})

	r.Get("/api/login/{user}/{password}", func(w http.ResponseWriter, r *http.Request) {

		user := chi.URLParam(r, "user")
		password := chi.URLParam(r, "user")

		respond := Login(user, password)

		if respond == "" {
			w.Write([]byte("user is not authorised, please login if registerred"))
		} else {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(map[string]string{"token": respond})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("user login succeed, auth token assigned:\n" + respond + "\n please use it manually in \"search.md\" file either in Swagger/Postman/other tools"))
		}

	})

	r.Get("/api/login/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("please login using path \"api/login/username/password\" or register first if not registered yet using the path \"/api/register/username/password\""))
	})

	return r

}
