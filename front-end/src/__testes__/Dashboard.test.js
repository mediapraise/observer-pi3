import { render, screen } from "@testing-library/react";
import Dashboard from "../components/Dashboard";
import axios from "axios";
import { act } from "react-dom/test-utils";

// Simula a resposta da API
jest.mock("axios");

describe("Dashboard Component", () => {
  it("exibe mensagem de carregamento inicialmente", () => {
    render(<Dashboard />);
    expect(screen.getByText("Carregando eventos...")).toBeInTheDocument();
  });

  it("exibe erro ao falhar na requisição", async () => {
    axios.get.mockRejectedValueOnce(new Error("Erro ao buscar eventos"));

    await act(async () => {
      render(<Dashboard />);
    });

    expect(screen.getByText("Não foi possível carregar os eventos.")).toBeInTheDocument();
  });

  it("exibe eventos corretamente ao carregar com sucesso", async () => {
    axios.get.mockResolvedValueOnce({
      data: [
        { id: 1, vehicle: "Carro", board: "ABC-1234", payday: "2025-06-07", status: "Pago" },
      ],
    });

    await act(async () => {
      render(<Dashboard />);
    });

    expect(screen.getByText("Carro")).toBeInTheDocument();
    expect(screen.getByText("ABC-1234")).toBeInTheDocument();
    expect(screen.getByText("2025-06-07")).toBeInTheDocument();
    expect(screen.getByText("Pago")).toBeInTheDocument();
  });
});
