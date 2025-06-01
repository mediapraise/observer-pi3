import { Link } from "react-router-dom";

function Navbar() {
  return (
    <nav className="navbar">
      <ul className="nav-list">
        <li>
          <Link to="/dashboard">Dashboard</Link>
        </li>
        <li>
          <Link to="/mensalistas">Cadastro</Link>
        </li>
        <li>
          <Link to="/pagamentos">Financeiro</Link>
        </li>
        <li>
          <Link to="/buscaplaca">Histórico</Link>
        </li>
        <li>
          <Link to="/pagamento">Pagamento</Link>
        </li>
      </ul>
    </nav>
  );
}

export default Navbar;
