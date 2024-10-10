package proxy

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	retaLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		retaLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.retaLimiter[url] == 0 {
		n.retaLimiter[url] = 1
	}
	if n.retaLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.retaLimiter[url] += 1
	return true
}
