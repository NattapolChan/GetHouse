package main

import (
	//"github.com/gofiber/fiber/v2"
	"context"
	"fmt"
	"html/template"
	"net/http"

	//"io"
	"log"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/pat"
	"github.com/joho/godotenv"

	// "github.com/labstack/echo/v4"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/github"

	"github.com/NattapolChan/ent/user"

	"github.com/dgrijalva/jwt-go"
	// "github.com/NattapolChan/controllers"

	"github.com/NattapolChan/ent"

	"time"
	"strconv"

	"github.com/go-co-op/gocron"
	_ "github.com/mattn/go-sqlite3"

	"github.com/NattapolChan/schedule"
	"github.com/NattapolChan/search"
	"encoding/json"
)

func main() {

	s := gocron.NewScheduler(time.UTC)

	err_env := godotenv.Load(".env")

	if err_env != nil {
		log.Println("[Error]: .env -> %s", err_env)
	}

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	} else {
		log.Println("Successfully connecting to database")
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	// Run scheduled URA API call
	_, err_cron := s.Every(1).Days().DoWithJobDetails(func(job gocron.Job) {
		log.Printf("\nprevious:\t%s\nlast:\t\t%s\nnow:\t\t%s\nnext:\t\t%s\n\n\n", job.PreviousRun(), job.LastRun(), time.Now().UTC(), job.NextRun())
		/*
		for i := 0 ;i < 4; i++ {
			url := fmt.Sprintf("https://www.ura.gov.sg/uraDataService/invokeUraDS?service=PMI_Resi_Rental&refPeriod=22q%d", i)
			schedule.UpdatePropertyListing(client, url)
		}
		*/
		url := "https://www.ura.gov.sg/uraDataService/invokeUraDS?service=PMI_Resi_Rental&refPeriod=22q1"
		schedule.UpdatePropertyListing(client, url)
	})
	if err_cron != nil {
		panic(err_cron)
	}
	s.StartAsync()


	key := "Secret-session-key"

	maxAge := 86400 * 30
	isProd := false       // Set to true when serving over https


	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"), 
			os.Getenv("GOOGLE_CLIENT_SECRET"), 
			"http://localhost:8080/auth/google/callback", 
			"email", "profile",
		),
		github.New(
			os.Getenv("GITHUB_KEY"), 
			os.Getenv("GITHUB_SECRET"), 
			"http://localhost:8080/auth/github/callback",
			"email", "profile",
		),
	)

	p := pat.New()

	p.Get("/search", func(res http.ResponseWriter, req *http.Request) {
		search.HandleSearchRequest(res, req, client)
	})

	p.Get("/test", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("connected to backend"))
		res.WriteHeader(http.StatusOK)
	})

	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user_gothic, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			fmt.Print("Error (auth) : ", err)
			return
		}
		fmt.Println(user_gothic.Email)
		fmt.Println(user_gothic.Name)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user_gothic.Email,
		})

		tokenString, err := token.SignedString([]byte("<add-a-real-secret-here>"))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(tokenString)
		u, err_user := client.User.
				Query().
				Where(
					user.IDEQ(tokenString),
				).
				Only(context.Background())
		email := "default"
		if user_gothic.Email!="" {
			email = user_gothic.Email
		} else {
			email = fmt.Sprintf("github:/%s", user_gothic.Name)
		}
		if err_user != nil {
			fmt.Println("Error user : ", err_user)
			u, err_create_user := client.User.
						Create().
						SetID(tokenString).
						SetName(user_gothic.Name).
						SetEmailAddr(email).
						SetFavorites(nil).
						Save(context.Background())

			if err_create_user != nil {
				fmt.Println("Error during Create user: ", err_create_user)
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte(tokenString))
				return
			}
			fmt.Println("User ", u," created")

		} else { fmt.Println("User = ", u)}

		//res.Write([]byte(tokenString))
		res.Header().Set("Location", fmt.Sprintf("http://localhost:3000/redirect/?token=%s", tokenString))
		res.WriteHeader(http.StatusTemporaryRedirect)

	})

	p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
			return
		}
	})

	p.Get("/get-favorite", func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.URL.Query().Get("accesstoken")
		u, err_user := client.User.
				Query().
					Where(
						user.IDEQ(tokenString),
					).
					Only(context.Background())
		if err_user != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		propertylistings, err_prop := u.QueryPropertylistings().All(context.Background())

		if err_prop != nil {
			res.WriteHeader(http.StatusBadGateway)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Println(propertylistings)
		json.NewEncoder(res).Encode(propertylistings)
	})

	p.Get("/add-favorite", func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.URL.Query().Get("accesstoken")
		listId := req.URL.Query().Get("listid")

		id, err := strconv.Atoi(listId)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}

		u, err_user := client.User.
				Query().
					Where(
						user.IDEQ(tokenString),
					).
					Only(context.Background())

		u, err_update := u.
				Update().
				AddPropertylistingIDs(id).
				Save(context.Background())

		if err_user != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		if err_update != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		propertylistings, err_prop := u.QueryPropertylistings().All(context.Background())

		if err_prop != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Println(propertylistings)
		json.NewEncoder(res).Encode(propertylistings)
	})

	p.Get("/remove-favorite", func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.URL.Query().Get("accesstoken")
		listId := req.URL.Query().Get("listid")
		id, err := strconv.Atoi(listId)
		if err != nil {
			http.Error(res, "List id not sent correctly", http.StatusBadRequest)
		}
		u, err_user := client.User.
				Query().
					Where(
						user.IDEQ(tokenString),
					).
					Only(context.Background())

		u, err_update := u.
				Update().
				RemovePropertylistingIDs(id).
				Save(context.Background())

		if err_user != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		if err_update != nil {
			res.WriteHeader(http.StatusBadRequest)
		}

		propertylistings, err_prop := u.QueryPropertylistings().All(context.Background())

		if err_prop != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Println(propertylistings)
		json.NewEncoder(res).Encode(propertylistings)
	})

	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Println("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(p)))
	//log.Fatal(http.ListenAndServe(":8080", p))

}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var indexTemplate = `
{{range $key,$value:=.Providers}}
<p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}
`

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
