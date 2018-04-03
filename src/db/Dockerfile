FROM mongo

COPY mongod.conf /etc/mongod.conf

COPY entries.json /entries.json

#RUN mongoimport --db entries --collection blog --type json --file /entries.json

CMD ["mongod", "-f", "/etc/mongod.conf"]