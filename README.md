## Number Classifier API
## Description
This project provides an API for number classification.


## Setup Instructions
To run this project locally, follow these steps:

### Prerequisites
- Go installed on your machine.
.
- Steps:
Clone the repository:

```bash
git clone https://github.com/your-username/your-project-name.git
cd your-project-name
```

Install dependencies:
- Install Gin using Go’s package manager:

```bash
go get github.com/gin-gonic/gin
```

## API Documentation
### Base URL
- All API endpoints are accessible via:

```
http://localhost:8080/api
```

### Endpoints

Get Number Info
- URL:  `/classify-number?number=42`
- Method: GET
  ##### Response Format:

```json
{
"number":42,
"is_prime":false,
"is_perfect":false,
"properties":["even"],
"digit_sum":6,
"fun_fact":"42 is the reciprocal of the sixth Bernoulli number."
}
```

- Example Usage
Fetching the Current Date and Time using cURL:
```bash
curl http://localhost:8080/api/classify-number?number=42
```

- Using JavaScript (fetch API):
```javascript
fetch('http://localhost:8080/api/classify-number?number=42')
  .then(response => response.json())
  .then(data => console.log(data));
```

### Backlink
For more information and to hire expert Golang developers, visit [HNG Tech - Hire Golang Developers](https://hng.tech/hire/golang-developers)

### License
This project is licensed under the MIT License - see the LICENSE file for details.
