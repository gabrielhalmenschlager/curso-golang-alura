import React, { useEffect, useState, useMemo } from "react";
import PokemonCard from "./PokemonCard";
import DeleteModal from "./DeleteModal"; // <-- 1. Importar o componente Modal

function PokemonList() {
    const [pokemons, setPokemons] = useState([]);
    const [favoritos, setFavoritos] = useState([]);

    // 2. Novos Estados para gerenciar o Modal
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [pokemonToDelete, setPokemonToDelete] = useState(null); // ID do Pokémon selecionado

    useEffect(() => {
        const fetchPokemons = fetch("http://localhost:8000/pokemon").then((res) => res.json());
        const fetchFavoritos = fetch("http://localhost:8000/favoritos").then((res) => res.json());

        Promise.all([fetchPokemons, fetchFavoritos])
          .then(([pokemonData, favoritoData]) => {
            setPokemons(pokemonData);
            setFavoritos(favoritoData);
          })
          .catch(error => console.error("Erro ao buscar dados:", error));
    }, []);

    const favoritoIds = useMemo(() =>
        new Set(favoritos.map((fav) => fav.pokemon.id))
    , [favoritos]);

    const adicionarFavorito = (id) => {
        fetch(`http://localhost:8000/favoritos/${id}`, {
            method: "POST",
        }).then(() => {
            const pokemonAdicionado = pokemons.find(p => p.id === id);
            if (pokemonAdicionado) {
                setFavoritos([...favoritos, { id: Math.random(), pokemon: pokemonAdicionado }]);
            }
        });
    };

    const removerFavorito = (id) => {
        fetch(`http://localhost:8000/favoritos/${id}`, {
            method: "DELETE",
        }).then(() => {
            setFavoritos(favoritos.filter((fav) => fav.pokemon.id !== id));
        });
    };
    
    // 3A. Abre o modal em vez de mostrar o alerta
    const openDeleteModal = (id) => {
        setPokemonToDelete(id); // Guarda o ID para deletar
        setIsModalOpen(true);   // Abre o modal
    };

    // 3B. Função para fechar o modal
    const closeDeleteModal = () => {
        setIsModalOpen(false);
        setPokemonToDelete(null); // Limpa o ID
    };

    // 3C. Função que executa a exclusão APÓS a confirmação no modal
    const handleDeleteConfirm = () => {
        if (!pokemonToDelete) return; // Se não houver ID, sai

        // Remove a lógica de window.confirm
        fetch(`http://localhost:8000/pokemon/${pokemonToDelete}`, {
            method: "DELETE",
        }).then(() => {
            setPokemons(pokemons.filter((poke) => poke.id !== pokemonToDelete));
            setFavoritos(favoritos.filter((fav) => fav.pokemon.id !== pokemonToDelete));
            closeDeleteModal(); // Fecha o modal após o sucesso
        }).catch(error => {
            console.error("Erro ao deletar Pokémon:", error);
            alert("Falha ao deletar Pokémon.");
            closeDeleteModal(); // Fecha mesmo em caso de erro
        });
    };

    // Encontra as informações do Pokémon para exibir o nome no modal
    const pokemonInfoToDelete = pokemons.find(p => p.id === pokemonToDelete);

    return (
        <div>
            <h2 className="page-title">Pokémon Encontrados</h2>
            <div className="list-grid">
                {pokemons.map((poke) => {
                    const isFavorited = favoritoIds.has(poke.id);

                    return (
                        <PokemonCard
                            key={poke.id}
                            id={poke.id}
                            name={poke.name}
                            type={poke.type}
                            image_url={poke.image_url}
                            
                            isFavorited={isFavorited}
                            onFavorite={adicionarFavorito}
                            onRemove={removerFavorito}
                            // MUDANÇA: Passa a função que abre o modal
                            onDelete={openDeleteModal}
                            showAdminActions={true}
                        />
                    );
                })}
            </div>

            {/* 4. Renderização do Modal */}
            <DeleteModal
                isOpen={isModalOpen}
                onClose={closeDeleteModal}
                onConfirm={handleDeleteConfirm}
                name={pokemonInfoToDelete?.name || 'este Pokémon'}
            />
        </div>
    );
}

export default PokemonList;