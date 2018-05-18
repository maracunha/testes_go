#language: pt

Funcionalidade: Login

    // para quem eu estou fazendo
    // o que eu vou fazer 
    // qual o valor de negocio que isso agrega

    Para que eu possa ter acesso as minhas tarefas
    Sendo um usuário
    Posso me autenticar com os meus dados previamente cadastrados

    Contexto: Formulário de login
        Dado que eu acessei a página principal

   @success
    Cenario: Login do usuário

        Quando faço login com "eu@papito.io" e "123456"
        Então sou autenticado com sucesso
    
     Cenario: Senha incorreta

        Quando faço login com "eu@papito.io" e "d4455454sdfs"
        Então deve ver a seguinte mensagem "Senha inválida."
    
    Cenario: Email incorreta

        Quando faço login com "eu@papit11o.io" e "123456"
        Então deve ver a seguinte mensagem "Usuário não cadastrado."