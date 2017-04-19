FROM ubuntu

RUN mkdir -p "/home/mundipagg/"
ADD boletoapi /home/mundipagg/
RUN chmod +x /home/mundipagg/boletoapi
RUN ls -la /home/mundipagg/
ENTRYPOINT ["/home/mundipagg/boletoapi"]
EXPOSE 3000
