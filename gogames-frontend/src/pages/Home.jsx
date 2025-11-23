import React, { useEffect, useState } from 'react';
import { listarJogos } from '../services/jogosService';
import CardJogo from '../components/CardJogo';
import '../styles/Home.css';

const Home = () => {
    const [jogos, setJogos] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        let isMounted = true;

        const fetchJogos = async () => {
            try {
                const data = await listarJogos();
                if (isMounted) setJogos(data);
            } catch (err) {
                console.error('Erro ao carregar jogos:', err);
                if (isMounted) setError('Não foi possível carregar os jogos.');
            } finally {
                if (isMounted) setLoading(false);
            }
        };

        fetchJogos();
        return () => { isMounted = false; };
    }, []);

    if (loading) {
        return (
            <div className="home-container loading-state">
                <h2 className="loading-text">Carregando dados...</h2>
            </div>
        );
    }

    if (error) {
        return (
            <div className="home-container error-state">
                <h2 className="error-text">Erro: {error}</h2>
            </div>
        );
    }

    return (
        <main className="home-container">
            <h1 className="main-title">Catálogo GoGames</h1>

            <section className="jogos-grid" aria-label="Lista de Jogos">
                {jogos.length > 0 ? (
                    jogos.map(jogo => <CardJogo key={jogo.id} jogo={jogo} />)
                ) : (
                    <div className="no-games-message-wrapper">
                        <p className="neon-message">
                            Nenhum jogo registrado. Adicione seu primeiro jogo!
                        </p>
                    </div>
                )}
            </section>
        </main>
    );
};

export default Home;
