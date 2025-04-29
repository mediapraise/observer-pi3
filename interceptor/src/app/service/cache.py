
import time
from datetime import datetime, timedelta

class CacheInterceptor:
    """ Classe para armazenar e gerenciar o cache de placas e seus status. """

    def __init__(self, tempo_expiracao_segundos=600):
        """
        Inicializa o cache para armazenar dados de placa e status.

        Args:
            tempo_expiracao_segundos (int): Tempo em segundos para expiração do cache (padrão: 10 minutos = 600 segundos).
        """

        self.cache = {}
        self.tempo_expiracao = timedelta(seconds=tempo_expiracao_segundos)


    def armazenar(self, placa: str, status: str):
        """
        Armazena o status de uma placa no cache.

        Args:
            placa (str): A identificação da placa.
            status (str): O status da placa a ser armazenado.
        """

        timestamp_expiracao = datetime.now() + self.tempo_expiracao
        self.cache[placa] = {"status": status, "expira_em": timestamp_expiracao}
        print(f"Cache atualizado para a placa '{placa}': {self.cache[placa]}")


    def obter_status(self, placa: str) -> str | None:
        """
        Obtém o status de uma placa do cache, se disponível e não expirado.

        Args:
            placa (str): A identificação da placa.

        Returns:
            str | None: O status da placa se encontrado e não expirado, caso contrário, None.
        """

        if placa in self.cache:
            dados_cache = self.cache[placa]
            if datetime.now() < dados_cache["expira_em"]:
                print(f"Status da placa '{placa}' encontrado no cache: {dados_cache['status']}")
                return dados_cache["status"]
            else:
                print(f"Cache para a placa '{placa}' expirou.")
                del self.cache[placa]
                return None
        else:
            print(f"Status da placa '{placa}' não encontrado no cache.")
            return None


    def limpar_cache(self):
        """
        Limpa todo o cache.
        """
        self.cache = {}
        print("Cache limpo.")
