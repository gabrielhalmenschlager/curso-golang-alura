// src/components/PokemonCard.jsx (Nenhuma mudança necessária neste arquivo)

import { Link } from 'react-router-dom';

function PokemonCard({ id, name, type, image_url, onFavorite, onRemove, onDelete, isFavorited, showAdminActions = false }) { 
    const typeColors = {
      fire: "#f08030",
      water: "#6890f0",
      grass: "#78c850",
      electric: "#f8d030",
      bug: "#a8b820",
      normal: "#a8a878",
      poison: "#a040a0",
      psychic: "#f85888",
      dragon: "#7038f8",
      fairy: "#f0b6bc",
    };
    
    const borderColor = typeColors[type?.toLowerCase()] || "#ccc";
    
    return (
      <div className="pokemon-card" style={{ borderColor }}>
        <img src={image_url} alt={name} className="pokemon-image" />
        <h3>{name}</h3>
        <span
          className="pokemon-type"
          style={{ backgroundColor: typeColors[type?.toLowerCase()] || "#ddd" }}
        >
          {type}
        </span>
        
        <div className="card-actions-wrapper"> 
            
            {isFavorited ? (
              <button className="btn remove-btn" onClick={() => onRemove(id)}>
                Desfavoritar
              </button>
            ) : (
              <button className="btn favorite-btn" onClick={() => onFavorite(id)}>
                Favoritar
              </button>
            )}

            {showAdminActions && (
              <div className="action-buttons">
                <Link to={`/editar/${id}`} className="btn edit-btn">
                  Editar
                </Link>
                <button className="btn delete-btn" onClick={() => onDelete(id)}>
                  Deletar
                </button>
              </div>
            )}
        </div>
      </div>
    );
}

export default PokemonCard;