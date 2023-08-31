package handlers

import (
	"encoding/json"
	"fmt"
	"time"
)

type nbpEurResponse struct {
	Table    string    `json:"table"`
	Currency string    `json:"currency"`
	Code     string    `json:"code"`
	Rates    []nbpRate `json:"rates"`
}

type nbpRate struct {
	No            string  `json:"no"`
	EffectiveDate string  `json:"effectiveDate"`
	Ratio         float64 `json:"mid"`
}

func (h *Handler) CurrencyCheck(x, y int) error {
	for i := 0; i < x; i++ {
		hResp, err := h.get(nbpEurLast100URL)
		if err != nil {
			return err
		}

		var nbpResp nbpEurResponse

		err = json.Unmarshal(hResp.Body, &nbpResp)
		if err != nil {
			return err
		}

		for _, rate := range nbpResp.Rates {
			if rate.Ratio > 4.7 || rate.Ratio < 4.5 {
				b := []byte(fmt.Sprintf(`{"EffectiveDate": %s, "Ratio": %f}`, rate.EffectiveDate, rate.Ratio))
				b = append(b, "\n"...)

				_, err = h.W.Write(b)
				if err != nil {
					return err
				}
			}
		}

		time.Sleep(time.Duration(y) * time.Second)

	}
	return nil
}
