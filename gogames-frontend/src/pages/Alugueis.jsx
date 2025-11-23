import React, { useState, useEffect } from 'react';
import { listarAlugueis, listarAlugueisAtivos, listarAlugueisInativos, devolverJogo } from '../services/alugueisService';
import { RefreshCw } from 'lucide-react'; 
import '../styles/Alugueis.css';

const Alugueis = () => {
    const [alugueis, setAlugueis] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [filter, setFilter] = useState('TODOS');
    const [message, setMessage] = useState('');

    const fetchAlugueis = async (currentFilter) => {
        setLoading(true);
        setError(null);
        setMessage('');

        let serviceCall = listarAlugueis;
        if (currentFilter === 'ALUGADO') serviceCall = listarAlugueisAtivos;
        if (currentFilter === 'DEVOLVIDO') serviceCall = listarAlugueisInativos;

        try {
            const data = await serviceCall();
            setAlugueis(data);
        } catch (err) {
            console.error('Erro ao carregar aluguéis:', err);
            const errorMessage = err.message.includes('Network Error')
                ? '⚠️ API indisponível. Não foi possível carregar os dados.'
                : '❌ Erro no processamento dos dados. Verifique a API.';
            setError(errorMessage);
            setAlugueis([]);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchAlugueis(filter);
    }, [filter]);

    const handleDevolucao = async (aluguelID) => {
        if (!window.confirm("Confirma a devolução deste aluguel?")) return;

        try {
            await devolverJogo(aluguelID);
            setMessage(`✅ Aluguel #${aluguelID} devolvido com sucesso.`);
            fetchAlugueis(filter);
        } catch (error) {
            console.error('Erro na devolução:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Erro na devolução: ${error.response?.data?.erro || 'Falha na conexão'}`);
        }
    };

    const formatDate = (dateStr) => {
        return dateStr ? new Date(dateStr).toLocaleDateString('pt-BR') : 'N/A';
    }

    return (
        <div className="alugueis-container">
            <h1 className="main-title alugueis-arcade-title">Registro de Aluguéis</h1>

            {message && (
                <p className={`feedback-message ${message.includes('Erro') || message.includes('❌') ? 'error' : 'success'}`}>
                    {message}
                </p>
            )}

            <div className="filter-tabs">
                <button 
                    onClick={() => setFilter('TODOS')}
                    className={`neon-tab ${filter === 'TODOS' ? 'active-cyan' : ''}`}
                >
                    Todos
                </button>
                <button 
                    onClick={() => setFilter('ALUGADO')}
                    className={`neon-tab ${filter === 'ALUGADO' ? 'active-red' : ''}`}
                >
                    Ativos
                </button>
                <button 
                    onClick={() => setFilter('DEVOLVIDO')}
                    className={`neon-tab ${filter === 'DEVOLVIDO' ? 'active-green' : ''}`}
                >
                    Inativos
                </button>
                <button 
                    onClick={() => fetchAlugueis(filter)}
                    className="neon-tab refresh-button"
                    disabled={loading}
                >
                    <RefreshCw size={16} className={loading ? 'spin' : ''} />
                </button>
            </div>

            {loading ? (
                <h2 className="loading-text">Carregando registros...</h2>
            ) : error ? (
                <h2 className="error-text">{error}</h2>
            ) : alugueis.length === 0 ? (
                <p className="neon-message">Nenhum registro encontrado para o filtro "{filter}".</p>
            ) : (
                <div className="alugueis-table-wrapper">
                    <table className="neon-table">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Jogo</th>
                                <th>Data de Aluguel</th>
                                <th>Devolução Prevista</th>
                                <th>Status</th>
                                <th>Ações / Data Real</th>
                            </tr>
                        </thead>
                        <tbody>
                            {alugueis.map(item => (
                                <tr key={item.id} className={item.status === 'ALUGADO' ? 'active-row' : 'inactive-row'}>
                                    <td>{item.id}</td>
                                    <td>{item.jogo.nome}</td>
                                    <td>{formatDate(item.data_aluguel)}</td>
                                    <td>{formatDate(item.data_devolucao_prevista)}</td>
                                    <td>
                                        <span className={`status-pill status-${item.status.toLowerCase()}`}>
                                            {item.status}
                                        </span>
                                    </td>
                                    <td>
                                        {item.status === 'ALUGADO' ? (
                                            <button 
                                                className="neon-button-small btn-devolver"
                                                onClick={() => handleDevolucao(item.id)}
                                            >
                                                Devolver
                                            </button>
                                        ) : (
                                            <span className="devolvido-date">{formatDate(item.data_devolucao_real)}</span>
                                        )}
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
};

export default Alugueis;
