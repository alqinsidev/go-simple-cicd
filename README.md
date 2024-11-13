## Go CI/CD Example Using GitHub Actions
This repository provides a straightforward example of implementing CI/CD for a Go application using **GitHub Actions**.

## Application Overview
The application is a simple Go program that includes a function to output a basic string, along with test code to validate this functionality. A Dockerfile is also provided to build an image of the app, which is then pushed to Docker Hub.

## Prerequisite
- Active Docker Hub Repo

## CI/CD Workflow
The CI/CD process is triggered by any **Pull Request** made to the `main` branch. The workflow includes the following steps:
   - Run tests to validate the code
   - Log in to Docker Hub
   - Create a Tag from GitHub short SHA from last commit (Pull Request)
   - Build a Docker image for the application
   - Push the Docker image to Docker Hub
