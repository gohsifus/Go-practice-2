package inMemory

import (
	"errors"
	"sync"
	"task11/errs"
	"time"
)
import "task11/domain/entity"

type inMemoryEventRepo struct {
	store     map[int]entity.Event
	currentId int
	sync.RWMutex
}

func NewInMemoryRepo() *inMemoryEventRepo {
	return &inMemoryEventRepo{
		store:     make(map[int]entity.Event),
		currentId: 0,
	}
}

func (im *inMemoryEventRepo) Create(event *entity.Event) (*entity.Event, error) {
	im.Lock()
	event.Id = im.currentId
	im.store[event.Id] = *event
	im.currentId++
	im.Unlock()

	return event, nil
}

func (im *inMemoryEventRepo) Update(id int, event *entity.Event) error {
	im.Lock()
	old, ok := im.store[id]
	if ok {
		if !event.Date.IsZero() {
			old.Date = event.Date
		}

		if event.Name != "" {
			old.Name = event.Name
		}

		if event.Description != "" {
			old.Description = event.Description
		}
	} else {
		im.Unlock()
		return errs.New(errors.New("item not found"), errs.BusinessLogicErr)
	}
	im.store[id] = old
	im.Unlock()

	return nil
}

func (im *inMemoryEventRepo) Delete(id int) error {
	im.Lock()
	delete(im.store, id)
	im.Unlock()

	return nil
}

func (im inMemoryEventRepo) GetEventsByDateInterval(from, to time.Time) ([]entity.Event, error) {
	events := []entity.Event{}

	im.RLock()
	for _, v := range im.store {
		if (v.Date.After(from) || v.Date.Equal(from)) && (v.Date.Before(to) || v.Date.Equal(to)) {
			events = append(events, v)
		}
	}
	im.RUnlock()

	return events, nil
}
