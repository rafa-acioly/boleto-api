FROM ubuntu

RUN mkdir -p "/home/mundipagg/boletoapi"
ADD boletoapi /home/mundipagg/boletoapi
RUN chmod 777 /home/mundipagg/boletoapi/boletoapi
ENTRYPOINT ["/home/mundipagg/boletoapi"]
EXPOSE 3000
