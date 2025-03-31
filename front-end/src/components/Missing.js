import { Link } from "react-router-dom";

const Missing = () => {
  return (
    <article style={{ padding: "100px" }}>
      <h1>Oops!</h1>
      <p>Página Não Encontrada!</p>
      <div className="flexGrow">
        <Link to="/">Voltar para a página inicial</Link>
      </div>
    </article>
  );
};

export default Missing;
