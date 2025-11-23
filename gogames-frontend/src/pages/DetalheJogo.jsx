import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { buscarJogoPorID, atualizarJogo, deletarJogo } from '../services/jogosService';
import { alugarJogo, buscarAluguelAtivoPorJogoID } from '../services/alugueisService'; 
import '../styles/DetalheJogo.css';
import { ArrowLeft } from 'lucide-react';

const categorias = ["BRONZE", "PRATA", "OURO"];

const formatDate = (dateStr) => dateStr ? new Date(dateStr).toLocaleDateString('pt-BR') : 'N/A';

const DetalheJogo = () => {
    const { id } = useParams();
    const navigate = useNavigate();

    const [jogo, setJogo] = useState(null);
    const [aluguelAtivo, setAluguelAtivo] = useState(null);
    const [isEditing, setIsEditing] = useState(false);
    const [formData, setFormData] = useState({});
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState('');

    const fetchJogo = async () => {
        setLoading(true);
        setAluguelAtivo(null);
        setMessage('');

        try {
            const jogoData = await buscarJogoPorID(id);
            setJogo(jogoData);
            setFormData({
                nome: jogoData.nome,
                descricao: jogoData.descricao,
                categoria: jogoData.categoria,
            });

            if (!jogoData.disponivel) {
                const aluguelData = await buscarAluguelAtivoPorJogoID(id);
                setAluguelAtivo(aluguelData);
            }
        } catch (error) {
            console.error('Erro ao buscar jogo:', error);
            setMessage('Jogo não encontrado ou erro de conexão.');
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchJogo();
    }, [id]);

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleUpdate = async (e) => {
        e.preventDefault();
        try {
            const updatedJogo = await atualizarJogo(id, formData);
            setJogo(updatedJogo);
            setIsEditing(false);
            setMessage('✨ Jogo atualizado com sucesso!');
        } catch (error) {
            console.error('Erro ao atualizar:', error);
            setMessage('🚨 Erro ao atualizar o jogo.');
        }
    };

    const handleAlugar = async () => {
        if (!window.confirm(`Deseja alugar o jogo "${jogo.nome}"?`)) return;

        try {
            await alugarJogo(jogo.id);
            setMessage('🎉 Jogo alugado com sucesso! Recarregando status...');
            await fetchJogo();
        } catch (error) {
            console.error('Erro ao alugar:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Erro ao alugar: ${error.response?.data?.erro || 'Serviço indisponível'}`);
        }
    };

    const handleDeletar = async () => {
        if (!window.confirm(`Confirma a exclusão de "${jogo.nome}"?`)) return;

        try {
            await deletarJogo(id);
            setMessage('🗑️ Jogo excluído com sucesso! Redirecionando...');
            setTimeout(() => navigate('/'), 2000);
        } catch (error) {
            console.error('Erro ao deletar:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Não foi possível excluir: ${error.response?.data?.erro || 'Verifique se o jogo está alugado.'}`);
        }
    };

    if (loading) return <h2 className="loading-text">BUSCANDO DADOS...</h2>;
    if (!jogo) return <h2 className="error-text">{message}</h2>;

    const imageBorderClass = jogo.disponivel ? 'status-available-border' : 'status-rented-border';

    return (
        <main className="detail-page-container container-neon">
            <button className="neon-button back-button" onClick={() => navigate('/')}>
                <ArrowLeft size={16} style={{ marginRight: '8px' }} />
                Voltar ao Catálogo
            </button>

            <h1 className="detail-title">{jogo.nome}</h1>

            {message && (
                <p className={`feedback-message ${message.includes('Erro') || message.includes('❌') ? 'error' : 'success'}`}>
                    {message}
                </p>
            )}

            {!isEditing ? (
                <div className="game-detail-content">
                    <div className="game-image-main">
                        <div className={`game-image-wrapper ${imageBorderClass}`}>
                            {jogo.imagem_url ? (
                                <img src={jogo.imagem_url} alt={`Arte de ${jogo.nome}`} className="game-main-image" />
                            ) : (
                                <div className="game-placeholder-image">
                                    <span>[ARTE NÃO ENCONTRADA]</span>
                                </div>
                            )}
                        </div>
                    </div>

                    <div className="game-info-actions neon-data-panel">
                        <div className="data-row status-info">
                            <h2 className="status-label">Status Atual:</h2>
                            <span className={`status-tag status-${jogo.disponivel ? 'available' : 'rented'}`}>
                                {jogo.disponivel ? 'Disponível (Online)' : 'Alugado (Offline)'}
                            </span>
                        </div>

                        {!jogo.disponivel && aluguelAtivo && (
                            <div className="data-row rental-details-info">
                                <h3 className="rental-details-title">🕹️ Aluguel Ativo</h3>
                                <div className="info-grid rental-grid">
                                    <p className="label-data">ID Aluguel:</p>
                                    <span className="value-data">{aluguelAtivo.id}</span>

                                    <p className="label-data">Alugado em:</p>
                                    <span className="value-data">{formatDate(aluguelAtivo.data_aluguel)}</span>

                                    <p className="label-data">Devolução Prevista:</p>
                                    <span className="value-data status-alert-date">
                                        {formatDate(aluguelAtivo.data_devolucao_prevista)}
                                    </span>

                                    {aluguelAtivo.cliente && (
                                        <>
                                            <p className="label-data">Cliente:</p>
                                            <span className="value-data">{aluguelAtivo.cliente.nome || aluguelAtivo.cliente.id}</span>
                                        </>
                                    )}
                                </div>
                                <hr className="neon-divider"/>
                            </div>
                        )}

                        <div className="data-row info-grid">
                            <p className="label-data">Descrição:</p>
                            <p className="value-data description-text">{jogo.descricao || 'Nenhuma descrição fornecida.'}</p>

                            <p className="label-data">Categoria:</p>
                            <span className={`category-tag category-${jogo.categoria}`}>{jogo.categoria}</span>

                            <p className="label-data">Data de Cadastro:</p>
                            <span className="value-data">{formatDate(jogo.data_cadastro)}</span>
                        </div>

                        <div className="action-buttons">
                            <button
                                className={`neon-button ${jogo.disponivel ? 'btn-green' : 'btn-disabled'}`}
                                onClick={handleAlugar}
                                disabled={!jogo.disponivel}
                            >
                                {jogo.disponivel ? 'Alugar' : 'Indisponível'}
                            </button>

                            <button className="neon-button btn-cyan" onClick={() => setIsEditing(true)}>
                                Editar Registro
                            </button>
                            <button
                                className="neon-button btn-red"
                                onClick={handleDeletar}
                                disabled={!jogo.disponivel}
                            >
                                Excluir Registro
                            </button>
                        </div>
                    </div>
                </div>
            ) : (
                <form onSubmit={handleUpdate} className="edit-form neon-data-panel">
                    <h2>Editar Dados do Jogo</h2>
                    <label className="neon-label">
                        Nome do Jogo:
                        <input type="text" name="nome" value={formData.nome || ''} onChange={handleChange} required className="neon-input" />
                    </label>
                    <label className="neon-label">
                        Descrição:
                        <textarea name="descricao" value={formData.descricao || ''} onChange={handleChange} required className="neon-input"></textarea>
                    </label>
                    <label className="neon-label">
                        Categoria:
                        <select name="categoria" value={formData.categoria || ''} onChange={handleChange} required className="neon-input">
                            {categorias.map(cat => <option key={cat} value={cat}>{cat}</option>)}
                        </select>
                    </label>
                    <div className="action-buttons edit-mode-buttons">
                        <button type="submit" className="neon-button btn-green">Salvar Alterações</button>
                        <button type="button" className="neon-button btn-red" onClick={() => setIsEditing(false)}>Cancelar Edição</button>
                    </div>
                </form>
            )}
        </main>
    );
};

export default DetalheJogo;
