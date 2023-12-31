package main

import "github.com/joho/godotenv"

func GetEnv() map[string]string {
	godotenv.Load(".env")
	env, _ := godotenv.Read()

	return env
}
