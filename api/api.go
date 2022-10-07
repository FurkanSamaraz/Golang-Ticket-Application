package api

import (
	"encoding/json"
	"fmt"
	"main/db"
	"main/isvalid"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Ticket_options(w http.ResponseWriter, r *http.Request) {
	var (
		db  = db.Openconnention()
		opt model.Ticket_options
	)

	opt.Name = r.FormValue("name")
	opt.Desc = r.FormValue("desc")
	optAlcain := r.FormValue("allocation")
	optAlcainInt, _ := strconv.Atoi(optAlcain)
	opt.Allocation = optAlcainInt
	if optAlcainInt < 0 {

		w.WriteHeader(http.StatusBadRequest)

	} else {

		opt.Create_Date = time.Now().String()

		db.Exec("INSERT INTO toptions(namee,descc,allocation,create_date) VALUES($1,$2,$3,$4)", opt.Name, opt.Desc, optAlcain, opt.Create_Date)
		peopleByte, err := json.MarshalIndent(opt, "", "\t")
		isvalid.CheckErr(err)
		w.Write(peopleByte)

	}

	defer db.Close()

}
func Ticket(w http.ResponseWriter, r *http.Request) {
	var (
		opt       model.Ticket_options
		db        = db.Openconnention()
		statusAll []model.Ticket_options
	)

	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := db.Query("SELECT * FROM toptions WHERE id=$1", id)

	for rows.Next() {
		rows.Scan(&opt.Id, &opt.Name, &opt.Desc, &opt.Allocation, &opt.Create_Date)

		statusAll = append(statusAll, opt)

	}
	peopleByte, err := json.MarshalIndent(statusAll, "", "\t")

	w.Write(peopleByte)
	isvalid.CheckErr(err)
	defer rows.Close()
	defer db.Close()
}
func Ticket_purchases(w http.ResponseWriter, r *http.Request) {
	var (
		prc       model.Ticket_purchases
		opt       model.Ticket_options
		statusAll []model.Ticket_options
		db        = db.Openconnention()
		ss        int
	)
	vars := mux.Vars(r)
	id := vars["id"]

	rows, err := db.Query("SELECT * FROM toptions WHERE id=$1", id)
	for rows.Next() {
		statusRows := rows.Scan(&opt.Id, &opt.Name, &opt.Desc, &opt.Allocation, &opt.Create_Date)
		if statusRows == nil {

			ss = opt.Allocation
		}
		statusAll = append(statusAll, opt)
	}

	prcQuty := r.FormValue("quantity")
	prc.User_id = r.FormValue("user_id")

	prcQ, err := strconv.Atoi(prcQuty)
	prc.Quantity = prcQ

	if ss >= prcQ {
		db.Exec("INSERT INTO tpurchases(quantity,user_id) VALUES($1,$2)", prc.Quantity, prc.User_id)
		peopleByte, err := json.MarshalIndent(prc, "", "\t")
		fmt.Println(peopleByte)
		isvalid.CheckErr(err)
		updateAl := ss - prcQ

		db.Exec("UPDATE toptions SET allocation=$1 WHERE id=$2", updateAl, id)
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(http.StatusText(http.StatusOK)))

	} else if ss <= prcQ {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))

	}
	isvalid.CheckErr(err)
	defer rows.Close()
	defer db.Close()
}

func Users(w http.ResponseWriter, r *http.Request) {
	var (
		db    = db.Openconnention()
		users model.Ticket_user
	)
	id := uuid.New()
	fmt.Println(id.String())

	users.User_id = id.String()
	users.Name = r.FormValue("name")

	db.Exec("INSERT INTO userss(user_id,name) VALUES($1,$2)", users.User_id, users.Name)
	peopleByte, err := json.MarshalIndent(users, "", "\t")
	isvalid.CheckErr(err)
	w.Write(peopleByte)

	defer db.Close()
}
