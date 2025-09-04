import { api } from './api';

export const accountsService = {
  // Получить все счета
  getAccounts: async () => {
    const response = await api.get('/accounts');
    return response.data.accounts;
  },
  
  // Получить общий баланс
  getTotalBalance: async () => {
    const response = await api.get('/balance');
    return response.data.total_balance;
  },
  
  // Перевод между счетами
  transfer: async (fromAccount, toAccount, amount, currency = 'RUB') => {
    const response = await api.post('/transfer/internal', {
      from_account: fromAccount,
      to_account: toAccount,
      amount: amount,
      currency: currency
    });
    return response.data;
  }, 
  // Получить историю транзакций
  getTransactionHistory: async () => {
    const response = await api.get('/transfer/history');
    return response.data.transactions;
  }
};