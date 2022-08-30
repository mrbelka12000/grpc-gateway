package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrbelka12000/grpc-gateway/internal/models"
	"github.com/mrbelka12000/grpc-gateway/proto"
	"net/http"
	"time"
)

const apiURL = "https://www.rusprofile.ru/ajax.php?query=%v&action=search"

type implementer struct{}

func New() *implementer {
	return &implementer{}
}

func (i *implementer) GetInfoByIIN(ctx context.Context, message *proto.Message) (*proto.Response, error) {

	send := time.Now()
	resp, err := http.Get(fmt.Sprintf(apiURL, message.IIN))
	if err != nil {
		return nil, err
	}
	fmt.Printf("send request time: %v \n", time.Since(send).Seconds())

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not ok resp status")
	}

	defer resp.Body.Close()

	respData := &models.Data{}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	if len(respData.Ul) < 1 {
		return nil, errors.New("customer not found by iin")
	}

	return &proto.Response{
		CompanyName: respData.Ul[0].Name,
		KPP:         "not implemented",
		Fio:         respData.Ul[0].CeoName,
		IIn:         respData.Ul[0].Inn,
	}, nil
}
