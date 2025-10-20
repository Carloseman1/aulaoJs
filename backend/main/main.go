package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = "bocadas"
	cfg.Passwd = "root!"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "db_clientes"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pessoas", getPessoas)
	r.GET("/enderecos", getEnderecos)
	r.GET("/telefones", getTelefones)
	r.POST("/clientes", saveClientes)

	r.Run()
}

type Pessoa struct {
	ID           int    `json:"id"`
	NomeCompleto string `json:"nome_completo"`
	DataNasc     string `json:"data_nasc"`
	Genero       string `json:"genero"`
	CPF          string `json:"cpf"`
}

func getPessoas(c *gin.Context) {
	var pessoas []Pessoa

	rows, err := db.Query("SELECT id_pessoa, nome_completo, data_nasc, genero, cpf FROM tb_pessoas")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var pessoa Pessoa
		if err := rows.Scan(&pessoa.ID, &pessoa.NomeCompleto, &pessoa.DataNasc, &pessoa.Genero, &pessoa.CPF); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		pessoas = append(pessoas, pessoa)
	}

	c.JSON(200, gin.H{
		"data": pessoas,
	})
}

type Endereco struct {
	ID       int    `json:"id"`
	IDPessoa int    `json:"id_pessoa"`
	Rua      string `json:"rua"`
	Numero   int    `json:"numero"`
	Cidade   string `json:"cidade"`
	Estado   string `json:"estado"`
	CEP      string `json:"cep"`
}

func getEnderecos(c *gin.Context) {
	var enderecos []Endereco
	rows, err := db.Query("SELECT id_endereco, id_pessoa, rua, numero, cidade, estado, cep FROM tb_enderecos")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var endereco Endereco
		if err := rows.Scan(&endereco.ID, &endereco.IDPessoa, &endereco.Rua, &endereco.Numero, &endereco.Cidade, &endereco.Estado, &endereco.CEP); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		enderecos = append(enderecos, endereco)
	}
	c.JSON(200, gin.H{
		"data": enderecos,
	})
}

type Telefone struct {
	ID       int    `json:"id"`
	IDPessoa int    `json:"id_pessoa"`
	Numero   string `json:"numero"`
	Tipo     string `json:"tipo"`
}

func getTelefones(c *gin.Context) {
	var telefones []Telefone
	rows, err := db.Query("SELECT id_telefone, id_pessoa, numero, tipo FROM tb_telefones")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var telefone Telefone
		if err := rows.Scan(&telefone.ID, &telefone.IDPessoa, &telefone.Numero, &telefone.Tipo); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		telefones = append(telefones, telefone)
	}
	c.JSON(200, gin.H{
		"data": telefones,
	})
}

type Cliente struct {
	ID               int      `json:"id"`
	PrimeiroNome     string   `json:"primeiro_nome"`
	UltimoNome       string   `json:"ultimo_nome"`
	CPFFormatado     string   `json:"cpf_formatado"`
	Nascimento       string   `json:"nascimento"`
	Genero           string   `json:"genero"`
	EnderecoCompleto string   `json:"endereco_completo"`
	Telefones        []string `json:"telefones"`
}

func saveClientes(c *gin.Context) {
	var clientes Cliente

	if err := c.BindJSON(&clientes); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	telefonesJSON, err := json.Marshal(clientes.Telefones)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("INSERT INTO clientes (primeiro_nome, ultimo_nome, cpf_formatado, nascimento, genero, endereco_completo, telefones) VALUES (?, ?, ?, ?, ?, ?, ?)",
		clientes.PrimeiroNome, clientes.UltimoNome, clientes.CPFFormatado, clientes.Nascimento, clientes.Genero, clientes.EnderecoCompleto, telefonesJSON)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cliente salvo com sucesso!",
	})
}
