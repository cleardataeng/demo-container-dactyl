FROM scratch
ADD dactyl /
EXPOSE 8080
USER 1000
CMD ["/dactyl"]
