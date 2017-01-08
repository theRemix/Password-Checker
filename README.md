# Password-Checker
SSH 2017 Hackathon Project

## Running this project

```sh
yarn install
npm start
```

## Release

```sh
npm run release
npm run docker-compile-bin
docker build -t theremix/password-checker:latest .

```
### Running release

```sh
docker run -it --rm --name password-checker -p 9090:80 theremix/password-checker
```
will run the server listening on port 9090
