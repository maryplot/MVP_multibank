import { useState, useEffect } from 'react';
import { accountsService } from '../services/accounts';

const TransactionHistory = () => {
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    loadHistory();
  }, []);

  const loadHistory = async () => {
    try {
      const data = await accountsService.getTransactionHistory();
      setTransactions(data);
    } catch (err) {
      setError('Ошибка загрузки истории операций');
      console.error('History load error:', err);
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString('ru-RU');
  };

  if (loading) return <div>Загрузка истории операций...</div>;
  if (error) return <div style={{ color: 'red' }}>{error}</div>;

  return (
    <div style={{ 
      marginTop: '40px',
      padding: '20px',
      background: '#fff',
      borderRadius: '10px',
      border: '2px solid #e0e0e0'
    }}>
      <h3 style={{ 
        margin: '0 0 20px 0',
        display: 'flex',
        alignItems: 'center',
        gap: '10px'
      }}>
        📋 История операций
      </h3>

      {transactions.length === 0 ? (
        <div style={{ textAlign: 'center', color: '#666', padding: '40px' }}>
          📭 Операций пока нет
        </div>
      ) : (
        <div style={{ display: 'grid', gap: '12px' }}>
          {transactions.map((transaction) => (
            <div
              key={transaction.id}
              style={{
                padding: '16px',
                border: '1px solid #eee',
                borderRadius: '8px',
                background: '#fafafa'
              }}
            >
              <div style={{ 
                display: 'flex', 
                justifyContent: 'space-between',
                alignItems: 'center',
                marginBottom: '8px'
              }}>
                <span style={{ fontWeight: 'bold' }}>
                  #{transaction.id}
                </span>
                <span style={{ 
                  color: transaction.status === 'completed' ? '#4caf50' : '#ff9800',
                  fontSize: '14px'
                }}>
                  {transaction.status === 'completed' ? '✅ Выполнено' : '⏳ В обработке'}
                </span>
              </div>

              <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '10px' }}>
                <div>
                  <div style={{ fontSize: '12px', color: '#666' }}>Откуда:</div>
                  <div>{transaction.from_account}</div>
                </div>
                <div>
                  <div style={{ fontSize: '12px', color: '#666' }}>Куда:</div>
                  <div>{transaction.to_account}</div>
                </div>
              </div>

              <div style={{ 
                marginTop: '12px',
                padding: '12px',
                background: '#e8f5e8',
                borderRadius: '6px',
                textAlign: 'center'
              }}>
                <div style={{ fontSize: '18px', fontWeight: 'bold', color: '#2e7d32' }}>
                  {transaction.amount.toLocaleString('ru-RU')} {transaction.currency}
                </div>
              </div>

              <div style={{ 
                marginTop: '8px',
                fontSize: '12px',
                color: '#666',
                textAlign: 'center'
              }}>
                {formatDate(transaction.created_at)}
              </div>
            </div>
          ))}
        </div>
      )}

      <button
        onClick={loadHistory}
        style={{
          marginTop: '20px',
          padding: '10px 20px',
          background: '#1976d2',
          color: 'white',
          border: 'none',
          borderRadius: '6px',
          cursor: 'pointer'
        }}
      >
        🔄 Обновить историю
      </button>
    </div>
  );
};

export default TransactionHistory;