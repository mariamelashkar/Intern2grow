package main

import (
    "io"
    "log"
    "net/http"
    "net/url"
)


// Handle requests and forward them to the target server
func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
    // Extract the target URL from the request
    requestURL := r.URL.Query().Get("request_url")
    if requestURL == "" {
        http.Error(w, "request_url query parameter is required", http.StatusBadRequest)
        return
    }

    // Parse the target URL
    proxyURL, err := url.Parse(requestURL)
    if err != nil {
        http.Error(w, "Invalid request_url", http.StatusBadRequest)
        return
    }

    // Create a new request to the target URL
    proxyReq, err := http.NewRequest(r.Method, proxyURL.String(), r.Body)
    if err != nil {
        http.Error(w, "Error creating request", http.StatusInternalServerError)
        return
    }

  

    // Remove headers that could reveal client information
    proxyReq.Header.Del("User-Agent")
    proxyReq.Header.Del("Referer")
    proxyReq.Header.Del("X-Forwarded-For")
    proxyReq.Header.Del("X-Real-IP")

    // Copy remaining headers from the original request
    for name, values := range r.Header {
        if name != "User-Agent" && name != "Referer" && name != "X-Forwarded-For" && name != "X-Real-IP" {
            for _, value := range values {
                proxyReq.Header.Add(name, value)
            }
        }
    }

   // Optionally, set your own IP address or a placeholder in X-Forwarded-For to hide client's IP
   proxyReq.Header.Set("X-Forwarded-For", "127.0.0.1") // Example of using a placeholder IP

    // Forward the request to the target server
    client := &http.Client{}
    resp, err := client.Do(proxyReq) // used to send an HTTP request and receive an HTTP response.
    if err != nil {
        http.Error(w, "Error forwarding request", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Copy headers and status code from the response
    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body) //This function copies data from a source (resp.Body) to a destination (w).
}

// This forward proxy server code handles incoming HTTP requests by extracting the target URL from a query parameter, 
// forwarding the request to the target server, and then relaying the response back to the client. Here are the steps in summary:
// Extract Target URL: Get the request_url query parameter from the request.
// Parse URL: Parse the request_url to ensure it's valid.
// Create Request: Create a new HTTP request to the target URL with the same method and body as the original request.
// Copy Headers: Copy headers from the original request to the new request.
// Send Request: Use an HTTP client to send the new request to the target server.
// Forward Response: Copy the response headers and body from the target server to the client's response, ensuring the client receives the correct data and status code.

func main() {
    // Set up the HTTP handler
    http.HandleFunc("/", handleRequestAndRedirect)
    
    // Define the address and port to run the proxy server
    addr := "127.0.0.1:8080"
    log.Println("Starting proxy server on", addr)
    
    // Start the HTTP server
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
