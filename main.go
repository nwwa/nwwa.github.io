package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

var (
	masterTmpl     *template.Template
	publishableKey = os.Getenv("PUBLISHABLE_KEY")
)

func handleSignup(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.Must(masterTmpl.Clone()).ParseFiles("signup.html")
	tmpl.Execute(w, map[string]string{"Key": publishableKey, "title": "testing"})
}

func handleCharge(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	customerParams := &stripe.CustomerParams{Email: r.Form.Get("stripeEmail")}
	customerParams.SetSource(r.Form.Get("stripeToken"))

	newCustomer, err := customer.New(customerParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chargeParams := &stripe.ChargeParams{
		Amount:   500,
		Currency: "usd",
		Desc:     "Sample Charge",
		Customer: newCustomer.ID,
	}

	if _, err := charge.New(chargeParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Charge completed successfully!")
}

func main() {
	stripe.Key = os.Getenv("SECRET_KEY")

	masterTmpl, _ = template.ParseFiles("master.html")
	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/charge", handleCharge)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":1313", nil)
}
