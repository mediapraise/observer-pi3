import { useState } from "react";
import axios from "axios";

const labelNames = {
  name: "Nome do Cliente",
  board: "Placa do Veículo",
  payDate: "Data do Pagamento",
  payAmount: "Valor Pago",
};

const Pagamento = ({ onSuccess }) => {
  const [formData, setFormData] = useState({
    name: "",
    board: "",
    payDate: new Date().toISOString().split("T")[0], // Captura a data atual automaticamente
    payAmount: "",
  });

  const [successMessage, setSuccessMessage] = useState(false);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        "https://observerpi3.onrender.com/finance",
        formData
      );

      if (response.status === 201) {
        setFormData({
          name: "",
          board: "",
          payDate: new Date().toISOString().split("T")[0], // Reatribui a data ao limpar
          payAmount: "",
        });
        setSuccessMessage(true);
        onSuccess();
      }
    } catch (error) {
      console.error("Erro ao cadastrar pagamento:", error);
    }
  };

  return (
    <div>
      {successMessage ? (
        <h2 style={{ color: "green" }}>✅ Pagamento registrado com sucesso!</h2>
      ) : (
        <form
          onSubmit={handleSubmit}
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          {Object.entries(formData).map(([key, value]) => (
            <div
              key={key}
              style={{
                display: "flex",
                flexDirection: "column",
                gap: "5px",
                marginBottom: "10px",
              }}
            >
              <label htmlFor={key} style={{ fontWeight: "bold" }}>
                {labelNames[key]}:
              </label>
              <input
                type={key === "payAmount" ? "number" : "text"}
                id={key}
                name={key}
                value={value}
                onChange={handleChange}
                disabled={key === "payDate"} // Impede a edição manual da data
                style={{
                  padding: "5px",
                  border: "1px solid #ccc",
                  borderRadius: "5px",
                  width: "100%",
                }}
              />
            </div>
          ))}
          <button
            type="submit"
            style={{
              padding: "10px 15px",
              backgroundColor: "#28a745",
              color: "white",
              border: "none",
              borderRadius: "5px",
              cursor: "pointer",
            }}
          >
            Registrar Pagamento
          </button>
        </form>
      )}
    </div>
  );
};

export default Pagamento;
