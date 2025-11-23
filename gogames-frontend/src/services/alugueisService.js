// src/services/alugueisService.js
import api from './api';

// GET /alugueis
export const listarAlugueis = async () => {
    const response = await api.get('/alugueis');
    return response.data;
};

// POST /alugueis/alugar/:id (O ID aqui é o JogoID)
export const alugarJogo = async (jogoId) => {
    // Sua API espera o ID na URL e não um corpo JSON, conforme o controlador:
    // r.POST("/alugueis/alugar/:id", controllers.AlugarJogo)
    const response = await api.post(`/alugueis/alugar/${jogoId}`);
    return response.data;
};

// PUT /alugueis/devolver/:id (O ID aqui é o AluguelID, que será buscado na tela de Aluguéis)
export const devolverJogo = async (aluguelId) => {
    const response = await api.put(`/alugueis/devolver/${aluguelId}`);
    return response.data;
};

// GET /alugueis/ativos
export const listarAlugueisAtivos = async () => {
    const response = await api.get('/alugueis/ativos');
    return response.data;
};

// GET /alugueis/inativos
export const listarAlugueisInativos = async () => {
    const response = await api.get('/alugueis/inativos');
    return response.data;
};

export const buscarAluguelAtivoPorJogoID = async (jogoID) => {
    try {
        const alugueisAtivos = await listarAlugueisAtivos();
        
        // Converte jogoID para número para garantir a comparação correta
        const idNumerico = Number(jogoID);
        
        // Filtra a lista de todos os ativos para encontrar o que corresponde ao jogo_id
        const aluguelDoJogo = alugueisAtivos.find(aluguel => aluguel.jogo_id === idNumerico);
        
        // Retorna o objeto de aluguel ou undefined (que será tratado como null no componente)
        return aluguelDoJogo || null;
    } catch (error) {
        console.error("Erro ao buscar aluguel ativo por Jogo ID:", error);
        return null;
    }
};