package main

import (
	"HackNu/helpers"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func StudentPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		return
	}

	student := &Student{}
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		log.Println(err)
		WriteJSON(w, http.StatusOK, "\"error\": \"You should provide name and goal of the student\"")
		return
	}

	js, err := os.ReadFile("./prompts/template.txt")
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := helpers.SendGPT(string(js), student.Goal, "gpt-4")
	if err != nil {
		log.Println(err)
		return
	}
	kkk := &Data{}
	if err := json.Unmarshal([]byte(resp), kkk); err != nil {
		log.Fatalf("JSON Unmarshalling failed: %s", err)
		return
	}
	student.Materials = kkk.Materials

	py, err := json.Marshal(student)
	if err != nil {
		log.Println(err)
		return
	}

	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/graphics/", bytes.NewBuffer(py))

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	t := &Python{}
	if err := json.NewDecoder(res.Body).Decode(t); err != nil {
		log.Println(err)
		return
	}

	file, err := os.ReadFile("./prompts/student.txt")
	if err != nil {
		log.Println(err)
		return
	}

	resp1, err := helpers.SendGPT(string(file), TranslateSubject(student), "gpt-4")
	if err != nil {
		log.Println(err)
		return
	}
	RR := &Prompt{}
	RR.Answer = resp1
	RR.Files = t.Files
	WriteJSON(w, http.StatusOK, RR)
}

func TeacherPage(w http.ResponseWriter, r *http.Request) {
	student := &Student{}
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		log.Println(err)
		WriteJSON(w, http.StatusOK, "\"error\": \"You should provide name and goal of the student\"")
		return
	}

	js, err := os.ReadFile("./prompts/template.txt")
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := helpers.SendGPT(string(js), student.Goal, "gpt-4")
	if err != nil {
		log.Println(err)
		return
	}
	kkk := &Data{}
	if err := json.Unmarshal([]byte(resp), kkk); err != nil {
		log.Fatalf("JSON Unmarshalling failed: %s", err)
	}
	student.Materials = kkk.Materials

	py, err := json.Marshal(student)
	if err != nil {
		log.Println(err)
		return
	}

	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/graphics/", bytes.NewBuffer(py))

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	t := &Python{}
	if err := json.NewDecoder(res.Body).Decode(t); err != nil {
		log.Println(err)
		return
	}

	file, err := os.ReadFile("./prompts/teacher.txt")
	if err != nil {
		log.Println(err)
		return
	}

	resp1, err := helpers.SendGPT(string(file), TranslateSubject(student), "gpt-4")
	if err != nil {
		log.Println(err)
		return
	}
	RR := &Prompt{}
	RR.Answer = resp1
	RR.Files = t.Files
	WriteJSON(w, http.StatusOK, RR)
}

func ParentPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		return
	}

	student := &Student{}
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		log.Println(err)
		WriteJSON(w, http.StatusOK, "\"error\": \"You should provide name and goal of the student\"")
		return
	}

	js, err := os.ReadFile("./prompts/template.txt")
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := helpers.SendGPT(string(js), student.Goal, "gpt-4")
	if err != nil {
		log.Println(err)
		return
	}
	kkk := &Data{}
	if err := json.Unmarshal([]byte(resp), kkk); err != nil {
		log.Fatalf("JSON Unmarshalling failed: %s", err)
	}
	student.Materials = kkk.Materials

	py, err := json.Marshal(student)
	if err != nil {
		log.Println(err)
		return
	}

	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/graphics/", bytes.NewBuffer(py))

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	t := &Python{}
	if err := json.NewDecoder(res.Body).Decode(t); err != nil {
		log.Println(err)
		return
	}

	file, err := os.ReadFile("./prompts/parent.txt")
	if err != nil {
		log.Println(err)
		return
	}

	resp1, err := helpers.SendGPT(string(file), TranslateSubject(student), "gpt-4")
	if err != nil {
		log.Println(err)
		return
	}
	RR := &Prompt{}
	RR.Answer = resp1
	RR.Files = t.Files
	WriteJSON(w, http.StatusOK, RR)
}
