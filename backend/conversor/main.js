const fetchPessoas = async() => { 
    try {  
        const res = await fetch("http://localhost:8080/pessoas")
        if(!res.ok) throw new Error (`http ${res.status}`)
        return await res.json()
    } catch (err) {
        console.error("Erro: ", err);
        throw err
    }
}

const fetchTelefones = async () => {
    try {
        const res = await fetch("http://localhost:8080/telefones")
        if(!res.ok) throw new Error (`http ${res.status}`)
        return await res.json()
    } catch (err) {
        console.log("Erro: ", err);
        throw err
    }
}

const formatarCPF = (cpf) => {
  
    cpf = cpf.replace(/\D/g, "");
  
    cpf = cpf.replace(/(\d{3})(\d)/, "$1.$2");
    cpf = cpf.replace(/(\d{3})(\d)/, "$1.$2");
    cpf = cpf.replace(/(\d{3})(\d{1,2})$/, "$1-$2");

  return cpf;
}

const fetchEnderecos = async () => {
    try {
        const res = await fetch("http://localhost:8080/enderecos")
        if(!res.ok) throw new Error (`http ${res.status}`)
        return await res.json()
    } catch (err) {
        console.log("Erro: ", err);
        throw err
    }
}

const salvarCliente = async (pessoa) => {
    try {
        const res = await fetch("http://localhost:8080/clientes",{
            method : "POST", 
            headers : {
                "Content-Type":"application/json"
            },
            body : JSON.stringify(pessoa)
        })
        if(!res.ok) throw new Error (`http ${res.status}`)
        return await res.json()
    } catch (err) {
        console.log("Erro: ", err);
        throw err
    }
}

const main = async () => {
    let pessoas = await fetchPessoas()
    let enderecos = await fetchEnderecos()
    let telefones = await fetchTelefones()

    const clientes = pessoas.data.map(pessoa => {
        let cliente = {}
        var cpf = pessoa.cpf
        var nomeCompleto = pessoa.nome_completo
        var genero = pessoa.genero
        var dataNascimento = pessoa.data_nasc

        var nomeSplit = nomeCompleto.split(" ");
        cliente.primeiro_nome = nomeSplit[0]
        cliente.ultimo_nome = nomeSplit[nomeSplit.length-1]
        
        cliente.cpf_formatado = formatarCPF(cpf)
        cliente.nascimento = dataNascimento
        cliente.genero = genero
        
        const endereco = enderecos.data.find(e => e.id_pessoa === pessoa.id)
        
        cliente.endereco_completo = `${endereco.rua} ,${endereco.numero} , ${endereco.cep} - ${endereco.cidade}/${endereco.estado}`
        
        const telefone = telefones.data.filter(e => e.id_pessoa === pessoa.id)
    
        cliente.telefones = telefone.map(telefone =>{
            return telefone.numero
        });

        return cliente

    });

    clientes.forEach(cliente => {
        salvarCliente(cliente)
    });
}

main()

