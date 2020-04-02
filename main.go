package main

import (
	"encoding/json"
	"log"

	"server/env"
	"server/game"

	"github.com/joho/godotenv"
)

type ENV struct {
	Maintain             bool   `json:"Maintain"`
	MaintainStartTime    string `json:"MaintainStartTime"`
	MaintainFinishTime   string `json:"MaintainFinishTime"`
	MaintainCheckoutTime string `json:"ULGMaintainCheckoutTime"`

	ServerIP         string `json:"IP"`
	ServerHTTPPort   string `json:"PORT"`
	ServerSocketPort string `json:"SocketPORT"`

	DBIP       string `json:"DBIP"`
	DBPort     string `json:"DBPORT"`
	DBUser     string `json:"DBUser"`
	DBPassword string `json:"DBPassword"`

	RedisURL string `json:"RedisURL"`

	AccountEncode string `json:"AccountEncodeStr"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panicf("No [ .env ] file found...\n")
	}

	env := ENV{
		Maintain:             env.GetEnvAsBool("MAINTAIN"),
		MaintainStartTime:    env.GetEnvAsString("MAINTAIN_START_TIME"),
		MaintainFinishTime:   env.GetEnvAsString("MAINTAIN_FINISH_TIME"),
		MaintainCheckoutTime: env.GetEnvAsString("MAINTAIN_CHECKOUT_TIME"),

		ServerIP:         env.GetEnvAsString("SERVER_IP"),
		ServerHTTPPort:   env.GetEnvAsString("SERVER_HTTP_PORT"),
		ServerSocketPort: env.GetEnvAsString("SERVER_SOCKET_PORT"),

		DBIP:       env.GetEnvAsString("DB_IP"),
		DBPort:     env.GetEnvAsString("DB_PORT"),
		DBUser:     env.GetEnvAsString("DB_USER"),
		DBPassword: env.GetEnvAsString("DB_PASSWORD"),

		RedisURL: env.GetEnvAsString("REDIS_URL"),

		AccountEncode: env.GetEnvAsString("ACCOUNT_ENCODE"),
	}

	jsonbyte, err := json.Marshal(env)
	if err != nil {
		log.Panicf("error:", err)
	}

	game.NewGameServer(string(jsonbyte))
}
