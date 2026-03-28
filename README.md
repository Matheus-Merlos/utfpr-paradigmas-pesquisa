<h1 align="center"> Pesquisador Paralelo </h1>
<p align="center">
  <a href="https://forthebadge.com">
    <img src="https://forthebadge.com/badges/made-with-go.svg" alt="Badge 1" />
  </a>
  <a href="https://forthebadge.com">
    <img src="https://forthebadge.com/badges/it-works-why.svg" alt="Badge 2" />
  </a>
</p>

Projeto desenvolvido para a disciplina de **Paradigmas de Programação** - UTFPR Dois Vizinhos.
Este projeto é um indexador de buscas desenvolvido em Go, focado em alta performance (e alto uso de CPU) através de paralelismo real.

### Tecnologias
- **Paralelismo com Goroutines**: Utiliza o modelo de concorrencia nativo do Go para processar múltiplos diretórios simultâneamnete.
- **Busca em Profundidade (DFS)**: Só porque é mais fácil de implementar com recursão,
- **Comunicação via Channels**: Implementa o padrão do próprio Go para centralizar os resultados da busca sem usar mutexes (eu e o Hasse estavamos com preguiça de usar Mutexes para falar a verdade)

### Como Executar

1. **Clone o Repositório**:
```bash
git clone https://github.com/seu-usuario/utfpr-paradigmas-pesquisa.git
cd utfpr-paradigmas-pesquisa
```

2. **Instale as depencências**:
```bash
go mod tidy
```

3. **Compile o projeto**:
```bash
go build -o super-grep
```

4. **Execute a busca**:
```bash
./super-grep prego
```