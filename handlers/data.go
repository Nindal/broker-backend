package handlers

import (
	"encoding/json"
	"net/http"
	"broker-backend/utils"
)

func GetHoldings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.MockHoldings())
}

func GetOrderbook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.MockOrderbook())
}

func GetPositions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.MockPositions())
}
