package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define the application struct to hold the application-wide dependencies for the
// web application. For now we'll only include the structured logger, but we'll
// add more to this as development progresses.
type application struct {
	logger *slog.Logger
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr variable.
	// You need to call this *before* you use the addr variable otherwise it will always
	// contain the default value of ":4000". If any errors are encountered during parsing
	// the application will be terminated.
	flag.Parse()

	// Use the slog.New() function to initialize a new structured logger, which writes
	// to the stadard out stream and uses the default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// Initialize a new instance of our application struct, containing the dependencies
	// (for now, just the structured logger)
	app := &application{
		logger: logger,
	}

	// Use the Info() method to log the starting server message at Info severity (along
	// with the listen address as an attribute)
	logger.Info("Starting server", slog.String("addr", *addr))

	// Call the new app.routes() method to get the servemux containing our routes,
	// and pass that to http.ListenAndServe()
	err := http.ListenAndServe(*addr, app.routes())

	// And we also use the Error() method to log any error message returned by
	// http.ListenAndServe() at Error severity (with no additional attributes),
	// and then call os.Exit(1) to terminate the application with exit code 1.
	logger.Error(err.Error())
	os.Exit(1)
}
