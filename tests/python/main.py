from http.server import HTTPServer, BaseHTTPRequestHandler
import json
from urllib.parse import urlparse, parse_qs
import re

class SimpleAPIHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # Parse path parameters
        item_id_match = re.match(r'/items/(\d+)', self.path)
        if not item_id_match:
            self.send_response(404)
            self.end_headers()
            return

        item_id = int(item_id_match.group(1))

        # Parse query parameters
        parsed_path = urlparse(self.path)
        query_params = parse_qs(parsed_path.query)
        color = query_params.get('color', [None])[0]

        # Parse headers
        request_id = self.headers.get('X-Request-ID')

        # Parse cookies
        cookies = {}
        if 'Cookie' in self.headers:
            for cookie in self.headers['Cookie'].split(';'):
                if '=' in cookie:
                    key, value = cookie.strip().split('=', 1)
                    cookies[key] = value
        session_id = cookies.get('session_id')

        # Parse JSON body (for GET this is unusual, but we'll handle it)
        content_length = int(self.headers.get('Content-Length', 0))
        post_data = self.rfile.read(content_length)
        try:
            json_data = json.loads(post_data) if content_length else {}
        except:
            json_data = {}

        # Prepare response
        response = {
            "itemId": item_id,
            "color": color,
            "hasDetails": json_data.get('details', False),
            "requestId": request_id,
            "sessionId": session_id
        }

        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps(response).encode('utf-8'))

if __name__ == '__main__':
    server_address = ('', 8000)
    httpd = HTTPServer(server_address, SimpleAPIHandler)
    print('Server running at http://localhost:8000')
    httpd.serve_forever()
