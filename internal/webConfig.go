package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venuyeredla/pan-services/internal/handlers"
	"github.com/venuyeredla/pan-services/internal/models"
	"github.com/venuyeredla/pan-services/internal/repository"
)

func Logging(ginHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Requested URL before processing = %v", c.Request.URL.Path)
		ginHandler(c)
		log.Printf("After = %v", c.Request.URL.Path)
	}
}

func JwtValidate(ginHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		log.Printf("Request path is : %v", gctx.Request.URL.Path)
		//Authorization: Bearer <token>
		authHeader := gctx.Request.Header.Get("Authorization")

		if authHeader != "" {
			log.Printf("Authorization : %v", authHeader)
			token := authHeader[7:]
			claims := handlers.ParseToken(token)
			if handlers.IsValidToken(claims) {
				log.Printf("Token validation is sucessful")
				ginHandler(gctx)
			} else {
				log.Printf("Token validation is failed")
				eresp := &models.ErrorResponse{Msg: "Invalid credentials"}
				gctx.JSON(http.StatusUnauthorized, eresp)
			}

		} else {
			log.Printf("Token doesn't exist")
			eresp := &models.ErrorResponse{Msg: "Token doesn't exist"}
			gctx.JSON(http.StatusUnauthorized, eresp)
		}

	}
}

const API_CONTEXT = "/api/v1"

func ApiPath(uri string) string {
	return API_CONTEXT + uri
}

func ApplicationConfig() *gin.Engine {
	log.Println("Setting up file server and api handlers")

	repository.IntializePool()
	//defer dbop.ClosePool()

	router := gin.Default()

	//router.Group(API_CONTEXT)

	//Static file handler
	router.StaticFile("/", "web/index.html")
	router.StaticFS("/static/", http.Dir("web/static/"))

	sc := &handlers.StatiContent{}
	//sc.Preprocess()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/tmpl", Logging(sc.GenrateForm))
	/*router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))
	*/
	var customerAPI handlers.CustomerAPI
	router.POST(ApiPath("/auth/signin"), Logging(handlers.Authenticate))
	router.POST(ApiPath("/auth/signup"), Logging(handlers.SignUp))

	router.GET(ApiPath("/customer/info/:id"), JwtValidate(customerAPI.CustInfo))

	router.GET(ApiPath("/product/getproducts"), Logging(handlers.GetProducts))

	router.GET("/chat", handlers.WebChat)

	return router
}

/*
func ConfigureHttp() {

	func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
		log.Println(r.URL.Path)
	}
}
	log.Println("Setting up file server and api handlers")
	//Static file handler
	fileHandler := http.FileServer(http.Dir("wstatic/"))
	http.Handle("/wstatic/", http.StripPrefix("/wstatic/", fileHandler))

	//opening introduction page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wstatic/index.html")
	})

	//Generating entitre HTML using libraries
	sc := &handlers.StatiContent{}
	sc.Preprocess()
	http.HandleFunc("/form", sc.GenrateForm)

	var customerAPI handlers.CustomerAPI

	http.HandleFunc("/api/authenticate", logging(customerAPI.Authenticate))
	http.HandleFunc("/signup", logging(customerAPI.SignUp))
	http.HandleFunc("/custinfo", logging(customerAPI.CustInfo))


	http.ListenAndServe(":2023", nil)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

*/
