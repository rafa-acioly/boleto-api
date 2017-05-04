# Boleto API

## FAQ

### No campo `codigoTipoTitulo` nós podemos enviar estes códigos:

1 = CHEQUE

2 = DUPLICATA-MERCANTIL

4 = DUPLICATA-SERVICO

6 = DUPLICATA-RURAL

7 = LETRA-DE-CAMBIO

12 = NOTA-PROMISSORIA

13 = NOTA-PROMISSORIA-RURAL

17 = RECIBO

19 = NOTA-DE-DEBITO

23 = DIVIDA-ATIVA-UNIAO

24 = DIVIDA-ATIVA-ESTADO

25 = DIVIDA-ATIVA-MUNICIPIO

Qual código nós devemos enviar? Estamos mandando por padrão `19`

__R.__ Este parâmetro é definido pela empresa.

### No campo `numeroCarteira` estamos mandando 17 por padrão, para eCommerce esta é a carteira ideal?

__R.__ Somente será utilizada a 17.

### No campo `codigoModalidadeTITULO` estamos mandando por padrão 1

As opções são:

1 – Carteira Simples; 4 – Vinculada 6 – Descontada; 8 – Vendor

__R.__ Continuar utilizando 1

### `dataEmissaoTITULO` e `dataVencimentoTITULO` podemos mandar um boleto que vence no mesmo dia? Data de emissão e vencimento ser a mesma?

__R.__ Pode.

### O que é `indicadorPermissaoRecebimentoParcial` Indicador de Recebimento Parcial, pra que serve? Estamos mandando sempre `N`

__R.__ Serve para o boleto ser pago parcialmente, podem ser feitos vários pagamentos até a liquidação total. (__OBS:__ não vamos usar isso)

### `textoNumeroTITULOBeneficiario` seu número, porque usar este se temos o nosso número que é de controle da empresa também?

__R.__ Campo padrão não obrigatório. (__OBS:__ não vamos usar isso)

### Por que existe o campo `textoMensagemBloquetoOcorrencia` se nós temos o trabalho de gerar o html do boleto?

__R.__ Campo padrão não obrigatório. (__OBS:__ não vamos usar isso)

### Existe alguma API onde a gente só envia as informações do boleto e vocês retornam o layout pra gente preenchido?

### Os campos que possuem Numérico(algum número) é necessário completar o número com zeros a esquerda?

### Por que quando eu envio um request passando um CPF para o pagador na API de homologação de vocês eu recebo um erro?

__R.__ Eles possuem uma relação de testes no ambiente de homologação

### Por que vem "Cliente nao informado" no campo de resposta `nomeLogradouroBeneficiario`, estamos mandando o Request corretamente?

#### Request

```xml
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:sch="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
   <soapenv:Header />
   <soapenv:Body>
      <sch:requisicao>
         <sch:numeroConvenio>1014051</sch:numeroConvenio>
         <sch:numeroCarteira>17</sch:numeroCarteira>
         <sch:numeroVariacaoCarteira>19</sch:numeroVariacaoCarteira>
         <sch:codigoModalidadeTitulo>1</sch:codigoModalidadeTitulo>
         <sch:dataEmissaoTitulo>18.04.2017</sch:dataEmissaoTitulo>
         <sch:dataVencimentoTitulo>19.04.2017</sch:dataVencimentoTitulo>
         <sch:valorOriginalTitulo>100</sch:valorOriginalTitulo>
         <sch:codigoTipoDesconto>0</sch:codigoTipoDesconto>
         <sch:codigoTipoMulta>0</sch:codigoTipoMulta>
         <sch:codigoAceiteTitulo>N</sch:codigoAceiteTitulo>
         <sch:codigoTipoTitulo>19</sch:codigoTipoTitulo>
         <sch:textoDescricaoTipoTitulo />
         <sch:indicadorPermissaoRecebimentoParcial>N</sch:indicadorPermissaoRecebimentoParcial>
         <sch:textoNumeroTituloBeneficiario />
         <sch:textoNumeroTituloCliente>00010140510000066673</sch:textoNumeroTituloCliente>
         <sch:textoMensagemBloquetoOcorrencia>Pagamento disponível até a data de vencimento</sch:textoMensagemBloquetoOcorrencia>
         <sch:codigoTipoInscricaoPagador>2</sch:codigoTipoInscricaoPagador>
         <sch:numeroInscricaoPagador>73400584000166</sch:numeroInscricaoPagador>
         <sch:nomePagador>Mundipagg Tecnologia em Pagamentos</sch:nomePagador>
         <sch:textoEnderecoPagador>R. Conde de Bonfim</sch:textoEnderecoPagador>
         <sch:numeroCepPagador>20520051</sch:numeroCepPagador>
         <sch:nomeMunicipioPagador>Rio de Janeiro</sch:nomeMunicipioPagador>
         <sch:nomeBairroPagador>Tijuca</sch:nomeBairroPagador>
         <sch:siglaUfPagador>RJ</sch:siglaUfPagador>
         <sch:codigoChaveUsuario>1</sch:codigoChaveUsuario>
         <sch:codigoTipoCanalSolicitacao>5</sch:codigoTipoCanalSolicitacao>
      </sch:requisicao>
   </soapenv:Body>
</soapenv:Envelope>
```

#### Response

