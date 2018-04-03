FROM nginx

COPY nginx.conf /etc/nginx/nginx.conf
COPY entrylocation.conf /etc/nginx/entrylocation.conf
COPY build/ /usr/share/nginx/html