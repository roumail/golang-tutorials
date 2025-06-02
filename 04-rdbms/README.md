# Accessing a relational database

This tutorial is following the content posted [here](https://go.dev/doc/tutorial/database-access). The reference material for further dep dive is found [here](https://go.dev/doc/database/).

After creating the directory `data-access`, we run the `go mod init example/data-access` command to create the `go.mod` file.

The next part mostly focused on getting an up and running `mysql` setup, for which I used docker-compose. I added data in a script folder called sql-scripts that I mount into the mysql service. The database is launched using `docker-compose up -d mysql`. To connect to the db in a cli session, I can use `docker-compose run --rm mysql-client` and then drop into `mysql -h mysql -u root -p`, where the password for root is the highly insecure `rootpassword`. Here we're able to discover the databased hosted under the service name `mysql` because of docker's network dns resolution.

Once we're in the shell, we can use `use recordings;` followed by `select * from album;` to verify that everything went well.
