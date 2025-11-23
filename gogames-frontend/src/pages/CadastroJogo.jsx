import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { criarJogo } from '../services/jogosService';
import { ArrowLeft } from 'lucide-react'; 
import '../styles/Formulario.css';

const categorias = ["BRONZE", "PRATA", "OURO"];

const CadastroJogo = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        nome: '',
        descricao: '',
        categoria: 'BRONZE',
        imagem_url: '',
    });
    const [loading, setLoading] = useState(false);
    const [message, setMessage] = useState('');

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setMessage('');

        try {
            await criarJogo(formData); 
            setMessage('🎉 Jogo cadastrado com sucesso! Redirecionando...');
            setTimeout(() => navigate('/'), 2000);
        } catch (error) {
            console.error('Erro ao cadastrar:', error.response ? error.response.data.erro : error.message);
            setMessage(`❌ Erro ao cadastrar: ${error.response?.data?.erro || 'Verifique a conexão e os dados.'}`);
        } finally {
            setLoading(false);
        }
    };

    return (
        <main className="form-container container-neon">
            <h1 className="main-title form-arcade-title">Cadastre um novo Jogo</h1>

            {message && (
                <p className={`feedback-message ${message.includes('Erro') || message.includes('❌') ? 'error' : 'success'}`}>
                    {message}
                </p>
            )}

            <form onSubmit={handleSubmit} className="neon-form" style={{ maxWidth: '700px', padding: '50px' }}>
                <label className="neon-label">
                    Nome do Jogo:
                    <input 
                        type="text" 
                        name="nome" 
                        value={formData.nome} 
                        onChange={handleChange} 
                        required 
                        className="neon-input" 
                        disabled={loading}
                        placeholder="Ex: Minecraft"
                    />
                </label>

                <label className="neon-label">
                    Descrição:
                    <textarea 
                        name="descricao" 
                        value={formData.descricao} 
                        onChange={handleChange} 
                        required 
                        className="neon-input textarea-neon" 
                        disabled={loading}
                        placeholder="Breve resumo sobre o gênero e jogabilidade."
                        style={{ minHeight: '150px' }}
                    ></textarea>
                </label>

                <label className="neon-label">
                    URL da Arte:
                    <input 
                        type="url" 
                        name="imagem_url" 
                        value={formData.imagem_url} 
                        onChange={handleChange} 
                        className="neon-input" 
                        disabled={loading}
                        placeholder="https://exemplo.com/capa-do-jogo.jpg"
                    />
                </label>

                <label className="neon-label">
                    Categoria:
                    <select 
                        name="categoria" 
                        value={formData.categoria} 
                        onChange={handleChange} 
                        required 
                        className="neon-input select-neon"
                        disabled={loading}
                    >
                        {categorias.map(cat => (
                            <option key={cat} value={cat}>{cat}</option>
                        ))}
                    </select>
                </label>

                <button 
                    type="submit" 
                    className="neon-button btn-submit-green"
                    disabled={loading}
                >
                    {loading ? 'PROCESSANDO DADOS...' : 'CONFIRMAR CADASTRO'}
                </button>
            </form>

            <button 
                className="neon-button btn-cancel-default" 
                onClick={() => navigate('/')} 
                disabled={loading}
            >
                <ArrowLeft size={16} style={{ marginRight: '8px' }} /> 
                Cancelar e Voltar
            </button>
        </main>
    );
};

export default CadastroJogo;
