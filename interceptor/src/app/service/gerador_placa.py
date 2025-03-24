import random

def gerar_placa():
    """Gera uma placa de veículo brasileira aleatória."""

    letras = ''.join(random.choices('ABCDEFGHIJKLMNOPQRSTUVWXYZ', k=3))
    numeros = ''.join(random.choices('0123456789', k=4))

    # Formato da placa: LLL-NNNN (onde L é letra e N é número)
    placa = f"{letras}-{numeros}"
    return placa

# Exemplo de uso
for _ in range(5):  # Gera 5 placas
    print(gerar_placa())