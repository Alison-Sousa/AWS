package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	sess := session.Must(session.NewSession())

	// Configurar o cliente do serviço S3
	svc := s3.New(sess)

	// Iterar sobre os eventos de S3
	for _, record := range s3Event.Records {
		s3Bucket := record.S3.Bucket.Name
		s3Object := record.S3.Object.Key

		// Obter metadados do objeto S3
		input := &s3.HeadObjectInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String(s3Object),
		}

		result, err := svc.HeadObject(input)
		if err != nil {
			log.Println("Erro ao obter metadados do objeto S3:", err)
			return err
		}

		// Imprimir metadados do objeto S3
		fmt.Println("Metadados do objeto S3:")
		fmt.Println("Nome do bucket:", s3Bucket)
		fmt.Println("Nome do objeto:", s3Object)
		fmt.Println("Tamanho do objeto:", *result.ContentLength)
		fmt.Println("Tipo de conteúdo:", *result.ContentType)
		fmt.Println("Última modificação:", *result.LastModified)
		fmt.Println("ETag:", *result.ETag)
	}

	return nil
}

func main() {
	// Se estiver executando localmente (fora do ambiente AWS Lambda)
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		// Simular um evento S3 para testar localmente
		s3Event := events.S3Event{
			Records: []events.S3EventRecord{
				{
					S3: events.S3Entity{
						Bucket: events.S3Bucket{Name: "my-bucket"},
						Object: events.S3Object{Key: "my-file.txt"},
					},
				},
			},
		}

		// Chamar o manipulador localmente
		if err := handler(context.Background(), s3Event); err != nil {
			log.Fatalf("Erro ao executar o manipulador: %v", err)
		}
		return
	}

	// Iniciar a execução da função Lambda
	lambda.Start(handler)
}
