package messages

import "net/http"

func GetConversation(w http.ResponseWriter, r *http.Request, userId string) {
	// TODO: retourner les messages échangés avec l’utilisateur
}

func SendMessage(w http.ResponseWriter, r *http.Request, userId string) {
	// TODO: envoyer un message privé
}
