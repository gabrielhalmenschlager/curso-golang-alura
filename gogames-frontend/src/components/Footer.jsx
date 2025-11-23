import React from 'react';
import { Github, Linkedin } from 'lucide-react';
import '../styles/Footer.css'; 

const Footer = () => {
  return (
    <footer className="neon-footer">
      <div className="footer-content">
        <p>
          &copy; 2024 **GoGames API**. Todos os direitos reservados.
        </p>
        
        <div className="social-links">
          <span className="developed-by">
            Desenvolvido com <span className="neon-heart">❤️</span> e Go & React
          </span>

          <a 
            href="https://linkedin.com/in/seu-perfil" 
            target="_blank" 
            rel="noopener noreferrer" 
            className="neon-link-footer"
            title="LinkedIn">
            <Linkedin size={20} /> 
          </a>
          <a 
            href="https://github.com/seu-usuario" 
            target="_blank" 
            rel="noopener noreferrer" 
            className="neon-link-footer"
            title="GitHub">
            <Github size={20} />
          </a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;