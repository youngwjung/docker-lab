FROM alpine
RUN mkdir /app
COPY --chmod=400 password.txt /app
RUN rm $(which cat)
USER 100:100
CMD ["sleep", "3600"]