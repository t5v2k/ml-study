from flask import Flask
import os

app = Flask(__name__)

@app.route('/')
def hello_world():
  return f'Hello, World!'

@app.route('/about')
def about():
  return f'This is the about page'

@app.route('/user/<username>')
def show_user_profile(username):
  return f'User {username}'

if __name__ == '__main__':
  app.run(host='0.0.0.0', port=5000)
