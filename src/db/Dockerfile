FROM mongo

COPY src/server/mongod.conf /etc/mongod.conf

COPY src/server/entries.json /entries.json

RUN mongoimport --db entries --collection blog --type json --file /entries.json

CMD ["mongod", "-f", "/etc/mongod.conf"]