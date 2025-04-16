# MCP Server Govbox

Um servidor MCP (Model Context Protocol) desenvolvido para facilitar a integração entre Large Language Models (LLMs) e tarefas do Rakefile do projeto Govbox.

## 📋 Sobre o Projeto

Este servidor atua como uma ponte entre LLMs e comandos Rake, permitindo que modelos de linguagem executem tarefas automatizadas de forma segura e controlada. O projeto foi desenvolvido para simplificar a execução de comandos Rake, tornando-os acessíveis, sem necessidade de conhecimento prévio dos comandos específicos.

### 🎯 Objetivo

O principal objetivo é permitir que LLMs possam:
- Executar tarefas do Rakefile de forma programática
- Automatizar processos comuns
- Reduzir a necessidade de conhecimento específico dos comandos Rake
- Prover uma interface padronizada para execução de tarefas

### 🚀 Funcionalidades

- Execução segura e controlada de tarefas
- Suporte a múltiplos projetos
- Retorno padronizado de resultados
- Tratamento de erros robusto

### 🛠️ Tecnologias Utilizadas

- Go (Golang) para o servidor MCP
- Ruby/Rake para execução das tarefas
- Docker para containerização

## ⚙️ Instalação e Configuração

### 📦 Passo 1: Obtenção do Código

- #### Clone o repositório
```bash
git clone https://github.com/giovanoh/mcp-server-govbox.git
cd mcp-server-govbox
```

### 🔨 Passo 2: Build

- #### Usando Docker
```bash
docker build -t mcp-server-govbox .
```

- #### Executando localmente
```bash
go build cmd/app/main.go
```

### 🔧 Passo 3: Configuração do Server

O servidor utiliza variáveis de ambiente para sua configuração. Todas as variáveis de ambiente devem ser configuradas para o funcionamento adequado do servidor. Para começar:

```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

As seguintes variáveis de ambiente são necessárias:

- `MCP_GOVBOX_PROJECT_PATH`: Caminho para o diretório do projeto Govbox
- `MCP_SHELL_PATH`: Shell a ser usado para execução dos comandos (ex: /bin/zsh)
- `MCP_SHELL_OPTIONS`: Argumentos adicionais para o shell (ex: -c)

**Nota importante**: O arquivo .env deve estar localizado na pasta raiz do projeto client que está chamando o servidor. Por exemplo:
- Se estiver utilizando com o Claude Desktop no Windows, o arquivo deve ficar em: `C:\Users\<Usuário>\AppData\Local\AnthropicClaude\app-<Versão>\.env`
- Se estiver utilizando com um client próprio, coloque o .env na pasta raiz do client

Certifique-se de que todas as variáveis estejam devidamente configuradas no arquivo .env antes de iniciar o servidor.

### 🔧 Passo 4: Configuração do Client

Para utilizar este servidor com o aplicativo Claude Desktop como client, ajuste a configuração do seu arquivo `claude_desktop_config.json` de acordo com os exemplos:

#### Para uso com Docker
```json
{
    "mcpServers": {
        "mcp-server-govbox": {
            "command": "docker",
            "args": [
                "run",
                "--rm",
                "-i",
                "mcp-server-govbox:latest"
            ]
        }
    }
}
```

#### Para uso do executável diretamente
```json
{
    "mcpServers": {
        "mcp-server-govbox": {
            "command": "/path/to/executable"
        }
    }
}
```

Para outros clients, copie a configuração desejada e configure de acordo com o modelo utilizado pelo seu client de LLM.

```bash
# Usando Docker
cp config_docker_example.json config.json
# OU para executável
cp config_executable_example.json config.json
```

## ❌ Possíveis Erros

Durante a execução do servidor, alguns erros podem ocorrer e serão retornados para o client. Abaixo estão os possíveis erros e suas causas:

### Erros de Inicialização
- `"Error loading .env file"`:
  - Ocorre quando o arquivo .env não é encontrado
  - Ou quando há problemas de permissão para ler o arquivo
  - Ou quando o arquivo está mal formatado

- `"Invalid shell configuration"`:
  - Ocorre quando as variáveis de ambiente do shell estão incorretas ou ausentes
  - Possíveis causas:
    - `MCP_GOVBOX_PROJECT_PATH` não foi definido ou está vazio
    - `MCP_SHELL_PATH` não foi definido ou está vazio
    - Shell especificado não existe no sistema

### Erros de Validação
- `"Param 'projects' must be a string"`: 
  - Ocorre quando o parâmetro 'projects' não foi fornecido na requisição
  - Ou quando o valor fornecido não é uma string válida

### Erros de Execução
- `"Error during command execution (start)"`:
  - Ocorre quando há problemas ao iniciar o comando
  - Possíveis causas: permissões insuficientes, shell não encontrado, variáveis de ambiente mal configuradas

- `"Error during command execution (wait)"`:
  - Ocorre quando há problemas durante a execução do comando
  - Possíveis causas: comando interrompido abruptamente, problemas de recursos do sistema

- `"Error creating stdout/stderr pipe"`:
  - Ocorre quando o servidor não consegue criar os pipes para capturar a saída do comando
  - Possíveis causas: problemas de permissão, limites do sistema operacional atingidos

### Como Resolver
1. Verifique se todas as variáveis de ambiente estão configuradas corretamente no arquivo .env
2. Certifique-se de que o arquivo .env está no local correto
3. Verifique se o usuário tem permissões adequadas para:
   - Ler o arquivo .env
   - Executar o shell configurado
   - Acessar o diretório do projeto Govbox
4. Confirme se o shell configurado existe e está acessível
5. Verifique se o caminho do projeto Govbox está correto
6. Verifique os logs do servidor para mais detalhes sobre o erro

## 📝 Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
