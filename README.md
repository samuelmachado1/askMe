AskMe
AskMe é uma aplicação web para perguntas e respostas desenvolvida em Go.

Pré-requisitos
Go: Certifique-se de ter o Go instalado em sua máquina. Você pode baixá-lo em https://go.dev/.
Lima: Ferramenta para provisionamento de máquinas virtuais Linux. Baixe em https://lima-vm.io/.
Nerdctl: Cliente de linha de comando para o containerd, usado pelo Lima.
Instalação e Execução
Clone o repositório:
git clone https://github.com/seu-usuario/askMe.git
   cd askMe
Inicie o contêiner Lima:
limactl start
Suba os serviços com Docker Compose:
lima nerdctl compose up
Este comando irá:

Construir a imagem Docker da aplicação.
Iniciar o banco de dados em um contêiner.
Iniciar a aplicação AskMe em outro contêiner.

Acesse a aplicação:

A aplicação estará disponível em: http://localhost:8081

Inicie a API (em um novo terminal):
go run cmd/askMe/main.go
A API estará disponível em uma porta definida no código fonte (verifique o arquivo cmd/askMe/main.go).

Uso
Com a aplicação e a API em execução, você pode:

Acessar a interface web: Utilize o navegador para interagir com a aplicação AskMe.
Utilizar a API: Envie requisições para a API para integrar com outras aplicações.
Documentação da API
A documentação da API estará disponível em breve.

Contribuindo
Sinta-se à vontade para contribuir com o projeto! Abra uma issue caso encontre algum problema ou tenha alguma sugestão.

Desenvolvido como projeto de estudo por:
@samuelmachado1
@techrastabr