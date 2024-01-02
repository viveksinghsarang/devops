Clone the git repo on the server.
git clone https://github.com/viveksinghsarang/devops

Copy the contents to /app directory.
cp devops /app

Then run the below cmd to up the docker containers:-
docker compose up --build

To check the service is up, test it from postman:-
Post http://ip:3000/process

{
  "data":[3, 4, 0, 7, 9, 2, 1, 5]
}
