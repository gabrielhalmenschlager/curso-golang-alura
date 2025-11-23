import React from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/CardJogo.css';

const getCardStyle = (categoria, disponivel) => {
    const neonColors = {
        OURO: '#ffd166',
        PRATA: '#9fb6c7',
        BRONZE: '#cd7f32',
    };

    return {
        '--neon-color': neonColors[categoria] || 'var(--neon-primary)',
        '--status-color': disponivel
            ? 'rgba(0,255,149,0.98)'
            : 'rgba(255,64,64,0.98)',
        cursor: 'pointer',
    };
};

const CardJogo = ({ jogo }) => {
    const navigate = useNavigate();

    if (!jogo) return null;

    const disponivel = Boolean(jogo.disponivel);
    const cardStyle = getCardStyle(jogo.categoria, disponivel);
    const statusText = disponivel ? 'DISPONÍVEL' : 'ALUGADO';

    const handleClick = () => {
        navigate(`/jogos/${jogo.id}`);
    };

    return (
        <article
            className={`jogo-card ${disponivel ? 'is-available' : 'is-rented'}`}
            style={cardStyle}
            onClick={handleClick}
            role="button"
            aria-label={`Abrir detalhes do jogo ${jogo.nome}`}
            tabIndex={0}
            onKeyDown={(e) => e.key === 'Enter' && handleClick()}
        >
            <div className="pulse-indicator" aria-hidden="true"></div>

            {jogo.imagem_url && (
                <img
                    src={jogo.imagem_url}
                    alt={jogo.nome || 'Imagem do jogo'}
                    className="card-image"
                    loading="lazy"
                />
            )}

            <header className="card-header">
                <h3 className="card-title">{jogo.nome}</h3>
                <span className="card-status">{statusText}</span>
            </header>

            <p className="card-description">
                {jogo.descricao || 'Descrição não disponível.'}
            </p>

            <footer className="card-footer">
                <div className="card-category">
                    <span className="category-value">{jogo.categoria}</span>
                </div>

                <div className="card-action-hint">
                    Detalhes...
                </div>
            </footer>
        </article>
    );
};

export default CardJogo;
