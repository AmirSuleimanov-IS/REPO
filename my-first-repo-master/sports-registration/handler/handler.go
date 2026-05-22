package handler

import (
"html/template"
"net/http"
"path/filepath"
"sports-registration/internal/database"
"sports-registration/models"
"strconv"
"strings"
)

type Handler struct {
tmpl *template.Template
}

func NewHandler() *Handler {
h := &Handler{}
h.loadTemplates()
return h
}

func (h *Handler) loadTemplates() {
funcs := template.FuncMap{
"eq": func(a, b string) bool { return a == b },
}

// Получаем абсолютные пути к шаблонам
baseDir, err := filepath.Abs("templates")
if err != nil {
panic(err)
}

// Сначала загружаем layout шаблоны
layoutFiles := []string{
filepath.Join(baseDir, "layout", "header.html"),
filepath.Join(baseDir, "layout", "footer.html"),
}

// Затем основные страницы
pageFiles := []string{
filepath.Join(baseDir, "home.html"),
filepath.Join(baseDir, "athletes.html"),
filepath.Join(baseDir, "events.html"),
filepath.Join(baseDir, "athlete-detail.html"),
filepath.Join(baseDir, "event-detail.html"),
filepath.Join(baseDir, "team-application.html"),
}

allFiles := append(layoutFiles, pageFiles...)
h.tmpl = template.Must(template.New("").Funcs(funcs).ParseFiles(allFiles...))
}

// 1. GET /services → поиск активных услуг (ORM)
func (h *Handler) ServicesHandler(w http.ResponseWriter, r *http.Request) {
search := r.URL.Query().Get("search")
var services []models.Service
q := database.DB.Where("status = ?", "active")
if search != "" {
q = q.Where("name ILIKE ? OR category ILIKE ?", "%"+search+"%", "%"+search+"%")
}
q.Find(&services)

var draftCount int64
// Для демо фиксируем user_id = 1
database.DB.Model(&models.Application{}).Where("user_id = 1 AND status = ?", "draft").Count(&draftCount)

data := map[string]interface{}{
"Services":    services,
"Search":      search,
"DraftCount":  draftCount,
"PageTitle":   "Услуги",
"CartCount":   draftCount,
}

h.tmpl.ExecuteTemplate(w, "events", data)
}

// 2. GET /service?id=X → детальная информация (ORM)
func (h *Handler) ServiceDetailHandler(w http.ResponseWriter, r *http.Request) {
idStr := r.URL.Query().Get("id")
id, err := strconv.Atoi(idStr)
if err != nil {
http.Error(w, "Неверный ID услуги", http.StatusBadRequest)
return
}

var svc models.Service
result := database.DB.Where("id = ? AND status = ?", id, "active").First(&svc)
if result.Error != nil {
http.NotFound(w, r)
return
}

data := map[string]interface{}{
"Service":   svc,
"PageTitle": svc.Name,
}

h.tmpl.ExecuteTemplate(w, "event-detail", data)
}

// 3. GET /application/draft → просмотр текущей заявки (ORM)
func (h *Handler) ViewDraftHandler(w http.ResponseWriter, r *http.Request) {
var app models.Application
err := database.DB.Where("user_id = 1 AND status = ?", "draft").Preload("Services.Service").First(&app).Error
if err != nil {
http.Error(w, "Черновик не найден", http.StatusNotFound)
return
}

data := map[string]interface{}{
"Application": app,
"PageTitle":   "Заявка #" + strconv.Itoa(int(app.ID)),
"CartCount":   1,
}

h.tmpl.ExecuteTemplate(w, "team-application", data)
}

// 4. POST /application/add → добавление услуги в заявку (ORM)
func (h *Handler) AddToApplicationHandler(w http.ResponseWriter, r *http.Request) {
serviceIDStr := r.FormValue("service_id")
serviceID, err := strconv.Atoi(serviceIDStr)
if err != nil {
http.Error(w, "Неверный ID услуги", http.StatusBadRequest)
return
}

// Находим или создаем черновик заявки
var app models.Application
database.DB.Where("user_id = 1 AND status = ?", "draft").FirstOrCreate(&app, models.Application{
UserID:   1,
Status:   "draft",
TeamName: "Kinetic Draft Team",
})

// Находим услугу
var svc models.Service
database.DB.First(&svc, serviceID)

// Проверяем, есть ли уже эта услуга в заявке
var appSvc models.ApplicationService
found := database.DB.Where("application_id = ? AND service_id = ?", app.ID, svc.ID).First(&appSvc)

if found.Error != nil {
// Услуги нет, создаем новую запись
appSvc = models.ApplicationService{
ApplicationID: app.ID,
ServiceID:     svc.ID,
Quantity:      1,
FinalPrice:    svc.BasePrice,
RoleInEvent:   "participant",
}
database.DB.Create(&appSvc)
} else {
// Услуга есть, увеличиваем количество
appSvc.Quantity++
appSvc.FinalPrice = svc.BasePrice * float64(appSvc.Quantity) // 📐 ФОРМУЛА
database.DB.Save(&appSvc)
}

// Пересчитываем общую сумму заявки
var total float64
database.DB.Model(&models.ApplicationService{}).Where("application_id = ?", app.ID).
Select("COALESCE(SUM(final_price), 0)").Scan(&total)
database.DB.Model(&app).Update("total_amount", total)

http.Redirect(w, r, "/services", http.StatusSeeOther)
}

// 5. POST /application/delete → логическое удаление (RAW SQL, без ORM)
func (h *Handler) DeleteDraftHandler(w http.ResponseWriter, r *http.Request) {
idStr := r.FormValue("id")
appID, err := strconv.Atoi(idStr)
if err != nil {
http.Error(w, "Неверный ID заявки", http.StatusBadRequest)
return
}

// Логическое удаление через SQL UPDATE
res := database.DB.Exec("UPDATE applications SET status = 'deleted', completed_at = NOW() WHERE id = ? AND status = 'draft'", appID)
if res.RowsAffected == 0 {
http.Error(w, "Заявка не найдена", http.StatusBadRequest)
return
}

http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func (h *Handler) AthletesHandler(w http.ResponseWriter, r *http.Request) {
h.ServicesHandler(w, r)
}

func (h *Handler) EventsHandler(w http.ResponseWriter, r *http.Request) {
h.ServicesHandler(w, r)
}

func (h *Handler) AthleteDetailHandler(w http.ResponseWriter, r *http.Request) {
h.ServiceDetailHandler(w, r)
}

func (h *Handler) EventDetailHandler(w http.ResponseWriter, r *http.Request) {
h.ServiceDetailHandler(w, r)
}

func (h *Handler) TeamApplicationHandler(w http.ResponseWriter, r *http.Request) {
h.ViewDraftHandler(w, r)
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
var draftCount int64
database.DB.Model(&models.Application{}).Where("user_id = 1 AND status = ?", "draft").Count(&draftCount)

data := map[string]interface{}{
"CartCount": draftCount,
"PageTitle": "Главная",
}

h.tmpl.ExecuteTemplate(w, "home", data)
}
