import { useState, useEffect } from 'react';
import { accountsService } from '../services/accounts';
import TransferForm from './TransferForm'; 
import TransactionHistory from './TransactionHistory';

const Dashboard = () => {
  const [accounts, setAccounts] = useState([]);
  const [totalBalance, setTotalBalance] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const [accountsData, balanceData] = await Promise.all([
        accountsService.getAccounts(),
        accountsService.getTotalBalance()
      ]);
      setAccounts(accountsData);
      setTotalBalance(balanceData);
    } catch (err) {
      setError('Ошибка загрузки данных: ' + err.message);
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <div>Загрузка данных о счетах...</div>;
  if (error) return <div style={{ color: 'red' }}>{error}</div>;

  return (
    <div style={{ padding: '20px', maxWidth: '1200px', margin: '0 auto' }}>
      <h2>Мои счета</h2>
      
      <div style={{ 
        marginBottom: '30px', 
        padding: '20px', 
        background: '#e8f5e8', 
        borderRadius: '8px',
        border: '2px solid #4caf50'
      }}>
        <h3 style={{ margin: 0, color: '#2e7d32' }}>
          💰 Общий баланс: {totalBalance.toLocaleString('ru-RU')} ₽
        </h3>
      </div>

      {/* Форма переводов */}
      <TransferForm accounts={accounts} onTransferComplete={loadData} />

      {/* Список счетов */}
      <div style={{ 
        display: 'grid', 
        gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))', 
        gap: '20px',
        marginTop: '30px'
      }}>
        {accounts.map(account => (
          <div key={account.id} style={{ 
            padding: '20px', 
            border: '2px solid #e0e0e0', 
            borderRadius: '12px',
            background: '#fff',
            boxShadow: '0 2px 8px rgba(0,0,0,0.1)'
          }}>
            <h4 style={{ 
              margin: '0 0 15px 0', 
              color: '#1976d2',
              display: 'flex',
              alignItems: 'center',
              gap: '10px'
            }}>
              🏦 {account.bank_name}
            </h4>
            
            <div style={{ lineHeight: '1.6' }}>
              <p>🔢 Счет: {account.account_number}</p>
              <p>💵 Баланс: 
                <strong style={{ fontSize: '1.2em', marginLeft: '8px' }}>
                  {account.balance.toLocaleString('ru-RU')} {account.currency}
                </strong>
              </p>
              <p>📊 Тип: {account.account_type}</p>
            </div>
          </div>
        ))}
      </div>
      <TransactionHistory />
      <div style={{ textAlign: 'center', marginTop: '30px' }}>
        <button 
          onClick={loadData}
          style={{ 
            padding: '12px 24px', 
            background: '#1976d2', 
            color: 'white', 
            border: 'none',
            borderRadius: '6px',
            cursor: 'pointer',
            fontSize: '16px'
          }}
        >
          🔄 Обновить данные
        </button>
      </div>
    </div>
  );
};

export default Dashboard;