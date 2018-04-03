FROM mongo

COPY entries.json /entries.json

# The "db" host is specified in docker-compose.yml and has to be differentiated 
# with mongodbs definition of databases (short db)
CMD mongoimport --host db --db entries --collection blog --type json --file /entries.json
