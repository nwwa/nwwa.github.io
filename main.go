package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
)

var (
	masterTmpl     *template.Template
	publishableKey = os.Getenv("PUBLISHABLE_KEY")
)

func logError(r *http.Request, err error, msg string) {
	// TODO: Do something with the request to give more context
	// Such as printing the email, phone, and name so I can contact them
	fmt.Printf("%s: %s\n", msg, err)
}

func renderSignupSuccess(w http.ResponseWriter) {
	tmpl, _ := template.Must(masterTmpl.Clone()).ParseFiles("templates/signup-success.html")
	tmpl.Execute(w, map[string]string{"title": "Membership Created"})
}

func renderSignupFailure(w http.ResponseWriter) {
	tmpl, _ := template.Must(masterTmpl.Clone()).ParseFiles("templates/signup-failure.html")
	tmpl.Execute(w, map[string]string{"title": "Membership Failed"})
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.Must(masterTmpl.Clone()).ParseFiles("templates/signup.html")
	tmpl.Execute(w, map[string]string{"Key": publishableKey, "title": "Membership Signup"})
}

func handleCharge(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	customerParams := &stripe.CustomerParams{Email: r.Form.Get("stripeEmail")}
	customerParams.SetSource(r.Form.Get("stripeToken"))

	newCustomer, err := customer.New(customerParams)
	if err != nil {
		logError(r, err, "Failed to create a customer during signup")
		renderSignupFailure(w)
		return
	}

	subParams := stripe.SubParams{
		Customer: newCustomer.ID,
		Items: []*stripe.SubItemsParams{
			{
				Plan: "membership",
			},
		},
	}
	if _, err := sub.New(&subParams); err != nil {
		logError(r, err, "Failed to create a subscription during signup")
		renderSignupFailure(w)
		return
	}

	renderSignupSuccess(w)
}

func main() {
	stripe.Key = os.Getenv("SECRET_KEY")

	masterTmpl, _ = template.ParseFiles("templates/master.html")
	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/charge", handleCharge)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":1313", nil)
}
