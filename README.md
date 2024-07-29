# ascii-art-web

`ascii-art-web` from [01Talent/Zone01 Oujda curriculum](https://01talent.com/) is a Go-based web application that allows users to generate ASCII art from input text and selected banner styles. The web interface provides a graphical way to create and view ASCII art, leveraging different banner styles.

## Description

`ascii-art-web` is designed to provide an interactive web interface for generating ASCII art using predefined banner styles. Users can input text and select a banner style, and the server will return the generated ASCII art. The application includes:
- A main page with a form for text input and banner selection.
- The ability to handle POST requests to generate ASCII art based on user input.
- Appropriate HTTP responses and status codes for different scenarios.

## Authors

- [Ayoub El Haddad](https://learn.zone01oujda.ma/git/aelhadda)

## Usage

To run the `ascii-art-web` server:

1. Clone the repository:
   ```sh
   git clone https://github.com/e-aub/ascii-art-web
   ```

2. Navigate to the project directory:
   ```sh
   cd ascii-art-web
   ```

3. Build and run the server:

   ```sh
   go run main.go
   ```
4. Open your web browser and navigate to http://localhost:8080 to access the web interface.

## Implementation Details
The application consists of the following main components:

### Handlers
`AsciiHandler:` Handles POST requests to /ascii-art to generate ASCII art. It validates the banner style, generates ASCII art using the asciiArt package, and handles errors appropriately.
### Endpoints
`GET /:` Serves the main page with a form for user input. This uses Go templates to render the page.

`POST /ascii-art:` Accepts form data (text and banner) and generates ASCII art. Redirects to the home page or displays an error based on the result.

### Error Handling
`500 Internal Server Error:` Returned if there is an issue during ASCII art generation.

`400 Bad Request:` Returned if the provided banner is invalid.

`405 Method Not Allowed:` Returned if a non-POST request is made to the /ascii-art endpoint or a non-GET request is made to the / endpoint.

### Instructions
1. Ensure that Go is installed on your system.
2. Place your HTML templates in the templates directory at the project root.
3. Follow the usage instructions to build and run the server.
4. Navigate to http://localhost:8080 in your web browser to use the application.
### HTTP Status Codes
`200 OK:` Returned for successful requests.

`404 Not Found:` Returned if the requested resource (e.g., template or banner) is not found.

`400 Bad Request:` Returned for invalid requests.

`500 Internal Server Error:` Returned for unhandled errors.

Contributing
If you find any issues or have suggestions for improvements, please submit an issue or propose a change via a pull request.

For any questions or further assistance, feel free to reach out.
