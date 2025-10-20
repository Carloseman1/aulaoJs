CREATE DATABASE db_cliente;
USE db_cliente;

CREATE TABLE tb_pessoas (
  id_pessoa INT PRIMARY KEY,
  nome_completo VARCHAR(100),
  data_nasc DATE,
  genero CHAR(1),
  cpf VARCHAR(11)
);

CREATE TABLE tb_enderecos (
  id_endereco INT PRIMARY KEY,
  id_pessoa INT,
  rua VARCHAR(100),
  numero INT,
  cidade VARCHAR(100),
  estado CHAR(2),
  cep VARCHAR(9),
  FOREIGN KEY (id_pessoa) REFERENCES tb_pessoas(id_pessoa)
);

CREATE TABLE tb_telefones (
  id_telefone INT PRIMARY KEY,
  id_pessoa INT,
  tipo VARCHAR(20),
  numero VARCHAR(15),
  FOREIGN KEY (id_pessoa) REFERENCES tb_pessoas(id_pessoa)
);

CREATE TABLE clientes (
  id_cliente INT PRIMARY KEY AUTO_INCREMENT,
  primeiro_nome VARCHAR(50),
  ultimo_nome VARCHAR(50),
  cpf_formatado VARCHAR(14),
  nascimento TIMESTAMP,
  genero CHAR(1),
  endereco_completo TEXT,
  telefones JSON
);


INSERT INTO tb_pessoas VALUES
(1, 'João da Silva', '1988-03-10', 'M', '11122233344'),
(2, 'Maria Oliveira', '1992-08-25', 'F', '22233344455'),
(3, 'Pedro Santos Neto', '1979-11-02', 'M', '33344455566'),
(4, 'Ana Paula Souza', '1995-05-15', 'F', '44455566677'),
(5, 'Carlos Alberto Lima', '1983-12-22', 'M', '55566677788'),
(6, 'Fernanda Rodrigues', '1990-09-03', 'F', '66677788899'),
(7, 'Ricardo Pereira', '1987-07-30', 'M', '77788899900'),
(8, 'Juliana Castro', '1994-01-09', 'F', '88899900011'),
(9, 'Marcos Vinicius Lopes', '1982-10-17', 'M', '99900011122'),
(10, 'Patrícia Nogueira', '1998-04-21', 'F', '00011122233');

-- Endereços
INSERT INTO tb_enderecos VALUES
(1, 1, 'Rua das Flores', 123, 'São Paulo', 'SP', '01001-000'),
(2, 2, 'Av. Brasil', 456, 'Rio de Janeiro', 'RJ', '20020-000'),
(3, 3, 'Rua das Palmeiras', 789, 'Belo Horizonte', 'MG', '30030-000'),
(4, 4, 'Rua dos Jasmins', 101, 'Curitiba', 'PR', '80040-000'),
(5, 5, 'Av. das Nações', 555, 'Brasília', 'DF', '70070-000'),
(6, 6, 'Rua da Liberdade', 333, 'Salvador', 'BA', '40040-000'),
(7, 7, 'Av. Independência', 890, 'Porto Alegre', 'RS', '90090-000'),
(8, 8, 'Rua XV de Novembro', 202, 'Florianópolis', 'SC', '88010-000'),
(9, 9, 'Rua São João', 404, 'Fortaleza', 'CE', '60060-000'),
(10, 10, 'Av. Beira Mar', 777, 'Recife', 'PE', '50050-000');

-- Telefones
INSERT INTO tb_telefones VALUES
(1, 1, 'celular', '11999998888'),
(2, 2, 'fixo', '2133334444'),
(3, 3, 'celular', '31988887777'),
(4, 4, 'celular', '41999996666'),
(5, 5, 'fixo', '6132223344'),
(6, 6, 'celular', '71988885555'),
(7, 7, 'celular', '51999994444'),
(8, 8, 'fixo', '4833332222'),
(9, 9, 'celular', '85999993333'),
(10, 10, 'celular', '81988882222');