const AWS = require('aws-sdk');

// Configuração da AWS (simulação)
AWS.config.update({
  region: 'us-east-1', // Simulação da região
  credentials: new AWS.Credentials('FAKE_ACCESS_KEY_ID', 'FAKE_SECRET_ACCESS_KEY', 'FAKE_SESSION_TOKEN'),
});

// Criando um novo cliente DynamoDB (simulação)
const dynamoDB = new AWS.DynamoDB.DocumentClient();

// Função para adicionar uma nova tarefa fictícia ao DynamoDB
async function adicionarTarefaFicticia(tarefa) {
  console.log(`Adicionando tarefa fictícia: ${tarefa.descricao}`);
  // Simulação de adição de tarefa ao DynamoDB
  console.log('Tarefa fictícia adicionada com sucesso!');
}

// Função para listar todas as tarefas fictícias do DynamoDB
async function listarTarefasFicticias() {
  console.log('Lista de tarefas fictícias:');
  console.log('- [ ] Tarefa 1');
  console.log('- [x] Tarefa 2 (concluída)');
}

// Função para marcar uma tarefa fictícia como concluída
async function marcarComoConcluidaFicticia(id) {
  console.log(`Marcando tarefa fictícia com id ${id} como concluída`);
  // Simulação de marcação de tarefa como concluída no DynamoDB
  console.log('Tarefa fictícia marcada como concluída com sucesso!');
}

// Exemplo de uso das funções fictícias
(async () => {
  await adicionarTarefaFicticia({ descricao: 'Fazer compras' });
  await listarTarefasFicticias();
  await marcarComoConcluidaFicticia('2');
})();
