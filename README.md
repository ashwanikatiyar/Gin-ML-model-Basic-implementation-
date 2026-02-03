Here is the complete text formatted as a clean, ready-to-use Markdown (`.md`) file. I've cleaned up the structure, properly nested the code blocks, and formatted the math using standard LaTeX for GitHub-style rendering.

```markdown
# 🧠 Sentiment Analysis API (ML from Scratch in Go)

A lightweight **Sentiment Analysis REST API** built using **Go (Gin)** with a **Logistic Regression model implemented from scratch** (no ML libraries).

This project demonstrates how to design, train, and serve a simple machine learning model as a backend service, with a clear separation between **ML logic** and **API infrastructure**.

---

## ✨ Features

- 🚀 **REST API** built with Gin
- 🧮 **Logistic Regression** implemented from scratch
- 📝 **Simple Bag-of-Words** text vectorization
- 🔁 **Clear separation** of training and prediction
- 📦 **No ML frameworks** (no sklearn, no TensorFlow)
- 🧠 **Fully explainable** and interpretable model

---

## 📌 Problem Statement

Given a short piece of text, predict whether its sentiment is:
- **Positive**
- **Negative**

The model outputs:
- A probability score (`0.0 – 1.0`)
- A sentiment label based on the score

---

## 🏗️ Project Structure

```text
sentiment-api/
│
├── main.go            # Gin server entry point
├── go.mod
│
├── api/
│   └── handlers.go    # HTTP request handlers
│
├── ml/
│   ├── vectorizer.go  # Bag-of-Words implementation
│   ├── model.go       # Logistic Regression model
│   ├── train.go       # Gradient Descent training
│   └── state.go       # In-memory model state
│
└── README.md

```

---

## 🧠 Machine Learning Details

### Model

* **Logistic Regression**: Binary classification (positive / negative).

### Text Vectorization

* **Bag-of-Words**:
* Vocabulary learned from training data.
* Text converted into numeric feature vectors.



### Training

* **Loss Function**: Binary Cross-Entropy.
* **Optimization**: Gradient Descent.
* Weights and bias updated iteratively.

### 🔢 Logistic Regression Model

The predicted output is calculated using the **Sigmoid Activation Function**, which converts the linear combination of inputs into a probability score:

$$\hat{y} = \sigma(w \cdot x + b)$$

**Where:**

* **$x$**: The input text feature vector (e.g., Bag-of-Words).
* **$w$**: The learned weights that determine the importance of each feature.
* **$b$**: The bias term (intercept).
* **$\sigma$**: The Sigmoid function, defined as $\sigma(z) = \frac{1}{1 + e^{-z}}$, which maps any real-valued number into a probability between $0$ and $1$.
---

## 🌐 API Endpoints

### 1️⃣ Health Check

`GET /health`

**Response:**

```json
{
  "status": "ok"
}

```

### 2️⃣ Train Model

`POST /train`

Trains the model using labeled text data.

**Request:**

```json
{
  "data": [
    { "text": "I love this product", "label": 1 },
    { "text": "This is terrible", "label": 0 }
  ]
}

```

**Response:**

```json
{
  "message": "model trained successfully",
  "vocab": 6
}

```

> 📌 **Note:** Training updates the in-memory model state.

### 3️⃣ Predict Sentiment

`POST /predict`

**Request:**

```json
{
  "text": "I really love this"
}

```

**Response:**

```json
{
  "text": "I really love this",
  "probability": 0.94,
  "sentiment": "positive"
}

```

---

## 🚀 Getting Started

### Prerequisites

* Go 1.20+
* macOS / Linux / Windows

### Run the Server

1. Initialize the module:
```bash
go mod init sentiment-api
go mod tidy

```


2. Start the application:
```bash
go run main.go

```


3. Server starts on: `http://localhost:8080`

---

## 🧪 Example Usage (curl)

### Train

```bash
curl -X POST http://localhost:8080/train \
-H "Content-Type: application/json" \
-d '{
  "data": [
    {"text": "I love this", "label": 1},
    {"text": "This is great", "label": 1},
    {"text": "I hate this", "label": 0},
    {"text": "This is terrible", "label": 0}
  ]
}'

```

### Predict

```bash
curl -X POST http://localhost:8080/predict \
-H "Content-Type: application/json" \
-d '{ "text": "I love this product" }'

```

---

## ⚠️ Limitations (By Design)

* Model is **in-memory only** (state is lost on server restart).
* No persistence or model versioning.
* Simple Bag-of-Words (no deep embeddings like BERT).
* Intended for learning, not high-scale production use.

---

## 🧠 What This Project Demonstrates

* How ML models work internally (**no black boxes**).
* How to integrate ML with **backend APIs**.
* Separation of **training** and **inference**.
* Practical ML engineering thinking.
* Clean, minimal Go backend design.

---

## 🔮 Possible Improvements

* Persist model to disk or object storage.
* Separate trainer and serving services.
* Add evaluation metrics (accuracy, precision).
* Support larger datasets.
* Dockerize the service.

---

## 👨‍💻 Author

Built as a learning-focused ML + backend engineering project using **Go** and **Gin**.

```

```