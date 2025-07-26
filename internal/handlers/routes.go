package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func SetupRoutes(r *chi.Mux, db *gorm.DB) {
	classHandler := &ClassHandler{DB: db}
	inquiryHandler := &InquiryHandler{DB: db}
	// subjectHandler := &SubjectHandler{DB: db}
	// teacherHandler := &TeacherHandler{DB: db}
	// studentHandler := &StudentHandler{DB: db}
	// announcementHandler := &AnnouncementHandler{DB: db}
	// homeworkHandler := &HomeworkHandler{DB: db}

	r.Route("/api", func(r chi.Router) {
		r.Get("/test", testAPIHandler)

		r.Route("/class", func(r chi.Router) {
			r.Post("/", classHandler.Create)
			r.Get("/{id}", classHandler.Get)
			r.Put("/{id}", classHandler.Update)
			r.Delete("/{id}", classHandler.Delete)
			r.Get("/", classHandler.List)
		})

		r.Route("/inquiry", func(r chi.Router) {
			r.Get("/all", inquiryHandler.GetInquiries)
		})

	})
}

// r.Route("/subjects", func(r *chi.Mux) {
// 	r.Post("/", subjectHandler.Create)
// 	r.Get("/{id}", subjectHandler.Get)
// 	r.Put("/{id}", subjectHandler.Update)
// 	r.Delete("/{id}", subjectHandler.Delete)
// 	r.Get("/", subjectHandler.List)
// })
// r.Route("/teachers", func(r *chi.Mux) {
// 	r.Post("/", teacherHandler.Create)
// 	r.Get("/{id}", teacherHandler.Get)
// 	r.Put("/{id}", teacherHandler.Update)
// 	r.Delete("/{id}", teacherHandler.Delete)
// 	r.Get("/", teacherHandler.List)
// })
// r.Route("/students", func(r *chi.Mux) {
// 	r.Post("/", studentHandler.Create)
// 	r.Get("/{id}", studentHandler.Get)
// 	r.Put("/{id}", studentHandler.Update)
// 	r.Delete("/{id}", studentHandler.Delete)
// 	r.Get("/", studentHandler.List)
// })
// r.Route("/announcements", func(r *chi.Mux) {
// 	r.Post("/", announcementHandler.Create)
// 	r.Get("/{id}", announcementHandler.Get)
// 	r.Put("/{id}", announcementHandler.Update)
// 	r.Delete("/{id}", announcementHandler.Delete)
// 	r.Get("/", announcementHandler.List)
// })
// r.Route("/homework", func(r *chi.Mux) {
// 	r.Post("/", homeworkHandler.Create)
// 	r.Get("/{id}", homeworkHandler.Get)
// 	r.Put("/{id}", homeworkHandler.Update)
// 	r.Delete("/{id}", homeworkHandler.Delete)
// 	r.Get("/", homeworkHandler.List)
// })

func testAPIHandler(w http.ResponseWriter, r *http.Request) {
	response := "API IS SUCCESSFULLY RUN"
	json.NewEncoder(w).Encode(response)
}
