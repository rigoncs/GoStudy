package proxy

import "fmt"

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCpde, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCpde, body)

	httpCpde, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCpde, body)

	httpCpde, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCpde, body)

	httpCpde, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCpde, body)

}
