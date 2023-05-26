package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/naoki914/lotter/internal/domain"
	"github.com/naoki914/lotter/internal/domain/dhlotto"
)

type DHLottoApiImpl struct {
	baseUrl string
}

func NewDHLottoApiImpl(baseUrl string) *DHLottoApiImpl {
	return &DHLottoApiImpl{
		baseUrl: baseUrl,
	}
}

func (api *DHLottoApiImpl) FetchDrawWithID(id int) (domain.Draw, error) {
	url := api.baseUrl + fmt.Sprint(id)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("couln't read response body")
	}
	dhdraw := dhlotto.DHDraw{}
	json.Unmarshal(body, &dhdraw)

	if dhdraw.Success != "fail" {
		return &dhdraw, nil
	} else {
		return nil, fmt.Errorf("draw number (ID) does not exist")
	}

}

func (api *DHLottoApiImpl) FetchAll(id int) (domain.Draw, error) {
	url := api.baseUrl + fmt.Sprint(id)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("couln't read response body")
	}

	dhdraw := dhlotto.DHDraw{}
	json.Unmarshal(body, &dhdraw)

	if dhdraw.Success != "fail" {
		return &dhdraw, nil
	} else {
		return nil, fmt.Errorf("draw number (ID) does not exist")
	}
}
