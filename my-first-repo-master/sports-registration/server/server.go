package server

import (
	"log"
	"net/http"
	"sports-registration/handler"
)

func StartServer() {
	h := handler.NewHandler()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Основные маршруты
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/athletes", h.AthletesHandler)
	http.HandleFunc("/athlete", h.AthleteDetailHandler)
	http.HandleFunc("/events", h.EventsHandler)
	http.HandleFunc("/event", h.EventDetailHandler)
	http.HandleFunc("/team-application", h.TeamApplicationHandler)

	// Маршруты для работы с услугами и заявками (Лабораторная 12)
	http.HandleFunc("/services", h.ServicesHandler)
	http.HandleFunc("/service", h.ServiceDetailHandler)
	http.HandleFunc("/application/draft", h.ViewDraftHandler)
	http.HandleFunc("/application/add", h.AddToApplicationHandler)
	http.HandleFunc("/application/delete", h.DeleteDraftHandler)

	log.Println("🏃 Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
