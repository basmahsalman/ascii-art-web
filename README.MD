# ascii-art-web

## Description
Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of the last project, ascii-art.
Ascii-art-web takes an input (text) and banner (font), and disply the output as a converted string into ascii art.


## Authors
@fhamza https://learn.reboot01.com/git/fhamza
@bsalman https://learn.reboot01.com/git/bsalman
@zasaed https://learn.reboot01.com/git/zasaeed

## Usage: how to run
-run the server: $go run . server.go
-open the webpage: localhost:8080
-enter a text, choose a banner, and submit

## Implementation details: algorithm
the following HTTP endpoints are implemented:

GET /: Sends HTML response, the main page.
1.1. GET Tip: go templates to receive and display data from the server.

POST /ascii-art: that sends data to Go server (text and a banner)
2.1. POST Tip: use form and other types of tags to make the post request.\

the result is displayed in the route /ascii-art after the POST is completed. So going from the home page to another page.

The main page have:
text input
select object to switch between banners
button, which sends a POST request to '/ascii-art' and outputs the result on the page.

HTTP status code
The endpoints return appropriate HTTP status codes.
OK (200), if everything went without errors.
Not Found, if nothing is found, for example templates or banners.
Bad Request, for incorrect requests.
Internal Server Error, for unhandled errors.