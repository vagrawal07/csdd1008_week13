# csdd1008_week13

# Instructor

| Instructor  | Maziar Sojoudian  |

# Go Application with Docker

# Docker Hub Link
https://hub.docker.com/repository/docker/natishkumar02/goapiweek13/general

# git link

 https://github.com/vagrawal07/csdd1008_week13

Go Application with Docker

Overview
This repository contains a Go application packaged with Docker. The application demonstrates basic usage of Docker with a Go application, including building and running the application in a container.
## Development

### Running the Go Application

Instructions:
1. Open the Project:
Launch your preferred code editor.
Open the project directory containing main.go.
2. Run the Application:
Open the terminal within the code editor.
Execute the following command to run the Go application
3. Access the Application:
Open your web browser.
Navigate to http://localhost:8080 to view the application.

Docker Image Buildup

Build the Image:
docker build -t natishkumar02/goapiweek13 .

Run the Container:
docker run -d -p 8080:8080 natishkumar02/goapiweek13

Push to Docker Hub:
docker push natishkumar02/goapiweek13

Azure VM Creation and Configuration

Create a Debian-Based VM
1. Log in to Azure Portal.
2. Go to “Virtual machines” and click “Add”.
3. Enter VM details:
• Choose subscription, resource group, and region.
• Select “Debian” image and VM size.
• Set authentication to “Password” and enter credentials.
4. Configure disks, networking, and ports
 Allow: Port 22 (SSH),Port 80 (HTTP),Port 443 (HTTPS)
5. Review and create the VM.

Swarm Deployment Steps

Initialize Swarm Mode : sudo docker swarm init
Create and Deploy Service :
sudo docker service create --name myservice --publish 80:8080 --replicas VaibhawAggarwal/go-hostname-app:latest
Monitor Deployment:
Visit  http://20.77.32.133:80

Conclusion

Verify Directory: Ensure you are in the correct directory containing the Dockerfile.
Check for Dockerfile: Confirm the presence and correct naming of the Dockerfile using dir.
Create/Rename Dockerfile: If missing or misnamed, create or rename the file to Dockerfile.
Ensure Docker is Running: Verify that Docker Desktop is active.
Retry Build Command: Execute docker build -t natishkumar02/goapiweek13 . to build the Docker image. 


# Team Members

| Name           | Student ID  |
|----------------|-------------|
| Disha Sharma   | 500234080   |
| Sachin Dhamija | 500223167   |
| Mandeep Kour   | 500235136   |
| Vaibhav Agrawal| 500231194   |
| Natish Kumar   | 500235460   |
| Veerpal Kaur   | 500224580   |
