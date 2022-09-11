package httpServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task11/infrastructure/inMemory"
	"task11/interfaces/httpServer/configs"
	"task11/logger"
	"task11/middleware"
	"task11/service"
)

type Server struct {
	service service.Calendar
	mux     *http.ServeMux
	config  *configs.ServerConfig
	log     *logger.Log
}

func NewServer(config *configs.ServerConfig) (*Server, error) {
	calendar := service.NewCalendar(inMemory.NewInMemoryRepo())
	mux := http.NewServeMux()
	logger, err := logger.NewLogger(config.PathToLog)
	if err != nil {
		return nil, err
	}

	return &Server{
		service: calendar,
		mux:     mux,
		config:  config,
		log: logger,
	}, nil
}

func (s Server) Start() {
	fmt.Println("start server")
	s.log.Info("start server")
	s.ConfigureServer()

	addr := s.config.Host + ":" + s.config.Port

	http.ListenAndServe(addr, s.mux)
}

func (s Server) ConfigureServer() {
	s.mux.Handle("/", middleware.Logging(http.HandlerFunc(s.helloHandler), s.log))
	s.mux.Handle("/create_event", middleware.Logging(http.HandlerFunc(s.createEventHandler), s.log))
	s.mux.Handle("/update_event", middleware.Logging(http.HandlerFunc(s.updateEventHandler), s.log))
	s.mux.Handle("/delete_event", middleware.Logging(http.HandlerFunc(s.deleteEventHandler), s.log))
	s.mux.Handle("/events_for_day", middleware.Logging(http.HandlerFunc(s.getEventsForDayHandler), s.log))
	s.mux.Handle("/events_for_week", middleware.Logging(http.HandlerFunc(s.getEventsForWeekHandler), s.log))
	s.mux.Handle("/events_for_month", middleware.Logging(http.HandlerFunc(s.getEventsForMonthHandler), s.log))
}

//Response функция для формирования ответа на запрос
func (s Server) Response(w http.ResponseWriter, status int, data interface{}, success bool) {
	w.WriteHeader(status)

	respData := make(map[string]interface{})

	if success {
		respData["result"] = data
	} else {
		respData["error"] = data
	}

	bytes, err := json.Marshal(respData)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error: " + err.Error()))
		return
	}

	w.Write(bytes)
}
