FROM traefik:2.9.6
WORKDIR /
ADD defaultcsp-plugin.tar.gz /
RUN ls -R /plugins-local/
COPY integration/traefik.yml /traefik.yml
COPY integration/runtime-config.yml /runtime-config.yml