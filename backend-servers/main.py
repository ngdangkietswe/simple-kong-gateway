from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/api/health', methods=['GET'])
def health():
    return jsonify({'message': 'Application is running in python backend server'}), 200

@app.route('/api/check-authorize', methods=['POST'])
def checkAuthorize():
    data = request.get_json()
    
    username = data.get('username')
    
    return jsonify({'message': 'User {} is authorized'.format(username)}), 200

if __name__ == '__main__':
    app.run(host="127.0.0.1", debug=True, port=5000)