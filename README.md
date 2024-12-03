<p align="center">
  🔥 Helps teams manage their daily scrums.
</p>

<p align="center">
  <img src="https://github.com/user-attachments/assets/34eb52fc-701f-4a1e-8f18-2b1d788ae323"  width="20%" />
  <img src="https://github.com/user-attachments/assets/a1ad5868-7b72-48c0-ac5c-ece4a8de2d9e" width="20%"/>
  <img src="https://github.com/user-attachments/assets/18d54f6d-7507-4fa2-86fb-44a9ea2d8158" width="20%" />
</p>
  
## Table of Contents
- [📼 Demo](#-demo)
- [🔔 About](#-about)
- [⭐ Inspiration](#-inspiration)
- [🚀 Goals](#-goals)
- [📦 Monorepo](#-monorepo)
  - [🚧 Prerequisites](#-prerequisites)
    - [Docker](#docker)
    - [Go](#go)
    - [Node](#node)
- [🔬 Enironments](#-environments)
  - [Development](#development)
 
# 📼 Demo
[scrumdinger_demo.webm](https://github.com/user-attachments/assets/6b71085c-b8da-4ea9-a622-82dd212a181e)

# 🔔 About

Many software engineering teams use daily meetings, known as scrums, to plan their work for the day. Scrums are short meetings where each attendee discusses what they accomplished yesterday, what they are working on today, and any obstacles that might impact their work.

# ⭐ Inspiration

Scrumdinger originally is an iOS [tutorial](https://developer.apple.com/tutorials/app-dev-training/getting-started-with-scrumdinger) used to teach people the basics of SwiftUI.

# 🚀 Goals

* Create a **web service** using OpenAPI Specification.
* Create a responsive **web app** that can be used on desktop and mobile.
* Using a machine learning model, provide a way for users to transcribe recorded meetings.
* Provide a Helm chart for easy deployment to Kubernetes.

# 📦 Monorepo

```text

📦 Scrumdinger
 ┣ 📂api
 ┃ ┣ 📂frontends // All frontend code
 ┃ ┃ ┣ 📂...
 ┃ ┣ 📂services // All web service code
 ┃ ┃ ┣ 📂...
 ┣ 📂app // web controllers and routes for web services
 ┃ ┃ ┣ 📂...
 ┣ 📂business // Business logic 
 ┃ ┃ ┣ 📂...
 ┣ 📂foundation // Common Go code for web services
 ┃ ┃ ┣ 📂...
 ┣ 📂zarf // Deployment files

```

## 🚧 Prerequisites

### Docker

I recommend installing docker in [rootless mode](https://docs.docker.com/engine/security/rootless/). Also install [docker compose](https://docs.docker.com/compose/install/).

### Go

[Download](https://go.dev/dl/) and follow these install [instructions](https://go.dev/doc/install#install).

### Node

Download [nvm](https://nodejs.org/en/download/package-manager) and follow the instructions to install.


# 🔬 Environments

## Development

Run `make build` and `make compose-up`

Navigate to the [web app](api/frontends/scrumdinger) and run `npm install && npm run dev`. The test email is `admin@example.com` and the password is `gophers`.
