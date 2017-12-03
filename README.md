NW WoodWorkers Association
--------------------------

This repo contains the source for the NWWA website, which can be found at
https://nwwoodworkers.org.

Building & Running
==================

1. Visit the [Google Sheets Go
   Quickstart](https://developers.google.com/sheets/api/quickstart/go) and
   follow the instructions to retrieve a `client_secret.json` file with your
   credentials which enable the OAuth process.

2. `go run main.go` to get the OAuth URL. Visit the URL and go through the
   authorization flow to retrieve a token. Paste the token back into the
   terminal. It will create the `sheets_api_secret_cache` file. Kill the server.

3. `./build.sh`

4. Either launch the server with:

   ```
   docker run --rm \ -v /etc/ssl/certs:/etc/ssl/certs \ -p 8586:8586 nwwa
   ```

   or create the container and set it to launch automatically, then start.

   ```
   docker create \
       --name nwwa-website \
       -p 8586:8586 \
       -v /etc/ssl/certs:/etc/ssl/certs \
       --expose 8586 \
       --restart always \
       nwwa:latest
   docker start nwwa-website
   ```

