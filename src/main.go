package main

import (
	"github.com/microservice_template/src/core"
)

const serviceName = "service_A"

func main() {


	// define plugins - dependencies/infrastructure for all further processing
	// (including all required configuration, CLI flags etc.)
	//metricsPlugin := metrics_plugin.New(serviceName)
	// note: most of the plugins are self-contained and independent from each other
	//dbPlugin := db_plugin.New(serviceName)
	// note: some of the plugins cannot be fully decoupled from the processing logic itself...
	//restPlugin := restapi_plugin.New(serviceName, func(router chi.Router) {
	//	router.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
	//		// order of plugin definition matters - we can use previously defined ones here:
	//		writer.Write([]byte("test + " + dbPlugin.Name()))
	//	})
	//})

	// business logic flow is also a plugin! (and could use any other plugins it's built on top of)
	//flowPlugin := NewFlowPlugin(dbPlugin)

	// define application - app lifecycle manager, that would register & initialize all the plugins and flows
	//app := bootstrap.New(serviceName, []bootstrap.Plugin{metricsPlugin, dbPlugin, restPlugin, flowPlugin})
	app := core.New()
	//ctx, ctxCancel := context.WithCancel(context.Background())
	// starting the app means: initialize plugins (setup infrastructure), start all flows, handle interruption signals etc.
	//app.Start(ctx, ctxCancel)
	app.Start()
}
