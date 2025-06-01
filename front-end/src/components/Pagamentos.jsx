import { useState, useEffect } from "react";
import axios from "axios";

const Pagamentos = () => {
  const [payments, setPayments] = useState([]);

  const fetchPayments = async () => {
    try {
      const response = await axios.get(
        "https://observerpi3.onrender.com/finance"
      );
      setPayments(response.data);
    } catch (error) {
      console.error("Erro ao buscar pagamentos:", error);
    }
  };

  useEffect(() => {
    fetchPayments();
  }, []);

  return (
    <div>
      <h1>📜 Pagamentos - Relatório</h1>

      <table border="1">
        <thead>
          <tr>
            <th>Nome</th>
            <th>Placa</th>
            <th>Data do Pagamento</th>
            <th>Valor</th>
          </tr>
        </thead>
        <tbody>
          {payments.length > 0 ? (
            payments.map((payment, index) => (
              <tr key={index}>
                <td>{payment.name}</td>
                <td>{payment.board}</td>
                <td>{payment.payDate}</td>
                <td>R$ {parseFloat(payment.payAmount).toFixed(2)}</td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="4" style={{ textAlign: "center", color: "red" }}>
                Nenhum pagamento registrado.
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default Pagamentos;
