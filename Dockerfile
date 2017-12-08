FROM nginx

# Website
COPY build /usr/share/nginx/website
RUN mkdir -p /usr/share/nginx/data/db

# Server environment
COPY srv/nginx.conf /etc/nginx/nginx.conf

# Database environment
COPY srv/mongodb.conf /etc/mongod.conf

# Install mongodb
RUN apt-get update
RUN apt-get install -y mongodb

# Run mongodb
CMD mongod --config /etc/mongod.conf

EXPOSE 80