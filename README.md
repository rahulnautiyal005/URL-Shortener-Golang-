# 🔗   URL Shortener (Golang)  

A simple and efficient URL Shortener service built using Go. This project allows users to convert long URLs into short, shareable links and redirect back to the original URL.

---

## 🚀 Features

* 🔗 Shorten long URLs into compact links
* 🔁 Redirect short URLs to original URLs
* 📊 Track number of clicks (optional)
* ⏳ Support for URL expiration (optional)
* 🧠 Unique ID generation using hashing/base62

---

## 🛠️ Tech Stack

* Language: Go
* Database: MongoDB / PostgreSQL (configurable)
* HTTP Server: net/http
* Optional: Redis (for caching)

---

## 📁 Project Structure

```
url-shortener/
│── main.go
│── handlers/
│     ├── shorten.go
│     ├── redirect.go
│── models/
│     └── url.go
│── utils/
│     └── base62.go
│── database/
│     └── db.go
│── go.mod
```

---

## ⚙️ How It Works

1. User sends a long URL to the `/shorten` endpoint
2. Server generates a unique short ID
3. The mapping is stored in the database
4. When user visits `/abc123`, it redirects to the original URL

---

## 📌 API Endpoints

### 🔹 Shorten URL

```
POST /shorten
```

**Request Body:**

```json
{
  "url": "https://example.com/very-long-link"
}
```

**Response:**

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

---

### 🔹 Redirect

```
GET /{short_id}
```

➡️ Redirects to original URL

---

## ▶️ Run Locally

### 1. Clone the repository

```
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```

### 2. Install dependencies

```
go mod tidy
```

### 3. Run the server

```
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 🧪 Example Usage

```
POST /shorten
→ returns short link

Visit:
http://localhost:8080/xyz123
→ redirects to original URL
```

---

## ✨ Future Improvements

* User authentication (JWT)
* Custom short URLs
* Analytics dashboard
* Rate limiting
* Distributed system using Redis

---

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first.

---

## 📄 License

This project is open-source and available under the MIT License.

---

## 👨‍💻 Author

Rahul Nautiyal
