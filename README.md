# Challenge Multithreading

- É realizado duas requisições simultaneamente para as seguintes APIs:

    - https://cdn.apicep.com/file/apicep/" + cep + ".json

    - http://viacep.com.br/ws/" + cep + "/json/

- Os requisitos para este desafio são:

  - Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
  
  - O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
  
  - Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

- Example Request: 
  - localhost:8080/?cep=10000123

- Example Response:
  - Via Cep:

        { 
          "cep": "01001-000",
          "logradouro": "Praça da Sé",
          "complemento": "lado ímpar",
          "bairro": "Sé",
          "localidade": "São Paulo",
          "uf": "SP",
          "ibge": "3550308",
          "gia": "1004",
          "ddd": "11",
          "siafi": "7107"
        }

  - BrasilAPI:
  
        {
          "cep": "89010025",
          "state": "SC",
          "city": "Blumenau",
          "neighborhood": "Centro",
          "street": "Rua Doutor Luiz de Freitas Melro",
          "service": "viacep"
        }