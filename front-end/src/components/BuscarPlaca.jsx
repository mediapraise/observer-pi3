import { useState } from "react";
import axios from "axios";

const BuscarPlaca = () => {
  const [board, setBoard] = useState("");
  const [vehicleData, setVehicleData] = useState(null);
  const [errorMessage, setErrorMessage] = useState("");

  const handleSearch = async () => {
    try {
      const response = await axios.get(
        `https://observerpi3.onrender.com/log?board=${board}`
      );
      if (response.data.length > 0) {
        setVehicleData(response.data);
        setErrorMessage("");
      } else {
        setVehicleData(null);
        setErrorMessage("Nenhum registro encontrado para esta placa.");
      }
    } catch (error) {
      console.error("Erro ao buscar placa:", error);
      setErrorMessage("Erro ao buscar dados. Tente novamente.");
    }
  };

  return (
    <div>
      <h1>🔍 Histórico de Acesso</h1>
      <label>
        <b>Digite a placa do veículo:</b>&nbsp;
      </label>
      <input
        type="text"
        value={board}
        onChange={(e) => setBoard(e.target.value)}
      />
      <button onClick={handleSearch}>Buscar</button>
      <br />
      <br />
      <br />

      {errorMessage && <p style={{ color: "red" }}>{errorMessage}</p>}

      {vehicleData && (
        <table border="1">
          <thead>
            <tr>
              <th>Nome</th>
              <th>Placa</th>
              <th>Entrada</th>
              <th>Saída</th>
            </tr>
          </thead>
          <tbody>
            {vehicleData.map((entry, index) => (
              <tr key={index}>
                <td>{entry.name}</td>
                <td>{entry.board}</td>
                <td>{entry.entrance}</td>
                <td>{entry.exit}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default BuscarPlaca;
