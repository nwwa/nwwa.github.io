FROM scratch

COPY client_secret.json /
COPY sheets_api_secret_cache /
COPY templates templates
COPY static static
ADD main /

EXPOSE 8586

CMD ["/main"]
