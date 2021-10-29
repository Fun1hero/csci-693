package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	firebaseconn "web-survey/firebase-conn"
	quizquestions "web-survey/quiz-questions"
)

var (
	responses    = make(map[string]interface{})
	nextQuestion = "general"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/survey", SurveyPage)
	http.HandleFunc("/surveyAdditional", AdditionalSurveyPage)
	http.HandleFunc("/submit", GeneralForm)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	LoadPageFromHTML(&w, "index.html", quizquestions.Questions{})
}

func SurveyPage(w http.ResponseWriter, r *http.Request) {
	LoadPageFromHTML(&w, "survey.html", quizquestions.MyQuestions)
}

func AdditionalSurveyPage(w http.ResponseWriter, r *http.Request) {
	LoadPageFromHTML(&w, "survey.html", quizquestions.Languages[nextQuestion])
}

func GeneralForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	language := ""
	for _, v := range quizquestions.Languages[nextQuestion].QuestionList {
		value := r.Form[strconv.Itoa(v.Id)][0]
		responses[v.Text] = value
		if _, ok := quizquestions.Languages[value]; ok {
			language = value
		}
	}
	if nextQuestion == "general" {
		nextQuestion = language
		http.Redirect(w, r, "/surveyAdditional", http.StatusFound)
	} else {
		firebaseconn.SaveResponse(responses)
		LoadPageFromHTML(&w, "thanks.html", quizquestions.Questions{})
	}
}

func LoadPageFromHTML(w *http.ResponseWriter, templateName string, data quizquestions.Questions) {
	t, err := template.ParseFiles(templateName) //parse the html file homepage.html
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(*w, data) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {           // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
