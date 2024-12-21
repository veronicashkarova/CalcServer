package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/veronicashkarova/CalcServer/pkg/calculation"
	"net/http"
	"os"
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

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		switch {
		case errors.Is(err, calculation.ErrInvalidExpression):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, calculation.ErrEmptyExpression):
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "result: %f", result)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	fmt.Println("Server started")
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
