package main

import (
	"log"

	"github.com/YWJSonic/GameServer/alien/game"

	"server/env"

	"github.com/joho/godotenv"
)

type ENV struct {
	MAINTAIN bool `json:"Maintain"`
	MAINTAIN_START_TIME string `json:"MaintainStartTime"`
	MAINTAIN_FINISH_TIME string `json:"MaintainFinishTime"`
	MAINTAIN_CHECKOUT_TIME string `json:"ULGMaintainCheckoutTime"`

	SERVER_IP string `json:"IP"`
	SERVER_HTTP_PORT string `json:"PORT"`
	SERVER_SOCKET_PORT string `json:"SocketPORT"`

	DB_IP string `json:"DBIP"`
	DB_PORT string `json:"DBPORT"`
	DB_USER string `json:"DBUser"`
	DB_PASSWORD string `json:"DBPassword"`

	REDIS_URL string `json:"RedisURL"`

	ACCOUNT_ENCODE: string `json:"AccountEncodeStr"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panicf("No [ .env ] file found...\n")
	}

	env := Env{
		MAINTAIN: env.GetEnvAsBool("MAINTAIN"),
		MAINTAIN_START_TIME: env.GetEnvAsString("MAINTAIN_START_TIME"),
		MAINTAIN_FINISH_TIME: env.GetEnvAsString("MAINTAIN_FINISH_TIME"),
		MAINTAIN_CHECKOUT_TIME: env.GetEnvAsString("MAINTAIN_CHECKOUT_TIME"),

		SERVER_IP: env.GetEnvAsString("SERVER_IP"),
		SERVER_HTTP_PORT: env.GetEnvAsString("SERVER_HTTP_PORT"),
		SERVER_SOCKET_PORT: env.GetEnvAsString("SERVER_SOCKET_PORT"),

		DB_IP: env.GetEnvAsString("DB_IP"),
		DB_PORT: env.GetEnvAsString("DB_PORT"),
		DB_USER: env.GetEnvAsString("DB_USER"),
		DB_PASSWORD: env.GetEnvAsString("DB_PASSWORD"),

		REDIS_URL: env.GetEnvAsString("REDIS_URL"),

		ACCOUNT_ENCODE: env.GetEnvAsString("ACCOUNT_ENCODE"),
	}

	jsonbyte, err := json.Marshal(env)
	if err != nil {
		log.Panicf("error:", err)
	}

	game.NewGameServer(string(jsonbyte))
}
