FROM golang:1.8

RUN mkdir -p "/home/mundipagg/"
ADD boletoapi /home/mundipagg/
RUN chmod +x /home/mundipagg/boletoapi
RUN mkdir  -p "/home/upMongo"

ENV WKHTML_MAJOR 0.12
ENV WKHTML_MINOR 4
ENV DOWNLOAD_URL "https://downloads.wkhtmltopdf.org/0.12/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz"
RUN mkdir -p "/tmp"
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl \
       xz-utils \
       fontconfig libfontconfig1 libfreetype6 \
       libpng12-0 libjpeg62-turbo \
       libssl1.0.0 libx11-6 libxext6 libxrender1 \
       xfonts-base xfonts-75dpi && \
    curl -o /tmp/wkhtmltox.tar.xz $DOWNLOAD_URL && \    
    cd /tmp && \
    tar xf wkhtmltox.tar.xz && \
    rm wkhtmltox.tar.xz && \
    
    mv ./wkhtmltox /usr/local/ && \
    
    echo 'export PATH="$PATH:/usr/local/wkhtmltox/bin"' >> ~/.bashrc && \
    apt-get purge -y curl && \
    rm -rf /var/lib/apt/lists/*

ENV PATH $PATH:/usr/local/wkhtmltox/bin

ENTRYPOINT ["/home/mundipagg/boletoapi"]
EXPOSE 3000
