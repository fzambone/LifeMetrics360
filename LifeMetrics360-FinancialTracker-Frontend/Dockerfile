# Use the official Node.js image as a base image
FROM node:20.9.0-alpine

# Set the current working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source code into the container
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000

# Start the React application
CMD ["npm", "start"]