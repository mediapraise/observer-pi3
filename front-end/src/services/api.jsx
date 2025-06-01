import axios from "axios";

// Configurando a API Axios para se conectar ao backend
const api = axios.create({
  baseURL: "https://observerpi3.onrender.com/", // Backend rodando na porta 3001
});

export default api;
