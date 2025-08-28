import { useState, useEffect } from 'react';
import { accountsService } from '../services/accounts';

const TransferForm = ({ accounts, onTransferComplete }) => {
  const [fromAccount, setFromAccount] = useState('');
  const [toAccount, setToAccount] = useState('');
  const [amount, setAmount] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage('');

    try {
      const result = await accountsService.transfer(fromAccount, toAccount, parseFloat(amount));
      setMessage(`✅ Перевод успешно выполнен! ID: ${result.transaction.id}`);
      setFromAccount('');
      setToAccount('');
      setAmount('');
      if (onTransferComplete) onTransferComplete();
    } catch (error) {
      setMessage('❌ Ошибка перевода: ' + (error.response?.data?.error || error.message));
    } finally {
      setLoading(false);
    }
  };

  const isSameAccount = fromAccount && toAccount && fromAccount === toAccount;
  const isFormValid = fromAccount && toAccount && amount > 0 && !isSameAccount;

  return (
    <div style={{ 
      padding: '20px', 
      border: '2px solid #e0e0e0', 
      borderRadius: '10px',
      margin: '20px 0',
      background: '#fff'
    }}>
      <h3>💸 Перевод между своими счетами</h3>
      
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: '15px' }}>
          <label>Откуда:</label>
          <select 
            value={fromAccount} 
            onChange={(e) => setFromAccount(e.target.value)}
            style={{ width: '100%', padding: '8px' }}
            required
          >
            <option value="">Выберите счет</option>
            {accounts.map(account => (
              <option key={account.id + '_from'} value={account.id}>
                {account.bank_name} - {account.account_number} ({account.balance} {account.currency})
              </option>
            ))}
          </select>
        </div>

        <div style={{ marginBottom: '15px' }}>
          <label>Куда:</label>
          <select 
            value={toAccount} 
            onChange={(e) => setToAccount(e.target.value)}
            style={{ width: '100%', padding: '8px' }}
            required
          >
            <option value="">Выберите счет</option>
            {accounts.map(account => (
              <option key={account.id + '_to'} value={account.id}>
                {account.bank_name} - {account.account_number} ({account.balance} {account.currency})
              </option>
            ))}
          </select>
        </div>

        <div style={{ marginBottom: '15px' }}>
          <label>Сумма:</label>
          <input
            type="number"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            placeholder="Введите сумму"
            style={{ width: '100%', padding: '8px' }}
            min="0.01"
            step="0.01"
            required
          />
        </div>

        {isSameAccount && (
          <div style={{ color: 'red', marginBottom: '15px' }}>
            ❌ Нельзя переводить на тот же счет!
          </div>
        )}

        {message && (
          <div style={{ 
            marginBottom: '15px', 
            color: message.includes('✅') ? 'green' : 'red',
            padding: '10px',
            background: message.includes('✅') ? '#e8f5e8' : '#ffe6e6',
            borderRadius: '5px'
          }}>
            {message}
          </div>
        )}

        <button 
          type="submit" 
          disabled={!isFormValid || loading}
          style={{
            width: '100%',
            padding: '12px',
            background: isFormValid ? '#1976d2' : '#ccc',
            color: 'white',
            border: 'none',
            borderRadius: '5px',
            cursor: isFormValid ? 'pointer' : 'not-allowed'
          }}
        >
          {loading ? '⏳ Выполнение...' : '🚀 Выполнить перевод'}
        </button>
      </form>
    </div>
  );
};

export default TransferForm;