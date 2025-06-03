import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Dashboard from "./components/Dashboard";
import CadClient from "./components/CadClient";
import Mensalistas from "./components/Mensalistas";
import Pagamento from "./components/Pagamento";
import Pagamentos from "./components/Pagamentos";
import BuscarPlaca from "./components/BuscarPlaca";
import Navbar from "./components/Navbar";

function App() {
  return (
    <BrowserRouter basename="/observer-pi3/">
      <Navbar/>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/cadclient" element={<CadClient />} />
        <Route path="/mensalistas" element={<Mensalistas />} />
        <Route path="/pagamentos" element={<Pagamentos />} />
        <Route path="/pagamento" element={<Pagamento />} />
        <Route path="/buscaplaca" element={<BuscarPlaca />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
