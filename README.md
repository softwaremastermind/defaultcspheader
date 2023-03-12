# traefik-default-csp-header-plugin
This is a traefik middleware plugin that adds a default CSP header on every upstream response without a CSP header.

This plugin implements the following behavior:
- If any upstream traefik middleware or service set a `Content-Security-Policy` header in the response, this plugin does nothing.
- Otherwise, this plugin set a default `Content-Security-Policy` header. You can provide the default `Content-Security-Policy` in the `traefik` configuration.

## Setup
To configure this plugin you should add its static and dynamic configuration to the Traefik dynamic configuration as explained [here](https://docs.traefik.io/getting-started/configuration-overview/#the-dynamic-configuration).

You find an example for the [static configuration](integration/traefik.yml) and the [dynamic configuration](integration/runtime-config.yml) in the integration test.

## Configuration
You must set the plugin configuration in the runtime configuration of traefik.

| Option  | Required  | Description  |
|---|---|---|
| `defaultCSPHeader`  | true  | The `Content-Security-Policy` header value that is set on every http response, that does not contains a `Content-Security-Policy` header  |

