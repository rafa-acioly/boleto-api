# IMPORTANTE: COMO USAR AS ISSUES DO BITBUCKET DO BOLETOONLINE

* As issues do Bitbucket DEVERÃO SER USADAS APENAS para reportar bugs e para novas FEATURES BEM DETALHADAS. Qualquer coisa além disso deverá ser discutido pelo canal do Slack.

# Como enviar um patch para uma nova feature

1. Antes de mandar qualquer código ou mudança grande para o projeto, poste sua ideia no grupo do Slack do BoletoOnline. Tente demonstrar porque sua feature é importante, ouça o feedback do time de boleto.
2. Caso sua feature seja aprovada pelo time do projeto utilize os seguintes procedimentos:

    a. Faça o [fork](https://confluence.atlassian.com/bitbucket/forking-a-repository-221449527.html) da API de Boleto no Bitbucket 
    
    b. Crie uma branch (git checkout -b sua_feature_incrivel)
    
    c. Envie para sua branch (git push origin sua_feature_incrivel)
    
    d. Inicie um [Pull Request no Bitbucket](https://confluence.atlassian.com/bitbucket/create-a-pull-request-774243413.html)
    

# Como lidar com o problema de import path em Projetos Go

A solução para este problema está [aqui](http://code.openark.org/blog/development/forking-golang-repositories-on-github-and-managing-the-import-path)