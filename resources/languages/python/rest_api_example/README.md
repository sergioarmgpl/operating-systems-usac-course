# Python REST API example.

 APIs can be written on Python pretty quick and easy using Flask. However, we also have Flask Restful which is
an extension for Flask that adds support for quickly building REST APIs.

## Requirements
 - Python 3.x
   * Can be installed from [https://www.python.org/](https://www.python.org/)
 - flask
   * Can be installed using the command
   ```
    pip install flask
   ```
 - flask-restful
* Can be installed using the command
   ```
    pip install flask-restful
   ```

## About this example
This example API has versioning. So v1 uses Flask along Flask Restful and v2 uses Flask only.
Routes available on this example are:
```
/* RESTfull API fully available*/
/api/v1/books

/* Only get all items endpoint available for demo purposes*/
/api/v2/books
```

The development server runs by default on port 5000 and can be started using
```bash
    python api.py
```

Details and clarification about implementation can be found on the documentation of the api.py file itself.

A Postman Collection was shared along with this example, so you can easily test the endpoints.
If you preffer to use cURL, endpoints can be reach using:

GET
> curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:5000/api/v1/books

POST

> curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "New Book Created"}' \
  http://localhost:5000/api/v1/books

PUT
> curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"name": "Edit Book Created"}' \
  http://localhost:5000/api/v1/books/4

DELETE
>curl --header "Content-Type: application/json" \
  --request DELETE \
  http://localhost:5000/api/v1/books/4