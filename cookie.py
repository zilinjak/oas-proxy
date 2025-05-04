from http.server import BaseHTTPRequestHandler, HTTPServer

class MyHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # Send response status code
        self.send_response(200)

        # Send headers, including multiple Set-Cookie headers
        self.send_header('Content-type', 'text/html')
        self.send_header('Set-Cookie', 'sessionid=abc123; HttpOnly')
        self.send_header('Set-Cookie', 'userid=42; Path=/')

        # Must call end_headers() before writing body
        self.end_headers()

        # Write response body
        self.wfile.write(b"<html><body><h1>Hello, world!</h1></body></html>")

if __name__ == "__main__":
    PORT = 8000
    server = HTTPServer(('', PORT), MyHandler)
    print(f"Serving on http://localhost:{PORT}")
    server.serve_forever()
