FROM mockserver/mockserver:5.14.0
COPY integration//mockserver.json /mockserver.json
ENV MOCKSERVER_WATCH_INITIALIZATION_JSON=true
ENV MOCKSERVER_INITIALIZATION_JSON_PATH=/mockserver.json