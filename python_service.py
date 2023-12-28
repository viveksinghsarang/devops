# Install required packages:
# pip install flask

from flask import Flask, request, jsonify
import logging

app = Flask(__name__)

@app.route('/process', methods=['POST'])
def process_data():
    data = request.json['data']
    sorted_data = sorted(data)

    # Log the request and response
    logging.info(f'Received data: {data}')
    logging.info(f'Sorted data: {sorted_data}')

    return jsonify({'result': sorted_data})

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    app.run(host="0.0.0.0", port=9001)

