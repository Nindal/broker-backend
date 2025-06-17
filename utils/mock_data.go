package utils

func MockHoldings() map[string]interface{} {
	return map[string]interface{}{
		"holdings": []map[string]interface{}{
			{"stock": "AAPL", "qty": 10, "avg_price": 150},
			{"stock": "TSLA", "qty": 5, "avg_price": 700},
		},
	}
}

func MockOrderbook() map[string]interface{} {
	return map[string]interface{}{
		"orders": []map[string]interface{}{
			{"stock": "AAPL", "price": 160, "qty": 10},
		},
		"realized_pnl":   250,
		"unrealized_pnl": -120,
	}
}

func MockPositions() map[string]interface{} {
	return map[string]interface{}{
		"positions": []map[string]interface{}{
			{"stock": "GOOG", "qty": 8, "pnl": 100},
		},
		"total_pnl": 100,
	}
}
