# Use the official Redis image from the Docker Hub
FROM redis:latest

# Define the volume for Redis data
VOLUME /redis-data

EXPOSE 6379

# Define the default command to run Redis
CMD ["redis-server", "--appendonly", "yes"]