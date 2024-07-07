GoLang Web Server

This project is a simple web server built with GoLang.

Features

HTTP Server: Implements a basic HTTP server using Go's standard library.
Routing: Demonstrates routing using http.Handle and http.HandleFunc.
Static File Serving: Serves static files using http.FileServer.
Form Handling: Includes basic form handling with http.ParseForm.
Prerequisites

Before running this server locally, make sure you have:

Go installed on your machine. You can download it from golang.org.
Clone this repository to your local machine.

Usage

Endpoints
Root: http://localhost:8080/

Serves a welcome message or homepage.
Hello Route: http://localhost:8080/hello

Demonstrates a simple route with plain text response.
Form Route: http://localhost:8080/form

Shows basic form handling and post request will be send to "/form" and will display the response.
Static Files: http://localhost:8080/form.html

Serves static files (CSS, JavaScript, images) from the ./static directory.
