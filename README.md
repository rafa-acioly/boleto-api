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

### No campo `numeroCarteira` estamos mandando 17 por padrão, para eCommerce esta é a carteira ideal?

### No campo `codigoModalidadeTITULO` estamos mandando por padrão 1

As opções são:

1 – Carteira Simples; 4 – Vinculada 6 – Descontada; 8 – Vendor

### `dataEmissaoTITULO` e `dataVencimentoTITULO` podemos mandar um boleto que vence no mesmo dia? Data de emissão e vencimento ser a mesma?

### O que é `indicadorPermissaoRecebimentoParcial` Indicador de Recebimento Parcial, pra que serve? Estamos mandando sempre `N`

### `textoNumeroTITULOBeneficiario` seu número, porque usar este se temos o nosso número que é de controle da empresa também?

### Por que existe o campo `textoMensagemBloquetoOcorrencia` se nós temos o trabalho de gerar o html do boleto?

### Existe alguma API onde a gente só envia as informações do boleto e vocês retornam o layout pra gente preenchido?

### Os campos que possuem Numérico(algum número) é necessário completar o número com zeros a esquerda?

### `codigoChaveUsuario` Campo para identificar usuário/sistema interno que solicitou o registro – controle empresa – Default: 1 pra que serve? Vocês tem uma API de consulta para verificar as requisições feitas?