FROM traefik:2.9.6
WORKDIR /
RUN mkdir -p /plugins-local/src/github.com/softwaremastermind/defaultcspheader

COPY defaultcsp.go /plugins-local/src/github.com/softwaremastermind/defaultcspheader/defaultcsp.go
COPY .traefik.yml /plugins-local/src/github.com/softwaremastermind/defaultcspheader/.traefik.yml
# COPY . /plugins-local/src/github.com/softwaremastermind/defaultcspheader
RUN ls -R /plugins-local/
COPY integration/traefik.yml /traefik.yml
COPY integration/runtime-config.yml /runtime-config.yml