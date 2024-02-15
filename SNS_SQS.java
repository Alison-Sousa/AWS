import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.sns.AmazonSNS;
import com.amazonaws.services.sns.AmazonSNSClientBuilder;
import com.amazonaws.services.sns.model.PublishRequest;
import com.amazonaws.services.sns.model.PublishResult;
import com.amazonaws.services.sqs.AmazonSQS;
import com.amazonaws.services.sqs.AmazonSQSClientBuilder;
import com.amazonaws.services.sqs.model.ReceiveMessageRequest;
import com.amazonaws.services.sqs.model.ReceiveMessageResult;
import com.amazonaws.services.sqs.model.Message;

public class SnsSqsExample {
    public static void main(String[] args) {
        // Configuração fictícia das credenciais
        String accessKey = "FAKE_ACCESS_KEY";
        String secretKey = "FAKE_SECRET_KEY";

        // Criando credenciais fictícias
        BasicAWSCredentials credentials = new BasicAWSCredentials(accessKey, secretKey);

        // Configurando o cliente SNS
        AmazonSNS snsClient = AmazonSNSClientBuilder.standard()
                .withRegion(Regions.US_EAST_1) // Simulação da região
                .withCredentials(new AWSStaticCredentialsProvider(credentials))
                .build();

        // Configurando o cliente SQS
        AmazonSQS sqsClient = AmazonSQSClientBuilder.standard()
                .withRegion(Regions.US_EAST_1) // Simulação da região
                .withCredentials(new AWSStaticCredentialsProvider(credentials))
                .build();

        // Simulando a publicação de uma mensagem no tópico do SNS
        String topicArn = "arn:aws:sns:us-east-1:123456789012:MeuTopicoSNS"; // ARN fictício
        String message = "Mensagem de exemplo do SNS";
        PublishRequest publishRequest = new PublishRequest(topicArn, message);
        PublishResult publishResult = snsClient.publish(publishRequest);
        System.out.println("Mensagem publicada no tópico do SNS.");

        // Simulando a recepção da mensagem da fila SQS
        String queueUrl = "https://sqs.us-east-1.amazonaws.com/123456789012/MinhaFilaSQS"; // URL fictícia
        ReceiveMessageRequest receiveMessageRequest = new ReceiveMessageRequest(queueUrl);
        ReceiveMessageResult receiveMessageResult = sqsClient.receiveMessage(receiveMessageRequest);
        for (Message msg : receiveMessageResult.getMessages()) {
            System.out.println("Mensagem recebida da fila SQS: " + msg.getBody());
        }
    }
}
