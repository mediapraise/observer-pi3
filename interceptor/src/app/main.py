import os
from utils import status
from settings import parameters

objeto_json = {}

def main():
    '''Função principal do programa'''

    print("--------- Interceptor Project ---------")
    print("--------- Versão 1.0.0 ---------")
    print("--------- Autor: DRP01-PJI310-SALA-002GRUPO-014 ---------")
    print("--------- Data: 2025-05-20 ---------")
    print("--------- Inicializando Programa ---------")

    # Exemplo e teste de uso de enum
    print(f"Status Enum Value: {status.Status.entrada.value}")
    # print(f"Status Enum Value: {status.Status.saida.value}")
    print(f"URL: {parameters.Parametros.url_api_golang.value}")
    
    # TODO: Capturar imagens da câmera

    # TODO: Extração de strings de imagens

    # TODO: Criar um dict [objeto JSON]  com a strings extraídas


    # TODO: Coletar data e hora do sistema

    # TODO: Enviar o dict [objeto JSON] para API com: PLACA, DATA_HORA, STATUS

    
    print("--------- Finalizando Programa ---------")

if __name__ == "__main__":
    main()
