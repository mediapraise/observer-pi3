import { useState } from "react";
import axios from "axios";

const labelNames = {
  name: "Nome Completo",
  document: "Documento (CPF/CNPJ)",
  email: "E-mail",
  phone: "Telefone",
  vehicle: "Veículo",
  pay_day: "Dia de Pagamento",
};

const CadClient = ({ onSuccess }) => {
  const [formData, setFormData] = useState({
    name: "",
    document: "",
    email: "",
    phone: "",
    vehicle: "",
    pay_day: "",
  });

  const [successMessage, setSuccessMessage] = useState(false);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        "https://observerpi3.onrender.com/clients",
        formData
      );

      if (response.status === 201) {
        setFormData({
          name: "",
          document: "",
          email: "",
          phone: "",
          vehicle: "",
          pay_day: "",
        });
        setSuccessMessage(true);
        onSuccess();
      }
    } catch (error) {
      console.error("Erro ao cadastrar cliente:", error);
    }
  };

  return (
    <div>
      <h1>Cadastrar Mensalista</h1>
      {successMessage ? (
        <h2 style={{ color: "green" }}>
          ✅ Mensalista cadastrado com sucesso!
        </h2>
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
                type="text"
                id={key}
                name={key}
                value={value}
                onChange={handleChange}
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
            Cadastrar
          </button>
        </form>
      )}
    </div>
  );
};

export default CadClient;
