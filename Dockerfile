FROM scratch

ENV PORT 8000
EXPOSE 443

COPY advent /
CMD ["/advent"]
