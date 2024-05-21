package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Я не тупой, просто сделал неправильный пуш и потом проглядел и сделал реквест на практикум, после чего мой комп, с которого писал все сдох и я пытался как то сделать это все с телефона, простите за эту дичь(
// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Ниже напишите обработчики для каждого эндпоинта
// ...
func getTasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func postTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks[task.ID] = task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func getIdTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задачи нет", http.StatusNoContent)
		return
	}
	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func deliteIdTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задачи нет", http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	r := chi.NewRouter()

	// здесь регистрируйте ваши обработчики
	// ...
	r.Get("/tasks", getTasks)
	r.Post("/tasks", postTask)
	r.Get("/tasks/{id}", getIdTask)
	r.Delete("/tasks/{id}", deliteIdTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
