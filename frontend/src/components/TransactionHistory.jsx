  import { useState, useEffect } from 'react';
import { accountsService } from '../services/accounts';
import './TransactionHistory.css';

const TransactionHistory = () => {
  console.log('TransactionHistory component rendered');
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  // Map account IDs to user-friendly names
  const accountNames = {
    'tinkoff_123': 'Тинькофф (****1234)',
    'alfa_111': 'Альфа-Банк (****1111)',
    'sber_789': 'Сбербанк (****9012)',
    'sber_345': 'Сбербанк (****3456)'
  };

  // Get user-friendly name for account ID
  const getAccountName = (accountId) => {
    return accountNames[accountId] || accountId;
  };

  useEffect(() => {
    console.log('TransactionHistory useEffect triggered');
    loadHistory();
  }, []);

  const loadHistory = async () => {
    try {
      console.log('Loading transaction history...');
      const data = await accountsService.getTransactionHistory();
      console.log('Transaction history loaded:', data);
      // Check if data is an object with a transactions property or just an array
      if (data && data.transactions) {
        setTransactions(data.transactions);
      } else if (Array.isArray(data)) {
        setTransactions(data);
      } else {
        setTransactions([]);
      }
    } catch (err) {
      console.error('Error loading transaction history:', err);
      setError('Ошибка загрузки истории операций');
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString('ru-RU');
  };

  if (loading) return <div className="transaction-history-container">Загрузка истории операций...</div>;
  if (error) return <div className="transaction-history-container" style={{ color: 'red' }}>{error}</div>;

  return (
    <div className="transaction-history-container">
      <h3 className="transaction-history-title">
        📋 История операций
      </h3>

      {transactions.length === 0 ? (
        <div className="transaction-history-empty">
          📭 Операций пока нет
        </div>
      ) : (
        <div className="transaction-list">
          {transactions.map((transaction) => (
            <div
              key={transaction.id}
              className="transaction-item"
            >
              <div className="transaction-header">
                <span className="transaction-type">
                  Перевод средств
                </span>
                <span className={`transaction-status ${transaction.status === 'completed' ? 'completed' : 'pending'}`}>
                  {transaction.status === 'completed' ? '✅ Выполнено' : '⏳ В обработке'}
                </span>
              </div>

              <div className="transaction-details">
                <div>
                  <div className="transaction-detail-label">Откуда:</div>
                  <div className="transaction-detail-value">{getAccountName(transaction.from_account)}</div>
                </div>
                <div>
                  <div className="transaction-detail-label">Куда:</div>
                  <div className="transaction-detail-value">{getAccountName(transaction.to_account)}</div>
                </div>
              </div>

              <div className="transaction-amount-container">
                <div className="transaction-amount">
                  {transaction.amount.toLocaleString('ru-RU')} {transaction.currency}
                </div>
              </div>

              <div className="transaction-date">
                {formatDate(transaction.created_at)}
              </div>
            </div>
          ))}
        </div>
      )}

      <button
        onClick={loadHistory}
        className="transaction-refresh-button"
      >
        🔄 Обновить историю
      </button>
    </div>
  );
};

export default TransactionHistory;