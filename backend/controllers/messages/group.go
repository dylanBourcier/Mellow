package messages

import "net/http"

func GetGroupMessages(w http.ResponseWriter, r *http.Request, groupId string) {
	// TODO: retourner les messages du groupe
}

func SendGroupMessage(w http.ResponseWriter, r *http.Request, groupId string) {
	// TODO: envoyer un message au groupe
}
