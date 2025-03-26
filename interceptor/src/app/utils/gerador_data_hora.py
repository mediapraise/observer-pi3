import datetime

class GeradorDataHora:
    '''Classe para gerar data e hora '''

    def gerador_data_hora():
        '''Função para gerar data e hora atual. Formato -> 2025-03-25 23:27:15.470660 '''

        today = datetime.datetime.now()

        # Converte string em datetime
        #today = today.strftime("%Y-%m-%d %H:%M:%S")

        # Converte string em datetime
        # today = datetime.datetime.strptime(today, "%Y-%m-%d %H:%M:%S")

        return today
