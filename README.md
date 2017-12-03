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

2. Visit your Stripe.com profile and retrieve your PublishableKey and SecretKey.

3. `go run main.go` to get the OAuth URL. Visit the URL and go through the
   authorization flow to retrieve a token. Paste the token back into the
   terminal. It will create the `sheets_api_secret_cache` file. Kill the server.

4. `./build.sh`

5. Retrieve the `PUBLISHABLE_KEY` and `SECRET_KEY` from the Stripe API. They can
   be found at https://dashboard.stripe.com/account/apikeys.

6. Launch the server with:

    docker run --rm -e PUBLISHABLE_KEY=<key> -e SECRET_KEY=<key> \
        -v /etc/ssl/certs:/etc/ssl/certs \
        -p 8586:8586 nwwa

