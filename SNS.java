import com.amazonaws.services.sns.AmazonSNS;
import com.amazonaws.services.sns.AmazonSNSClient;
import com.amazonaws.services.sns.model.PublishRequest;
import com.amazonaws.services.sns.model.PublishResult;

public class SnsExample {
    public static void main(String[] args) {
        // Configurando o cliente SNS
        AmazonSNS snsClient = AmazonSNSClient.builder().build();

        // Simulando a publicação de uma mensagem no tópico do SNS
        String topicArn = "arn:aws:sns:us-east-1:123456789012:MeuTopicoSNS"; // ARN fictício
        String message = "Mensagem de exemplo do SNS";
        PublishRequest publishRequest = new PublishRequest(topicArn, message);
        PublishResult publishResult = snsClient.publish(publishRequest);
        System.out.println("Mensagem publicada no tópico do SNS.");
    }
}
