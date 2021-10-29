package quizquestions

var Languages = map[string]Questions{
	"general": MyQuestions,
	"c++":     CppQuestions,
	"java":    JavaQuestions,
	"python":  PythonQuestions,
	"other":   OtherQuestions,
}

type Question struct {
	Id   int
	Text string
}

type Questions struct {
	QuestionList []Question
	Answers      [][]string
}

var (
	MyQuestions = Questions{
		QuestionList: []Question{
			{0, "Have you ever used Golang before?"},
			{1, "It was easy to understand the material and to finish the challenge?"},
			{2, "Was it helpful to have previous experience in programming to understand the basics of Golang?"},
			{3, "How many years of prior experience do you have with any other programming languages?"},
			{4, "Have you used any additional resources for completing the challenge?"},
			{5, "Would you consider integrating Golang in daily usage?"},
			{6, "I have a clear understanding of where to apply acquired Golang knowledge?"},
			{7, "What is your primary language for writing code?"},
		},
		Answers: [][]string{
			{"weekly", "monthly", "once per year", "never"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"0-3", "3-5", "5+"},
			{"yes", "no"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"python", "c++", "java", "other"},
		},
	}
	CppQuestions = Questions{
		QuestionList: []Question{
			{0, "Considering your background in C++, have you felt a lot of similarities writing code in Go?"},
			{1, "Do you think variable declarations in Go is an advancement over C++?"},
			{2, "Would you recommend Go as an lightweight substitute of C++?"},
		},
		Answers: [][]string{
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
		},
	}
	JavaQuestions = Questions{
		QuestionList: []Question{
			{0, "Did you like the idea of not using semicolons on the end of every line?"},
			{1, "Do you feel like Go has a more modern syntax?"},
			{2, "Did you find the concept of functions in Go easy to pickup compared to Java?"},
		},
		Answers: [][]string{
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
		},
	}
	PythonQuestions = Questions{
		QuestionList: []Question{
			{0, "Given your knowledge of python did you fell Go syntax?"},
			{1, "Would you recommend Go for a beginner programmer?"},
			{2, "The way code is structured in Go is comprehensive and modern?"},
		},
		Answers: [][]string{
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
		},
	}
	OtherQuestions = Questions{
		QuestionList: []Question{
			{0, "Was there a lot of in common between Go and your primary programming language?"},
			{1, "Did you have a different learning experience compared to your primary programming language?"},
			{2, "Would you write less code using your primary programming language to perform the same task?"},
		},
		Answers: [][]string{
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
			{"strongly agree", "agree", "disagree", "strongly disagree"},
		},
	}
)
