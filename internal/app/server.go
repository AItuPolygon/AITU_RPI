package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	openapiv1 "rpi-test/pkg/openapi/v1"
	"syscall"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
)

type App struct {
	Config  *Config
	Classes map[string]map[string]*ClassFieldConfig
	Tasks   []*Task
	Port    string
}

func Run() {
	var err error
	app := new(App)
	err = app.InitConfig()
	if err != nil {
		log.Println("Config create error:", err)
		return
	}

	err = app.InitTasks()
	if err != nil {
		log.Println("Init tasks error:", err)
		return
	}

	for _, task := range app.Tasks {
		fmt.Println(task.Name, task.config.Elements)
	}

	err = app.InitModbus()
	if err != nil {
		log.Println("Init modbus error:", err)
		return
	}
	// openapiv1.StrictHandlerFunc
	// fmt.Printf("Tasks: %+v\n", tasks)

	// tlsConfig, err := app.InitTls()
	// if err != nil {
	// 	log.Println("tls create error:", err)
	// 	return
	// }

	//TODO: TESTING!!!!
	// for _, task := range app.Tasks {
	// 	if task.Name == "airport" {
	// 		p := make(map[string]interface{})
	// 		p["lights_id"] = []int{0x10, 0x11, 0x12, 0x14}
	// 		app.processApiCommand(task, "turn_on", p)
	// 		break
	// 	}
	// }
	// ==================

	mux := echo.New()
	srv := newServer(app)
	swagger, err := openapiv1.GetSwagger()
	if err != nil {
		log.Fatalf("Failed to load swagger: %v", err)
	}
	mux.Use(middleware.OapiRequestValidator(swagger))
	mux.Use(TaskElementExtractMiddleware)

	strictServer := openapiv1.NewStrictHandler(srv, []openapiv1.StrictMiddlewareFunc{})
	openapiv1.RegisterHandlers(mux, strictServer)

	// srv := &http.Server{
	// 	Addr:    ":" + strconv.Itoa(app.Config.App.Port),
	// 	Handler: mux,
	// 	// TLSConfig:    tlsConfig,
	// 	// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	// }

	errChan := make(chan error, 1)
	go func() {
		log.Printf("Starting server on port: %d\n", app.Config.App.Port)
		// if err := srv.ListenAndServeTLS(app.Config.Cert.ServerCert.PublicKey, app.Config.Cert.ServerCert.PrivateKey); err != nil && err != http.ErrServerClosed {
		// 	errChan <- err
		// 	return
		// }
		port := fmt.Sprintf(":%d", app.Config.App.Port)
		if err := mux.Start(port); err != nil && err != http.ErrServerClosed {
			errChan <- err
			return
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-errChan:
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("Error!!!")
			ctx, cancel := context.WithTimeout(context.Background(), app.Config.App.ShutdownTimeout)
			defer cancel()
			if err := mux.Shutdown(ctx); err != nil {
				log.Printf("Server forced to shutdown: %s", err)
				return
			}
		}
	case <-quit:
	}

}
