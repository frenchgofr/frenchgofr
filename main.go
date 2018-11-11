package main

import (
	"expvar"
	"fmt"
    "github.com/frenchgofr/frenchgofr/domain"
    "log"
	"os"
	"text/template"
	"time"

	"bytes"
	"github.com/go-recaptcha/recaptcha"
	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
	"github.com/paulbellamy/ratecounter"
	"net"
	"net/http"
)

var indexTemplate = template.Must(template.New("index.tmpl").ParseFiles("inviter/templates/index.tmpl"))

var (
	api     *slack.Client
	captcha *recaptcha.Recaptcha
	counter *ratecounter.RateCounter

	ourTeam = new(domain.Team)

	m *expvar.Map
	hitsPerMinute,
	requests,
	inviteErrors,
	missingFirstName,
	missingLastName,
	missingEmail,
	missingCoC,
	successfulCaptcha,
	failedCaptcha,
	invalidCaptcha,
	successfulInvites,
	userCount,
	activeUserCount expvar.Int
)

var c Specification

// Specification is the config struct
type Specification struct {
	CaptchaSitekey string `required:"true"`
	CaptchaSecret  string `required:"true"`
	SlackToken     string `required:"true"`
	CocUrl         string `required:"false" default:"https://frenchgo.fr"`
	EnforceHTTPS   bool
	Debug          bool // toggles nlopes/slack client's debug flag
}

func main() {

	err := envconfig.Process("slackinviter", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	counter = ratecounter.NewRateCounter(1 * time.Minute)
	m = expvar.NewMap("metrics")
	m.Set("hits_per_minute", &hitsPerMinute)
	m.Set("requests", &requests)
	m.Set("invite_errors", &inviteErrors)
	m.Set("missing_first_name", &missingFirstName)
	m.Set("missing_last_name", &missingLastName)
	m.Set("missing_email", &missingEmail)
	m.Set("missing_coc", &missingCoC)
	m.Set("failed_captcha", &failedCaptcha)
	m.Set("invalid_captcha", &invalidCaptcha)
	m.Set("successful_captcha", &successfulCaptcha)
	m.Set("successful_invites", &successfulInvites)
	m.Set("active_user_count", &activeUserCount)
	m.Set("user_count", &userCount)

	captcha = recaptcha.New(c.CaptchaSecret)
	api = slack.New(c.SlackToken)

	if c.Debug {
		api.SetDebug(true)
	}

	go pollSlack()

    http.Handle("/", http.FileServer(http.Dir("blog/public")))
	http.HandleFunc("/inviter/invite/", handleInvite)
    http.Handle("/inviter/static/", http.StripPrefix("/inviter/static/", http.FileServer(http.Dir("./inviter/static"))))
	http.HandleFunc("/inviter", enforceHTTPSFunc(homepage))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func enforceHTTPSFunc(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if xfp := r.Header.Get("X-Forwarded-Proto"); c.EnforceHTTPS && xfp == "http" {
            u := *r.URL
            u.Scheme = "https"
            if u.Host == "" {
                u.Host = r.Host
            }
            http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
            return
        }
        h(w, r)
    }
}

// Updates the globals from the slack API
// returns the length of time to sleep before the function
// should be called again
func updateFromSlack() time.Duration {
	users, err := api.GetUsers()
	if err != nil {
		log.Println("error polling slack for users:", err)
		return time.Minute
	}
	var uCount, aCount int64 // users and active users
	for _, u := range users {
		if u.ID != "USLACKBOT" && !u.IsBot && !u.Deleted {
			uCount++
			if u.Presence == "active" {
				aCount++
			}
		}
	}
	userCount.Set(uCount)
	activeUserCount.Set(aCount)

	st, err := api.GetTeamInfo()
	if err != nil {
		log.Println("error polling slack for team info:", err)
		return time.Minute
	}
	ourTeam.Update(st)
	return 10 * time.Minute
}

// pollSlack over and over again
func pollSlack() {
	for {
		time.Sleep(updateFromSlack())
	}
}

// Homepage renders the homepage
func homepage(w http.ResponseWriter, r *http.Request) {
	counter.Incr(1)
	hitsPerMinute.Set(counter.Rate())
	requests.Add(1)

	var buf bytes.Buffer
	err := indexTemplate.Execute(
		&buf,
		struct {
			SiteKey,
			UserCount,
			ActiveCount string
			Team   *domain.Team
			CocUrl string
		}{
			c.CaptchaSitekey,
			userCount.String(),
			activeUserCount.String(),
			ourTeam,
			c.CocUrl,
		},
	)
	if err != nil {
		log.Println("error rendering template:", err)
		http.Error(w, "error rendering template :-(", http.StatusInternalServerError)
		return
	}
	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

// ShowPost renders a single post
func handleInvite(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	captchaResponse := r.FormValue("g-recaptcha-response")
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		failedCaptcha.Add(1)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	valid, err := captcha.Verify(captchaResponse, remoteIP)
	if err != nil {
		failedCaptcha.Add(1)
		http.Error(w, "Error validating recaptcha.. Did you click it?", http.StatusPreconditionFailed)
		return
	}
	if !valid {
		invalidCaptcha.Add(1)
		http.Error(w, "Invalid recaptcha", http.StatusInternalServerError)
		return

	}
	successfulCaptcha.Add(1)
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	coc := r.FormValue("coc")
	if email == "" {
		missingEmail.Add(1)
		http.Error(w, "Missing email", http.StatusPreconditionFailed)
		return
	}
	if fname == "" {
		missingFirstName.Add(1)
		http.Error(w, "Missing first name", http.StatusPreconditionFailed)
		return
	}
	if lname == "" {
		missingLastName.Add(1)
		http.Error(w, "Missing last name", http.StatusPreconditionFailed)
		return
	}
	if coc != "1" {
		missingCoC.Add(1)
		http.Error(w, "You need to accept the code of conduct", http.StatusPreconditionFailed)
		return
	}
	err = api.InviteToTeam(ourTeam.Domain(), fname, lname, email)
	if err != nil {
		log.Println("InviteToTeam error:", err)
		inviteErrors.Add(1)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successfulInvites.Add(1)
}
