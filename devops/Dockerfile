FROM golang:1.8

RUN mkdir -p "/home/mundipagg/"
ADD boletoapi /home/mundipagg/
ADD wkhtmltox.tar.xz /usr/local/
RUN chmod +x /home/mundipagg/boletoapi
RUN mkdir  -p "/home/upMongo"
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl \
       xz-utils \
       fontconfig libfontconfig1 libfreetype6 \
       libpng12-0 libjpeg62-turbo \
       libssl1.0.0 libx11-6 libxext6 libxrender1 \
       xfonts-base xfonts-75dpi && \        
    echo 'export PATH="$PATH:/usr/local/wkhtmltox/bin"' >> ~/.bashrc && \
    apt-get purge -y curl && \
    rm -rf /var/lib/apt/lists/*
ENV PATH $PATH:/usr/local/wkhtmltox/bin

ENTRYPOINT ["/home/mundipagg/boletoapi"]
EXPOSE 3000
