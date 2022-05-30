# or v1.0-onbuild
FROM fluent/fluentd:v1.14-debian-1

# COPY fluent.conf /fluentd/etc/

USER root
RUN gem install fluent-plugin-s3 && \
    gem install fluent-plugin-kinesis && \
    gem install fluent-plugin-rewrite-tag-filter -v 2.1.1
# USER fluent

# gem install fluent-plugin-rewrite-tag-filter -v 2.2.0