# Python REST API example.

# APIs can be written on Python pretty quick and easy using Flask. However we also have Flask Restful which is
# an extension for Flask that adds support for quickly building REST APIs.

# Example API has versioning. So v1 uses Flask along Flask Restful and v2 uses Flask only.
# You can run this example with the command "python api.py".
# A Postman Collection was shared along with this example so you can easily test the endpoints.

# First step is to import Flask and Flask Restful extension.
# - reqparse enables adding and parsing of multiple arguments.
# - abort raise a HTTPException for the given http_status_code.
# - Api is the main entry point for the application and must be initialized with flask application.
# - Resource represents an abstract restful resource.
from flask import Flask
from flask_restful import reqparse, abort, Api, Resource

# app is initialized using flash and the main entry point for the API es configured.
app = Flask(__name__)
api = Api(app)

# reqparse in configured. 'name' is the argument to parse.
parser = reqparse.RequestParser()
parser.add_argument('name')

# This example doesn't have db or any persist configuration, so it would be simulated using a dictionary.
books = {
    "1": {'name': 'Harry Potter'},
    "2": {'name': 'Game of Thrones'},
    "3": {'name': 'The Hunger Games'},
}


# This method is written in order to check if data exists and returns 404 if it doesn't.
def check_if_exists(book_id):
    if book_id not in books:
        abort(404, message='Book {} does not exist'.format(book_id))


# This class represents a resource: Book. It needs to contain get, delete and put methods. These methods represents
# the HTTP verbs that are going to be supported by the API. Logic inside the methods depends on the operation we
# want to perform, but usually they have small configuration and code. Validations and complex logic could be
# added through middlewares and/or services, this depends on the project architecture.
class Book(Resource):
    def get(self, book_id):
        check_if_exists(book_id)
        return books[book_id]

    def delete(self, book_id):
        check_if_exists(book_id)
        del books[book_id]
        return '', 200

    def put(self, book_id):
        args = parser.parse_args()
        book = {'name': args['name']}
        books[book_id] = book
        return book, 200


# This class also represents a resource. The difference between this one and Book, is that it performs operations
# on the complete list of resources instead of the item itself, like the previous one.
class BookList(Resource):
    def get(self):
        return books

    def post(self):
        args = parser.parse_args()
        book_id = str(int(max(books.keys())) + 1)
        books[book_id] = {'name': args['name']}
        return books[book_id], 200


# The resources classes are added to the api object defined at the beginning along with the route that resolves them.
api.add_resource(BookList, '/api/v1/books')
api.add_resource(Book, '/api/v1/books/<book_id>')


# This is not needed for the rest API itself, but it's recommended in order to handle unknown routes.
@app.errorhandler(404)
def page_not_found(e):
    return "<h1>404</h1><p>The resource could not be found.</p>", 404


# This implementation was added in order to show how to write an endpoint without using flask restful extension.
# methods array support more than one HTTP verb at the same time, but logic could get complex. And models handling
# must be added manually.
@app.route('/api/v2/books', methods=['GET'])
def get_all():
    return books

# This runs the server in default port 5000. Server is not intended to use in production environment.
if __name__ == '__main__':
    app.run(debug=True)

