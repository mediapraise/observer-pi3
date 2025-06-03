import { useState, useEffect } from "react";
import axios from "axios";

const Mensalistas = () => {
  const [clients, setClients] = useState([]);

  const fetchClients = async () => {
    try {
      const response = await axios.get(
        "https://observerpi3.onrender.com/clients"
      );
      setClients(response.data);
    } catch (error) {
      console.error("Erro ao buscar clientes:", error);
    }
  };

  useEffect(() => {
    fetchClients();
  }, []);

  return (
    <div>
      <h1>Mensalistas - Cadastro Geral</h1>

      <table border="1">
        <thead>
          <tr>
            <th>Nome</th>
            <th>Documento</th>
            <th>Email</th>
            <th>Telefone</th>
            <th>Veículo</th>
            <th>Dia do Pagamento</th>
          </tr>
        </thead>
        <tbody>
          {clients.map((client) => (
            <tr key={client.document}>
              <td>{client.name}</td>
              <td>{client.document}</td>
              <td>{client.email}</td>
              <td>{client.phone}</td>
              <td>{client.vehicle}</td>
              <td>{client.pay_day}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Mensalistas;