```xml
<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
   <SOAP-ENV:Body>
      <ns0:resposta xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
         <ns0:siglaSistemaMensagem />
         <ns0:codigoRetornoPrograma>0</ns0:codigoRetornoPrograma>
         <ns0:nomeProgramaErro />
         <ns0:textoMensagemErro />
         <ns0:numeroPosicaoErroPrograma>0</ns0:numeroPosicaoErroPrograma>
         <ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
         <ns0:textoNumeroTituloCobrancaBb>00010140510000066673</ns0:textoNumeroTituloCobrancaBb>
         <ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
         <ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
         <ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
         <ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
         <ns0:codigoCliente>932131545</ns0:codigoCliente>
         <ns0:linhaDigitavel>00190000090101405100500066673179971340000010000</ns0:linhaDigitavel>
         <ns0:codigoBarraNumerico>00199713400000100000000001014051000006667317</ns0:codigoBarraNumerico>
         <ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
         <ns0:nomeLogradouroBeneficiario>Cliente nao informado.</ns0:nomeLogradouroBeneficiario>
         <ns0:nomeBairroBeneficiario />
         <ns0:nomeMunicipioBeneficiario />
         <ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
         <ns0:siglaUfBeneficiario />
         <ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
         <ns0:indicadorComprovacaoBeneficiario />
         <ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
      </ns0:resposta>
   </SOAP-ENV:Body>
</SOAP-ENV:Envelope>
```

### Por que quando faço uma requisição pasando o meu CPF ele diz que não encontrou na base e me retorna um erro?

O código de erro `CBRSR005` não consta no Manual.
Eu preciso colocar zeros a esquerda do CPF?

__R.__ Não vai ocorrer em produção

#### Request

```xml
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:sch="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
   <soapenv:Header />
   <soapenv:Body>
      <sch:requisicao>
         <sch:numeroConvenio>1014051</sch:numeroConvenio>
         <sch:numeroCarteira>17</sch:numeroCarteira>
         <sch:numeroVariacaoCarteira>19</sch:numeroVariacaoCarteira>
         <sch:codigoModalidadeTitulo>1</sch:codigoModalidadeTitulo>
         <sch:dataEmissaoTitulo>18.04.2017</sch:dataEmissaoTitulo>
         <sch:dataVencimentoTitulo>19.04.2017</sch:dataVencimentoTitulo>
         <sch:valorOriginalTitulo>100</sch:valorOriginalTitulo>
         <sch:codigoTipoDesconto>0</sch:codigoTipoDesconto>
         <sch:codigoTipoMulta>0</sch:codigoTipoMulta>
         <sch:codigoAceiteTitulo>N</sch:codigoAceiteTitulo>
         <sch:codigoTipoTitulo>19</sch:codigoTipoTitulo>
         <sch:textoDescricaoTipoTitulo />
         <sch:indicadorPermissaoRecebimentoParcial>N</sch:indicadorPermissaoRecebimentoParcial>
         <sch:textoNumeroTituloBeneficiario />
         <sch:textoNumeroTituloCliente>00010140510000066673</sch:textoNumeroTituloCliente>
         <sch:textoMensagemBloquetoOcorrencia>Pagamento disponível até a data de vencimento</sch:textoMensagemBloquetoOcorrencia>
         <sch:codigoTipoInscricaoPagador>1</sch:codigoTipoInscricaoPagador>
         <sch:numeroInscricaoPagador>15075791794</sch:numeroInscricaoPagador>
         <sch:nomePagador>Mundipagg Tecnologia em Pagamentos</sch:nomePagador>
         <sch:textoEnderecoPagador>R. Conde de Bonfim</sch:textoEnderecoPagador>
         <sch:numeroCepPagador>20520051</sch:numeroCepPagador>
         <sch:nomeMunicipioPagador>Rio de Janeiro</sch:nomeMunicipioPagador>
         <sch:nomeBairroPagador>Tijuca</sch:nomeBairroPagador>
         <sch:siglaUfPagador>RJ</sch:siglaUfPagador>
         <sch:codigoChaveUsuario>1</sch:codigoChaveUsuario>
         <sch:codigoTipoCanalSolicitacao>5</sch:codigoTipoCanalSolicitacao>
      </sch:requisicao>
   </soapenv:Body>
</soapenv:Envelope>
```

#### Response

```xml
<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
   <SOAP-ENV:Body>
      <ns0:resposta xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
         <ns0:siglaSistemaMensagem />
         <ns0:codigoRetornoPrograma>5</ns0:codigoRetornoPrograma>
         <ns0:nomeProgramaErro>CBRSR005</ns0:nomeProgramaErro>
         <ns0:textoMensagemErro>?CPF do pagador nao encontrado na base.</ns0:textoMensagemErro>
         <ns0:numeroPosicaoErroPrograma>5</ns0:numeroPosicaoErroPrograma>
         <ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
         <ns0:textoNumeroTituloCobrancaBb />
         <ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
         <ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
         <ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
         <ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
         <ns0:codigoCliente>932131545</ns0:codigoCliente>
         <ns0:linhaDigitavel />
         <ns0:codigoBarraNumerico />
         <ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
         <ns0:nomeLogradouroBeneficiario />
         <ns0:nomeBairroBeneficiario />
         <ns0:nomeMunicipioBeneficiario />
         <ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
         <ns0:siglaUfBeneficiario />
         <ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
         <ns0:indicadorComprovacaoBeneficiario />
         <ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
      </ns0:resposta>
   </SOAP-ENV:Body>
</SOAP-ENV:Envelope>
```