package system

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetConfig(key string) string {
	for _, k := range SELECT_Settings_All() {
		if k.Key == key {
			return k.Value
		}
	}

	Warning("Couldn't retrieve key: "+key, true)
	return ""
}

func GetInt(key string) int {
	for _, k := range SELECT_Settings_All() {
		if k.Key == key {
			intKey, err := strconv.Atoi(k.Value)
			Error("Couldn't convert key to int", err)

			return intKey
		}
	}

	Warning("Couldn't retrieve key: "+key, true)
	return -1
}

func EncodeID(id string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(id)
	Error("Couldn't convert string ID to mongo ID Object", err)

	return objID
}

func InitSettings() {
	if len(SELECT_Settings_All()) > 0 {
		return
	}

	settings := []Setting{
		{
			Key:   "version",
			Value: "1.0",
		},
		{
			Key:   "patch",
			Value: "1",
		},
		{
			Key:   "runtime-log",
			Value: "res/logs/runtime.txt",
		},
		{
			Key:   "error-log",
			Value: "res/logs/error.txt",
		},
		{
			Key:   "print-log",
			Value: "res/logs/print.txt",
		},
		{
			Key:   "json-log",
			Value: "res/logs/jsonlog.txt",
		},
		{
			Key:   "server-port",
			Value: "8008",
		},
		{
			Key:   "mongo-enabled",
			Value: "0",
		},
		{
			Key:   "mongo-password-req",
			Value: "0",
		},
		{
			Key:   "mongo-user",
			Value: "root",
		},
		{
			Key:   "mongo-password",
			Value: "rootpassword",
		},
		{
			Key:   "mongo-ip",
			Value: "127.0.0.1",
		},
		{
			Key:   "mongo-port",
			Value: "27017",
		},
	}

	for _, setting := range settings {
		INSERT_Settings(setting.Key, setting.Value)
	}
}
