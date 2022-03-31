# phones-api

First, let’s download the Docker repository. Run the following:

`sudo apt-get update`

Followed by:
`sudo apt-get install 
    apt-transport-https 
    ca-certificates 
    curl 
    software-properties-common`

Next run:
`curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -`

To verify you have the repository run the following:
`sudo apt-key fingerprint 0EBFCD88`

And you should get something like this:
`pub   4096R/0EBFCD88 2017-02-22 Key fingerprint = 9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid                  Docker Release (CE deb) <docker@docker.com>
sub   4096R/F273FCD8 2017-02-22`

Now to install Docker you just need to do the following:
`sudo apt-get update && sudo apt-get install docker-ce`

Now to get Docker Compose run:
`sudo curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose`

To test if your installation was setup correctly, run:
`docker-compose --version`

And you should get something similar to:
`docker-compose version 1.22.0, build 1719ceb`

If you want to check mysql databases and tables, install the mysql client core:
`sudo apt install mysql-client-core-8.0`

Then, go to your repo folder:
`cd <patch/to/golang-rump-up>`

and run:
`docker-compose up -d`

You can choose two types of database drives: `mysql` or `orm`. Just change the `DB_IMPL` field in the environment file to `mysql` or `orm` and follow steps bellow:

`docker-compose build`

`docker-compose up`

If you encounter a non-existent database error, follow the steps below:

`docker-compose down`

`docker volume ls`

`docker volume rm -f golang-rump-up_my-db`

If the container is running, remove it and then repeat the command

`docker volume rm -f golang-rump-up_my-db`

After that, run the command below again:

`docker-compose up -d`
