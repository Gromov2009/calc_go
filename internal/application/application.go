package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	calculation "github.com/Gromov2009/calc_go/pkg"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {

	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//fmt.Println("error:", err.Error())
		http.Error(w, `{ "error": "Expression is not valid" }`, http.StatusUnprocessableEntity)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		//fmt.Fprintf(w, "err: %s", err.Error())
		if errors.Is(err, calculation.ErrInvalidExpression) || errors.Is(err, calculation.ErrMissingClosingBracket) || errors.Is(err, calculation.ErrDivByZero) {
			http.Error(w, `{ "error": "Expression is not valid" }`, http.StatusUnprocessableEntity)
		} else {
			http.Error(w, `{ "error": "unknown error" }`, http.StatusNotImplemented)
		}
	} else {
		fmt.Fprintf(w, `{ "result": "%f" }`, result)
	}

}
