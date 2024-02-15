const AWS = require('aws-sdk');
const s3 = new AWS.S3();

exports.handler = async (event) => {
    // Gerar dados fictícios
    const data = [
        { name: 'John', age: 30 },
        { name: 'Alice', age: 25 },
        { name: 'Bob', age: 35 }
    ];
    
    // Converter dados em formato JSON
    const jsonData = JSON.stringify(data);
    
    // Parâmetros para o upload do arquivo no S3
    const params = {
        Bucket: 'my-bucket', // Substitua pelo nome do seu bucket no S3
        Key: 'data.json', // Nome do arquivo no S3
        Body: jsonData
    };
    
    // Fazer upload do arquivo no S3
    try {
        await s3.upload(params).promise();
        return {
            statusCode: 200,
            body: 'Dados enviados para o S3 com sucesso.'
        };
    } catch (err) {
        console.error('Erro ao fazer upload dos dados no S3:', err);
        return {
            statusCode: 500,
            body: 'Erro ao fazer upload dos dados no S3.'
        };
    }
};
