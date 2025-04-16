# BloomFilter.io 🚀

BloomFilter.io is a high-performance, privacy-aware username availability checker built using **Go**, **React**, and **Bloom Filters**. It can check millions of usernames in milliseconds, making it ideal for scalable, low-latency applications.

---

## Table of Contents
- [What is a Bloom Filter?](#what-is-a-bloom-filter)
- [Theoretical Foundation](#theoretical-foundation)
- [Architecture Overview](#architecture-overview)
- [Features](#features)
- [Folder Structure](#folder-structure)
- [Setup Instructions](#setup-instructions)
  - [1. Server (Go)](#1-setting-up-the-server-go)
  - [2. Client (React)](#2-setting-up-the-client-react)
  - [3. Data Preparation](#3-data-preparation)

---

# What is a Bloom Filter?

A **Bloom Filter** is a space-efficient probabilistic data structure used to test whether an element is a member of a set. It can result in **false positives** (claiming an element is in the set when it isn't), but never false negatives.

This makes Bloom Filters especially useful in scenarios where fast membership queries are needed, and occasional false positives are acceptable.

<div style="display: flex; justify-content: space-between;">
    <img src="assets/success.png" alt="alt text" width="500" height="500" style="margin-right: 10px;">
    <img src="assets/failure.png" alt="alt text" width="500" height="500">
</div>

---

# Theoretical Foundation

## How a Bloom Filter Works
A Bloom Filter is an array of bits (initially all set to 0), and uses **k** different hash functions.

### Insertion:
To insert an element:
1. Hash it with all **k** hash functions.
2. Set the bits at each resulting index to 1.

### Query:
To test if an element is in the set:
1. Hash it with all **k** hash functions.
2. If **any** of the bits at the hashed positions are 0 → the element is **definitely not** in the set.
3. If **all** bits are 1 → the element is **probably** in the set (possible false positive).


### Advantages:
- Space and time efficient.
- Fast lookups, regardless of data size.

### Limitations:
- Does not support deletion (unless using Counting Bloom Filters).
- Cannot retrieve actual values, only membership status.
- False positives possible (no false negatives).

---

## Properties of Bloom Filters

### Key Characteristics:
- **Space-efficient**: Requires minimal memory to store large datasets.
- **Extremely fast lookup**: Provides quick membership checks.
- **Supports millions of entries**: Scales efficiently with large datasets.
- **No deletion**: Does not support removing elements unless using specialized variants like Counting Bloom Filters.

### Common Use Cases:
- **Username/Email Availability**: Quickly check if a username or email is already taken.
- **Web Caching**: Determine if a resource is cached without retrieving it.
- **Network Security**: Filter malicious URLs or detect spam efficiently.
- **Database Query Optimization**: Reduce unnecessary database lookups by pre-checking membership.
- **Distributed Systems**: Synchronize data across nodes with minimal overhead.

---

## Features

- ✅ Check username availability with millisecond response times
- ✅ Debounced input to reduce server load
- ✅ Bloom filter backed by a text file and initialized at server startup
- ✅ Real-time response logging (including timing in ms)
- ✅ CORS-enabled backend for safe frontend/backend separation

---

# Setup Instructions

## 1. Data Source
### **Download the dataset from [**Kaggle - Reddit Usernames**](https://www.kaggle.com/datasets/colinmorris/reddit-usernames?resource=download). After downloading, unzip the file and place the extracted `users.csv` file inside the `data` folder.**

## 2. Setting Up the Server (Go)

### Steps:
1. Clone the repository:
    ```bash
    git clone https://github.com/Prayag2003/bloom-filter
    ```
2. Navigate to the server directory:
    ```bash
    cd bloom-filter/server
    ```
3. Install dependencies:
    ```bash
    go mod tidy
    ```
4. Start the server:
    ```bash
    go run main.go
    ```

The server will be running at: `http://localhost:8080`

---

## 3. Setting Up the Client (React)

### Steps:
1. Navigate to the client directory:
    ```bash
    cd ../client
    ```
2. Install dependencies:
    ```bash
    pnpm install
    ```
3. Start the development server:
    ```bash
    pnpm run dev
    ```

The application will be available at: `http://localhost:5173`

