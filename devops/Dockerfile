FROM ubuntu

RUN mkdir -p "/home/mundipagg/"
ADD boletoapi /home/mundipagg/
RUN chmod +x /home/mundipagg/boletoapi
RUN mkdir  -p "/home/upMongo"
ENTRYPOINT ["/home/mundipagg/boletoapi"]
EXPOSE 3000
