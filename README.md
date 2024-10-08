# URL Shortener Server

## Start the server Locally

1. Make sure you have the necessary dependencies installed.
    - Go

2. Clone this repo into local
        git clone git@github.com:AvinashBarade/url-shortener.git

3. You can start the Go application by running:

        go run main.go

    This will start the server on `http://localhost:8080`.

4. Test URL Shortening
    
    You can use curl or a tool like Postman to test the endpoints

    Using curl:

    - Shorten a URL:
        
        Run this command in your terminal to shorten a URL:
            
            curl "http://localhost:8080/shorten?url=https://example.com"

        The server will respond with something like:

            Shortened URL: ab12cd

    - Redirect to Original URL:
        
        After shortening the URL, test the redirection by using the shortened URL (e.g., ab12cd):

            curl -I http://localhost:8080/ab12cd

        This will output the headers of the HTTP response. If the redirection works, you should see something like:

            HTTP/1.1 302 Found
            Location: https://example.com

    - Test Metrics API:

        Shorten multiple URLs and then test the `/metrics` endpoint to get the top 3 domains:

            curl http://localhost:8080/metrics

        The response should show the top 3 domains with the number of times they were shortened:

            example.com: 1
            another-domain.com: 3

5. Run Unit Tests
        
    - Use this command to run all the tests in your project:

            go test ./...

     - Check test coverage

            go test ./... -cover   

## Run URL Shortener as a Docker container 
1.  Pull the Docker Image
    
        docker pull avinashbarde/url-shortener:latest

2. Run the Docker Container

        docker run -p 8080:8080 avinashbarde/url-shortener

3. Now, you can test it in local using `curl` or `postman` on port 8080