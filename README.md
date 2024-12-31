# LibraryLabs - API REST para Biblioteca Online em Golang

**LibraryLabs** é uma aplicação em Golang que implementa uma API REST para gerenciar uma biblioteca online. A API permite o CRUD (Create, Read, Update, Delete) de livros, além de funcionalidades para consultar livros por nome ou gênero, e alterar o status de disponibilidade dos livros.

## Funcionalidades

### 1. Criar Livro
Permite a criação de um novo livro na biblioteca, informando detalhes como nome, autor, gênero e disponibilidade.  
A criação inclui validações de dados, garantindo que os campos obrigatórios sejam fornecidos corretamente.

### 2. Obter Todos os Livros
Retorna uma lista com todos os livros cadastrados na biblioteca.

### 3. Obter Livro por ID
Permite consultar um livro específico através de seu ID, retornando seus detalhes.

### 4. Atualizar Livro
Permite atualizar os dados de um livro existente (nome, autor, gênero e disponibilidade).  
Realiza verificações para garantir que o livro realmente exista antes de realizar a atualização.

### 5. Deletar Livro
Permite excluir um livro da biblioteca, identificado pelo seu ID.

### 6. Alterar Disponibilidade do Livro
Permite marcar um livro como disponível ou indisponível.  
A API garante que o livro esteja na disponibilidade desejada e que não haja tentativas de marcar o livro como disponível ou indisponível repetidamente sem alteração.

### 7. Buscar Livros por Nome
Permite buscar livros com base no nome fornecido. A consulta pode retornar vários livros que atendem ao critério.

### 8. Buscar Livros por Gênero
Permite buscar livros que pertencem a um determinado gênero (ex: "Fantasia", "Ficção Científica", etc.).
