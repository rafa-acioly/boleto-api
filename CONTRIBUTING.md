# Contributing

Este repositório contém informações de como será a política do time One em relação as branches e mensagens de commit, por favor leia antes de começar a utilizar nossos repositórios.

## Índice

[Como eu posso contribuir?](#como-eu-posso-contribuir)
  * [Reportando Bugs](#reportando-bugs)
  * [Fluxo de Desenvolvimento](#fluxo-de-desenvolvimento)
  * [Fluxo Git](#fluxo-git)
  * [Pull Requests](#pull-requests)
  * [Pull Request Labels](#pull-request-labels)

[Styleguides](#styleguides)
  * [Git Commit Messages](#git-commit-messages)

## Como eu posso contribuir?

### Reportando Bugs

Esta seção mostra como reportar um bug para o BoletoConciliationService. Seguir estes _guidelines_ ajuda o time a entender o seu _report_ :pencil: e reproduzir o comportamento :computer: :computer:.

Antes de criar um bug report verifique [esta lista](#antes-de-reportar-um-bug). Quando for criar um novo _report_ de bug, inclua [a maior quantidade de detalhes possíveis](#como-eu-crio-um-bom-report-de-bug) por favor. Preencha [o template](#template-for-submitting-bug-reports) para facilitar na resolução dos bugs.

#### Antes de reportar um Bug

* Verifique se ele já existe no [quadro de tarefas](https://mundipagg.visualstudio.com/MundiPagg/_backlogs/board/Backlog%20items).

#### Como eu crio um (Bom) report de Bug?

Bugs são rastreados como PBI's no quadro de tarefas e basicamente devem possuir uma descrição e um critério de aceite. Opcionalmente, pode-se anexar arquivos que possam auxiliar no entendimento do bug.

Explique o bug e inclua detalhes adicionais para ajudar na reprodução do problema:

* **Use um título claro e descritivo** no PBI para a identificação do problema.
* **Descreva os passos exatos que reproduzem o problema** com o máximo de detalhes possível. 
* **Se possível, forneça exemplos específicos para reproduzir os passos**. Inclua links ou arquivos que mostrem onde o problema ocorre.
* **Descreva o comportamento observado após seguir os passos** e aponte exatamente qual o problema nesse comportamento.
* **Explique qual era o comportamento esperado e o por quê.**

* **Se possível, inclua evidências do erro no SEQ**

### Pull Requests

* Siga os styleguides de código.
* Verifique se todos os testes estão passando.
* Siga o padrão das mensagens de commit.
* No **Título** descreva sucintamente o que o pull request representa, seguido de um dos labels adequados, descritos [aqui](#pull-request-labels).
* Na **Descrição** documente a alteração de acordo com o styleguide:
 
  - `O que foi feito?`
  - `Por que foi feito?`
  - `Como foi feito?`
  
- Coloque como **Reviewers** o time `MundiPagg Team`, assim todos os membros do time serão selecionados.

**Exemplo:**

```
Title:      [#bug] Correção de erros de validação no login

Description:  
        O que foi feito?
          
        - Corrige login.
        - Implementa validação de campos no login.

        Por que foi feito?
          
        - Quando se tentava logar com informações erradas a página estava 
        recarregando desnecessariamente.

        Como foi feito?

        - No momento do login, no estado não autorizado, foi retirado o
         URIEncoded (responsável pelo problema).
```

### Pull Request Labels

Essa seção lista as labels (etiquetas) usadas na mensagem de pull request.
Cada label representará uma seção no changelog, seguida da mensagem commit dos pull requests que descreverá a alteração.  

| Label |  Description |
| --- | --- |
| `enhancement` |  Feature nova ou melhorias |
| `bug` | Correção de bugs |
|`documentation`| Adição/alteração de documentação |
| `fire` |  Remoção de código |

## Styleguides

### Git Commit Messages

* Use o presente imperativo ("Adiciona feature" não "Adicionada feature"). Dica: pense na ação que o commit está realizando
* Limitar a primeira linha a 72 caracteres ou menos
* Considere iniciar a mensagem de commit com um emoji que se aplique:

| Código                | Emoji               | Descrição                                       |
|-----------------------|---------------------|-------------------------------------------------|
| `:art:`               | :art:               | when improving the format/structure of the code |
| `:racehorse:`         | :racehorse:         | when improving performance                      |
| `:non-potable_water:` | :non-potable_water: | when plugging memory leaks                      |
| `:memo:`              | :memo:              | when writing docs                               |
| `:checkered_flag:`    | :checkered_flag:    | when fixing something on windows                |
| `:bug:`               | :bug:               | when fixing a bug                               |
| `:fire:`              | :fire:              | when removing code or files                     |
| `:green_heart:`       | :green_heart:       | when fixing CI build                            |
| `:white_check_mark:`  | :white_check_mark:  | when adding tests                               |
| `:lock:`              | :lock:              | when dealing with security                      |
| `:arrow_up:`          | :arrow_up:          | when upgrading dependencies                     |
| `:arrow_down:`        | :arrow_down:        | when downgrading dependencies                   |
| `:shirt:`             | :shirt:             | when removing linter warnings                   |
| `:bulb:`              | :bulb:              | new idea                                        |
| `:construction:`      | :construction:      | work in progress                                |
| `:heavy_plus_sign:`   | :heavy_plus_sign:   | when adding feature                             |
| `:heavy_minus_sign:`  | :heavy_minus_sign:  | when removing feature                           |
| `:speaker:`           | :speaker:           | when adding logging                             |
| `:mute:`              | :mute:              | when removing logging                           |
| `:facepunch:`         | :facepunch:         | when resolving conflicts                        |
| `:wrench:`            | :wrench:            | when changing configuration files               |

[:arrow_left: Voltar](README.md)
    

# Como lidar com o problema de import path em Projetos Go

A solução para este problema está [aqui](http://code.openark.org/blog/development/forking-golang-repositories-on-github-and-managing-the-import-path)