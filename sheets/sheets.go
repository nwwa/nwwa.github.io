package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gsheets "google.golang.org/api/sheets/v4"
)

const (
	clientSecretsFile string = "client_secret.json"
	tokenCacheFile    string = "sheets_api_secret_cache"
)

// SignupInfo represents information from a user after filling the signup form
type SignupInfo struct {
	FirstName string
	LastName  string
	Addr      string
	City      string
	State     string
	Zip       string
	Phone     string
	Email     string
	Since     string
}

type Appender struct {
	// spreadsheetID is the Google Sheets document which will be appended to
	// with new user information
	spreadsheetID string

	// sheetsClient is the connection to Google Sheets which enables the server
	// to write new signups
	sheetsClient *gsheets.Service
}

func New(sheetID string) (Appender, error) {
	client, err := createSheetsClient()

	return Appender{
		spreadsheetID: sheetID,
		sheetsClient:  client,
	}, err
}

func (a *Appender) WriteNewSignup(info SignupInfo) error {
	myval := []interface{}{
		info.FirstName,
		info.LastName,
		info.Addr,
		info.City,
		info.State,
		info.Zip,
		info.Phone,
		info.Email,
		info.Since,
	}

	writeRange := "Members!A1"
	var vr gsheets.ValueRange
	vr.Values = append(vr.Values, myval)

	_, err := a.sheetsClient.Spreadsheets.Values.
		Append(a.spreadsheetID, writeRange, &vr).
		ValueInputOption("RAW").
		Do()
	return err
}

func createSheetsClient() (*gsheets.Service, error) {
	ctx := context.Background()

	b, err := ioutil.ReadFile(clientSecretsFile)
	if err != nil {
		return nil,
			fmt.Errorf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	config, err := google.ConfigFromJSON(b,
		"https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil,
			fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := gsheets.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets Client %v", err)
	}
	return srv, nil
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	tok, err := tokenFromFile(tokenCacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenCacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
