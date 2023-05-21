# Campaign Demo

This is a demo application for managing campaigns and products.

## Overview

Campaign Demo is a command-line application that allows you to create, manage, and monitor campaigns and products. It provides functionality for creating products, placing orders, creating campaigns, and retrieving information about products and campaigns. The application is written in Go and can be built and run locally or inside a Docker container.

## Prerequisites

- Go 1.16 or later
- Docker (optional)

## Getting Started

To get started with the Campaign Demo application, follow these steps:

1. Clone the repository:
    ```shell
    git clone https://github.com/eylmzer/campaingdemo.git
    ```

2. Navigate to the project directory:
    ```shell
    cd campaingdemo
    ```
    
3. Build the application:
    ```shell
    make build
    ```
    
4. Run the application:
    ```shell
    ./campaingdemo
    ```
5. Run the tests:
    ```shell
    make test
    ```
    
## Docker Support
Alternatively, you can run the application using Docker. Docker allows you to isolate the application's dependencies and run it in a containerized environment.

To use Docker, follow these steps:

1. Build the Docker image:
    ```shell
    make docker-build
    ```
    
2. Run the Docker container:
    ```shell
    make docker-run
    ```

3. Stop and clean up the Docker container and image:
    ```shell
    make docker-clean
    ```
