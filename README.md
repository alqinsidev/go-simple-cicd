## Go CI/CD Example Using GitHub Actions
This repository provides a straightforward example of implementing CI/CD for a Go application using **GitHub Actions**.

## Application Overview
The application is a simple Go program that includes a function to output a basic string, along with test code to validate this functionality. A Dockerfile is also provided to build an image of the app, which is then pushed to Docker Hub.

## Prerequisite
- Active Docker Hub Repo

## CI/CD Process
### CI Workflow

The Continuous Integration (CI) workflow is triggered by any Pull Request to the `main` branch. This workflow is structured to ensure that new changes meet quality standards before merging. It includes the following steps:

1. **Checkout Code**: Uses GitHub’s `actions/checkout@v4` to pull the latest code from the repository.
2. **Set Up Go Environment**: Configures the Go environment using `actions/setup-go@v5` with version `1.21.0`.
3. **Install Dependencies**: Downloads all required dependencies by running `go mod download`.
4. **Run Tests**: Executes tests across all packages using `go test ./...` to validate the code.

This workflow ensures that every pull request passes code validation and testing standards before merging.

---

### CD Workflow

The Continuous Deployment (CD) workflow is triggered by any push to the `main` branch, such as when a pull request is merged. The deployment workflow includes these steps:

1. **Checkout Code**: Uses `actions/checkout@v4` to pull the latest code.
2. **Set Up Go Environment**: Configures the Go environment with version `1.21.0`.
3. **Log in to Docker Hub**: Authenticates with Docker Hub using the `docker/login-action@v2` action, with credentials provided through GitHub Secrets (`DOCKER_USERNAME` and `DOCKER_PASSWORD`).
4. **Set Version Tag**: Generates a short SHA tag from the latest commit using `git rev-parse --short HEAD`, which is stored in the `sha_short` environment variable.
5. **Build Docker Image**: Builds a Docker image with the tag `go-simple-cicd:${sha_short}`.
6. **Push Docker Image to Docker Hub**: Uploads the tagged Docker image to Docker Hub under the specified repository.

This CD pipeline builds and pushes a Docker image based on the latest code from `main`, allowing for an updated image to be deployed to the target environment.
