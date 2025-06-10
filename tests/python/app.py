# This app is used as a test for the OAS Proxy, hence written in flask that does not validate the request body. 
from flask import Flask, request, jsonify
import json

app = Flask(__name__)

@app.route('/<integer>/<string>', methods=['POST'])
def get_item(integer, string):
    request_data = request.get_json()
    boolean_value = request_data.get('boolean', False)
    q = request.args.get('q')

    response = {
        'integer': integer,
        'string': string,
        'boolean': boolean_value
    }

    header_string = "\n".join(f"{k}: {v}" for k, v in request.headers.items())
    if q is None:
        q = ""
    else:
        q = f"?{q=}"
    print(f"POST /{integer}/{string}{q}\n"
        f"Headers:\n"
        f"{header_string}\n"
        f"Body: {json.dumps(request_data, indent=4)}\n"
        f"Response: {json.dumps(response, indent=4)}\n"
    )
    
    return jsonify(response)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)