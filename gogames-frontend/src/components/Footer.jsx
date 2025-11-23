import React from 'react';
import { Github, Linkedin } from 'lucide-react';
import '../styles/Footer.css'; 

const Footer = () => {
  return (
    <footer className="neon-footer">
      <div className="footer-content">
        <p>
          &copy; 2025 GoGames API. Todos os direitos reservados.
        </p>
        
        <div className="social-links">
          <span className="developed-by">
            Desenvolvido com Go e React
          </span>

          <a 
            href="https://www.linkedin.com/in/gabriel-halmenschlager/" 
            target="_blank" 
            rel="noopener noreferrer" 
            className="neon-link-footer"
            title="LinkedIn">
            <Linkedin size={20} /> 
          </a>
          <a 
            href="https://github.com/gabrielhalmenschlager" 
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