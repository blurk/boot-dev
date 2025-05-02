# This is a comment

# Use a lightweight debian os
# as the base image
FROM debian:stable-slim

# ADD command would also work here, but COPY is the better choice for us because we don't need the extra functionality that ADD offers.
# COPY source destination
COPY ./go-app/go-app /bin/go-app

ENV PORT=8080

CMD ["/bin/go-app"]