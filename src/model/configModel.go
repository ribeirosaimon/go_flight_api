package model

import "time"

type ConfigApiModel struct {
	Read        bool      `json:"read"`
	CurrentTime time.Time `json:"current_time"`
}
