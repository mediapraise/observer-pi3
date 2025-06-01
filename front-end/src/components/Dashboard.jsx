import React, { useState, useEffect } from "react";
import axios from "axios";

const Dashboard = () => {
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [notification, setNotification] = useState(null);

  const playBeep = () => {
    const audioContext = new (window.AudioContext || window.webkitAudioContext)();
    const oscillator = audioContext.createOscillator();
    oscillator.type = "square"; // Som clássico tipo "beep"
    oscillator.frequency.setValueAtTime(1000, audioContext.currentTime); // Frequência do som (1000Hz)
    
    const gainNode = audioContext.createGain();
    gainNode.gain.setValueAtTime(0.2, audioContext.currentTime); // Volume
    
    oscillator.connect(gainNode);
    gainNode.connect(audioContext.destination);
    
    oscillator.start();
    setTimeout(() => {
      oscillator.stop(); // Para o som após 300ms
      audioContext.close();
    }, 300);
  };

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await axios.get("https://observerpi3.onrender.com/events");
        const newEvents = response.data.reverse();

        if (newEvents.length > 0 && (!events.length || newEvents[0].id !== events[0].id)) {
          // 🚗 Mensagem + som ao detectar novo evento
          setNotification(`🚗 Novo evento: ${newEvents[0].vehicle} - ${newEvents[0].board}`);
          playBeep(); // 🔊 Toca o alerta sonoro

          setTimeout(() => {
            setNotification(null);
          }, 3000);
        }

        setEvents(newEvents);
        setLoading(false);
      } catch (error) {
        console.error("Erro ao buscar eventos:", error);
        setError("Não foi possível carregar os eventos.");
        setLoading(false);
      }
    };

    fetchEvents(); // Carregamento inicial

    const interval = setInterval(() => {
      fetchEvents();
    }, 10000);

    return () => clearInterval(interval);
  }, [events]); // Agora a lógica depende da mudança de `events`!

  return (
    <div>
      <h1>Status Atual</h1>

      {notification && (
        <div style={{
          position: "fixed",
          top: "50%",
          left: "50%",
          transform: "translate(-50%, -50%)",
          padding: "20px",
          background: "rgba(0, 128, 0, 0.9)",
          color: "white",
          borderRadius: "10px",
          fontWeight: "bold",
          fontSize: "24px",
          textAlign: "center",
          boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.3)",
          zIndex: 1000
        }}>
          {notification}
        </div>
      )}

      {loading ? (
        <p>Carregando eventos...</p>
      ) : error ? (
        <p style={{ color: "red" }}>{error}</p>
      ) : events.length === 0 ? (
        <p>Nenhum evento encontrado.</p>
      ) : (
        <table border="1">
          <thead>
            <tr>
              <th>Veículo</th>
              <th>Placa</th>
              <th>Dia do Pagamento</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {events.map((event) => (
              <tr key={event.id}>
                <td>{event.vehicle || "N/A"}</td>
                <td>{event.board || "N/A"}</td>
                <td>{event.payday || "N/A"}</td>
                <td>{event.status || "N/A"}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default Dashboard;
