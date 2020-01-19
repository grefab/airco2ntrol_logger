FROM envoyproxy/envoy:latest
RUN apt-get update && \
    apt-get install gettext -y
COPY envoy.yaml /tmpl/envoy.yaml
CMD cat /tmpl/envoy.yaml | envsubst \$STORAGE_HOST,\$STORAGE_PORT > /etc/envoy.yaml && /usr/local/bin/envoy -c /etc/envoy.yaml
