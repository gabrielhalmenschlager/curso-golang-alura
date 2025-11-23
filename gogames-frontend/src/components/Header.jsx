import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/Header.css'; 

const Header = () => {
  return (
    <header className="neon-header">
      <Link to="/" className="logo"> 
        <span className="neon-text-pink">GO</span>
        <span className="neon-text-cyan">GAMES</span>
      </Link>
      <nav className="nav-links">
        <Link to="/" className="nav-item neon-link">JOGOS</Link>
        <Link to="/alugueis" className="nav-item neon-link">ALUGUÉIS</Link>
        <Link to="/jogos/novo" className="nav-item neon-link **neon-button**">NOVO JOGO</Link> 
      </nav>
    </header>
  );
};

export default Header;