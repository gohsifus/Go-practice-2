package httpServer

import (
	"net/http"
	"net/url"
	"strconv"
	"task11/errs"
)

func (s Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func (s Server) createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		params, err := parseParams(r)
		if err != nil {
			cErr := errs.Wrap(err)
			s.Response(w, cErr.Status(), cErr.Error(), false)
			s.log.Error(cErr.Error())
		}

		event, err := s.service.CreateEvent(
			params.Get("date"),
			params.Get("name"),
			params.Get("description"),
		)

		if err != nil {
			cErr := errs.Wrap(err)
			s.Response(w, cErr.Status(), cErr.Error(), false)
			s.log.Error(cErr.Error())
		} else {
			s.Response(w, 200, event, true)
		}
	}
}

func (s Server) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		params, err := parseParams(r)
		if err != nil {
			cErr := errs.Wrap(err)
			s.Response(w, cErr.Status(), cErr.Error(), false)
			s.log.Error(cErr.Error())
		}

		err = s.service.UpdateEvent(
			params.Get("id"),
			params.Get("date"),
			params.Get("name"),
			params.Get("description"),
		)

		if err != nil {
			cErr := errs.Wrap(err)
			s.Response(w, cErr.Status(), cErr.Error(), false)
			s.log.Error(cErr.Error())
		} else {
			s.Response(w, 200, "update success", true)
		}
	}
}

func (s Server) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		id, err := strconv.Atoi(r.PostForm.Get("id"))

		if err != nil {
			cErr := errs.New(err, errs.IncorrectDataErr)
			s.Response(w, cErr.Status(), cErr.Error(), false)
			s.log.Error(cErr.Error())
		} else {
			err = s.service.DeleteEvent(id)
			if err != nil {
				cErr := errs.Wrap(err)
				s.Response(w, cErr.Status(), cErr.Error(), false)
				s.log.Error(cErr.Error())
			} else {
				s.Response(w, 200, "delete success", true)
			}
		}
	}
}

func (s Server) getEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()

	events, err := s.service.GetEventsForDay(args.Get("from"), args.Get("to"))
	if err != nil {
		cErr := errs.Wrap(err)
		s.Response(w, cErr.Status(), cErr.Error(), false)
		s.log.Error(cErr.Error())
	} else {
		s.Response(w, 200, events, true)
	}
}

func (s Server) getEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()

	events, err := s.service.GetEventsForWeek(args.Get("from"), args.Get("to"))
	if err != nil {
		cErr := errs.Wrap(err)
		s.Response(w, cErr.Status(), cErr.Error(), false)
		s.log.Error(cErr.Error())
	} else {
		s.Response(w, 200, events, true)
	}
}

func (s Server) getEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()

	events, err := s.service.GetEventsForMonth(args.Get("date"))
	if err != nil {
		cErr := errs.Wrap(err)
		s.Response(w, cErr.Status(), cErr.Error(), false)
		s.log.Error(cErr.Error())
	} else {
		s.Response(w, 200, events, true)
	}
}

// Для парсинга параметров метода на update и create
func parseParams(r *http.Request) (url.Values, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, errs.New(err, errs.IncorrectDataErr)
	}

	return r.PostForm, nil
}
