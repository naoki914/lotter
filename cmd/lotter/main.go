package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/naoki914/lotter/internal/adapters"
	"github.com/naoki914/lotter/internal/domain/dhlotto"
	"github.com/naoki914/lotter/internal/infra"
)

type Server struct {
	api     adapters.Api
	service adapters.Service
}

func NewServer(api adapters.Api, service adapters.Service) *Server {
	return &Server{
		api:     api,
		service: service,
	}
}

func main() {
	var lotteryType string
	var operation string
	var id int
	var num int

	flag.StringVar(&lotteryType, "lottery", "dhlotto", "The type of the lottery.")
	flag.StringVar(&operation, "operation", "output", "The operation to perform. (create, update, output)")
	flag.IntVar(&id, "id", -1, "The ID of the specific draw.")
	flag.IntVar(&num, "num", 1, "The number of draws to generate.")

	flag.Parse()

	db := infra.NewDbImpl("mongodb://localhost:27027", "lottery", "draws")
	// defer db.Client.Disconnect(context.Background())

	var server Server
	switch lotteryType {
	case "dhlotto":
		api := infra.NewDHLottoApiImpl("https://www.dhlottery.co.kr/common.do?method=getLottoNumber&drwNo=")
		service := dhlotto.DHService{}

		server = *NewServer(api, &service)

	default:
		fmt.Printf("Unsupported lottery type: %s\n", lotteryType)
		os.Exit(101)
	}

	switch operation {

	case "update-all":
		updateAll(server.api, db)
	case "update":
		if id < 0 {
			fmt.Println("id flag should be set with a value of 1 or above.")
			return
		}
		handleDraw(server.api, db, 1)
	case "output":
		draws, err := db.GetAll(lotteryType)
		if err != nil {
			fmt.Printf("Could not get all draws from db: %v\n", err)
			os.Exit(102)
		}
		if num < 1 {
			num = 1
		}
		for i := 0; i < num; i++ {
			solPri, solSec, err := server.service.LottoSolution(draws)
			if err != nil {
				fmt.Printf("couldn't get solution: %+v", err)
				continue
			}
			sort.Ints(solPri)
			fmt.Printf("%d: %+v,%+v\n", i+1, solPri, solSec)
		}

	default:
		fmt.Printf("Unsupported operation: %s\n", operation)
		os.Exit(1002)
	}

}

func updateAll(api adapters.Api, db adapters.DB) {
	done := make([]string, 0)
	for i := 1; ; i++ {
		draw, err := api.FetchDrawWithID(i)
		if err != nil {
			fmt.Println("Error fetching draw:", err)
			return
		}
		if !draw.IsSuccess() {
			break
		}
		err = db.Create(draw)
		if err != nil {
			fmt.Println("Error saving draw:", err)
		}
		done = append(done, draw.GetID())
	}
	fmt.Println("Successfully fetched and saved draw", done)
}

func handleDraw(api adapters.Api, db adapters.DB, id int) {
	draw, err := api.FetchDrawWithID(id)
	if err != nil {
		fmt.Println("Error fetching draw:", err)
		return
	}
	err = db.Create(draw)
	if err != nil {
		fmt.Println("Error saving draw:", err)
		return
	}
	fmt.Println("Successfully fetched and saved draw", draw.GetID())
}
