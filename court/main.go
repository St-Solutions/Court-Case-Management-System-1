package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	"github.com/Surafeljava/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Welcome To Court Case Management System")

	dbc, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=courttest2 password=123456")
	defer dbc.Close()

	//TODO: Creating tables on the database
	// dbc.AutoMigrate(&entity.Opponent{})
	// dbc.AutoMigrate(&entity.Case{})
	// dbc.AutoMigrate(&entity.Judge{})
	// dbc.AutoMigrate(&entity.Admin{})
	// dbc.AutoMigrate(&entity.Notification{})
	// dbc.AutoMigrate(&entity.Relation{})
	// dbc.AutoMigrate(&entity.Decision{})

	// hasher := md5.New()
	// hasher.Write([]byte("1234"))
	// pwdnew := hex.EncodeToString(hasher.Sum(nil))

	// ad := entity.Admin{AdminId: "AD1", AdminPwd: pwdnew}
	// dbc.Create(&ad)

	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	//Login repository and Service Creation
	loginRepo := repository.NewLoginRepositoryImpl(dbc)
	loginServ := service.NewLoginServiceImpl(loginRepo)

	//Case repository and Service Creation
	caseRepo := repository.NewCaseRepositoryImpl(dbc)
	caseServ := service.NewCaseServiceImpl(caseRepo)

	//Opponent repository and Service Creation
	oppRepo := repository.NewOpponentRepositoryImpl(dbc)
	oppServ := service.NewOpponentServiceImpl(oppRepo)

	//Judge repository and Service Creation
	adminJudgeRepo := repository.NewJudgeRepositoryImpl(dbc)
	adminJudgeServ := service.NewJudgeServiceImpl(adminJudgeRepo)

	loginHandler := handler.NewLoginHandler(tmpl, loginServ)
	newcaseHandler := handler.NewCaseHandler(tmpl, caseServ)
	opponentHandler := handler.NewOpponentHandler(tmpl, oppServ)
	adminJudgeHandler := handler.NewAdminJudgeHandler(tmpl, adminJudgeServ)

	fs := http.FileServer(http.Dir("../UI/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/login", loginHandler.AuthenticateUser)
	http.HandleFunc("/admin/cases/new", newcaseHandler.NewCase)
	http.HandleFunc("/admin/cases/update", newcaseHandler.UpdateCase)
	http.HandleFunc("/admin/cases/delete", newcaseHandler.DeleteCase)
	http.HandleFunc("/admin/cases", newcaseHandler.Cases)
	http.HandleFunc("/admin/opponent/new", opponentHandler.NewOpponent)
	http.HandleFunc("/admin/judge/new", adminJudgeHandler.NewJudge)

	http.HandleFunc("/judge/cases/close", newcaseHandler.CloseCase)

	//TODO: notification handlers
	// http.HandleFunc("/admin/notification/new", )
	// http.HandleFunc("/notification", )

	http.ListenAndServe(":8181", nil)

}
