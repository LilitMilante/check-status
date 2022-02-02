package service

import (
	"check-status/entity"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type StatusService struct {
	c *http.Client
}

func NewStatusService() *StatusService {
	return &StatusService{
		c: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (ss *StatusService) CheckStatus(addresses []string) string {
	var (
		wg       sync.WaitGroup
		countAdr = len(addresses)
		chSender = make(chan string)
		chStatus = make(chan entity.Status, countAdr)
	)

	defer close(chSender)

	go ss.readFromChan(chStatus, chSender)

	wg.Add(countAdr)

	for _, adr := range addresses {
		go ss.getServerStatus(&wg, chStatus, adr)

	}

	wg.Wait()

	close(chStatus)

	return <-chSender
}

func (ss *StatusService) getServerStatus(wg *sync.WaitGroup, chStatus chan entity.Status, adr string) {
	defer wg.Done()

	var s entity.Status
	s.URL = adr

	resp, err := ss.c.Get(adr)
	if err != nil {
		s.StatusCode = 0

		chStatus <- s

		return
	}

	defer resp.Body.Close()

	s.StatusCode = resp.StatusCode

	chStatus <- s
}

func (ss *StatusService) readFromChan(chStatus chan entity.Status, chSender chan string) {
	for s := range chStatus {
		if s.StatusCode != http.StatusOK {
			chSender <- fmt.Sprintf("Сервис %s не работает", s.URL)

			return
		}
	}

	chSender <- "Всё работает"
}
