FROM nginx

COPY srv/nginx.conf /etc/nginx/nginx.conf
COPY build /usr/share/nginx/website

EXPOSE 80