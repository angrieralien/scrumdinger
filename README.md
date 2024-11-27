🔥 Helps teams manage their daily scrums.

## Table of Contents
- [🔔 About](#-about)
- [⭐ Inspiration](#-inspiration)
- [🚀 Goals](#-goals)
- [📦 Monorepo](#-monorepo)

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
 ┣ 📂deploy
 ┣ 📂ml
 ┣ 📂service
 ┣ 📂web
 ┣ 📂zarf

```

| directory | description                                                           |
| --------- | --------------------------------------------------------------------- |
| deploy    | contains the helm charts required to deploy Scrumdinger to Kubernetes |
| ml        | contains the code required for transcribing recorded meetings         |
| service   | contains the Go service exposing the OpenAPI REST API                 |
| web       | contains all frontend web code                                        |
| zarf      | contains all build files and scripts                                  |
