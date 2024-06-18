# Use Node.js image with NPM and Node.js pre-installed
FROM node:14

# Set working directory inside the Docker image
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy all files from the frontend directory to the Docker image
COPY . .

# Build your application (if necessary)
# RUN npm run build

# Install a simple HTTP server
RUN npm install -g http-server

# Expose the port your app runs on
EXPOSE 8080

# Command to run your app using http-server
CMD ["http-server", "index.html"]
