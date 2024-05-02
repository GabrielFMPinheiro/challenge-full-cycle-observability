Full Cycle Challenge

API: http://localhost:8081/zipcode
Body: { "zipcode": "DIGITE UM CEP AQUI" }

Retorno: {"temp_C":26,"temp_F":78.8,"temp_K":299,"city":"Vitória"}

Como rodar em dev:

Inicie baixando o projeto na sua máquina;
Faça a criação de uma conta na API https://www.weatherapi.com/;
Na sua máquina faça uma cópia do arquivo .env.example e renomeie para .env;
Na sua conta feita nesse site: https://www.weatherapi.com/, procute pela Api Key e cole a mesma na variável de ambiente WEATHER_API_KEY;
Rode o docker compose utilizando o comando: docker compose up;
Acesse os containers app-a e app-b e inicie a aplicação com o comando: go run main.go;
