import React, { useState, useEffect } from 'react';
import './ExpensesPage.css';
import alfaIcon from '../assets/icon_alfa.png';
import sberIcon from '../assets/icon_sber.png';
import tbankIcon from '../assets/icon_tbank.png';

const ExpensesPage = ({ onBack }) => {
  const [activeTab, setActiveTab] = useState('banks');
  const [showHistoryModal, setShowHistoryModal] = useState(false);
  const [historyType, setHistoryType] = useState('');
  const [historyData, setHistoryData] = useState([]);
  
  const getBankIcon = (bankName) => {
    switch (bankName.toLowerCase()) {
      case 'альфа-банк':
        return { icon: alfaIcon, color: '#EF3124', isImage: true };
      case 'тинькофф':
        return { icon: tbankIcon, color: '#FFDD2D', isImage: true };
      case 'сбербанк':
        return { icon: sberIcon, color: '#21A038', isImage: true };
      default:
        return { icon: bankName.charAt(0), color: '#666', isImage: false };
    }
  };

  const [expenses, setExpenses] = useState({
    byBanks: [
      { bank: 'Альфа-Банк', amount: 4100, color: '#EF3124' },
      { bank: 'Сбербанк', amount: 6200, color: '#21A038' },
      { bank: 'Тинькофф', amount: 3200, color: '#FFDD2D' }
    ],
    byCategories: [
      { category: 'Продукты', amount: 6000, color: '#3b82f6', icon: '🛒' },
      { category: 'Транспорт', amount: 2500, color: '#8b5cf6', icon: '🚗' },
      { category: 'Развлечения', amount: 1800, color: '#f59e0b', icon: '🎬' },
      { category: 'Кафе', amount: 3200, color: '#ef4444', icon: '☕' }
    ]
  });

  const [limits, setLimits] = useState([
    { name: 'Продукты', current: 6000, limit: 8000, color: '#3b82f6' },
    { name: 'Транспорт', current: 2500, limit: 3000, color: '#8b5cf6' },
    { name: 'Развлечения', current: 1800, limit: 2000, color: '#f59e0b' },
    { name: 'Кафе', current: 3200, limit: 2500, color: '#ef4444' }
  ]);

  // История расходов
  const expenseHistory = {
    banks: {
      'Альфа-Банк': [
        { date: '18.09.2024', description: 'Магазин "Пятёрочка"', amount: -850, category: 'Продукты' },
        { date: '17.09.2024', description: 'Кафе "Шоколадница"', amount: -420, category: 'Кафе' },
        { date: '16.09.2024', description: 'Заправка BP', amount: -1200, category: 'Транспорт' },
        { date: '15.09.2024', description: 'Кинотеатр "Синема Парк"', amount: -680, category: 'Развлечения' },
        { date: '14.09.2024', description: 'Супермаркет "Ашан"', amount: -950, category: 'Продукты' }
      ],
      'Сбербанк': [
        { date: '18.09.2024', description: 'Ресторан "Тануки"', amount: -1200, category: 'Кафе' },
        { date: '17.09.2024', description: 'Метро (пополнение)', amount: -500, category: 'Транспорт' },
        { date: '16.09.2024', description: 'Магазин "Перекрёсток"', amount: -1800, category: 'Продукты' },
        { date: '15.09.2024', description: 'Боулинг "Космик"', amount: -1200, category: 'Развлечения' },
        { date: '14.09.2024', description: 'Кофейня "Старбакс"', amount: -500, category: 'Кафе' }
      ],
      'Тинькофф': [
        { date: '18.09.2024', description: 'Такси Яндекс', amount: -300, category: 'Транспорт' },
        { date: '17.09.2024', description: 'Доставка еды', amount: -800, category: 'Кафе' },
        { date: '16.09.2024', description: 'Магазин "Лента"', amount: -1100, category: 'Продукты' },
        { date: '15.09.2024', description: 'Кинотеатр IMAX', amount: -900, category: 'Развлечения' },
        { date: '14.09.2024', description: 'Заправка Лукойл', amount: -800, category: 'Транспорт' }
      ]
    },
    categories: {
      'Продукты': [
        { date: '18.09.2024', description: 'Магазин "Пятёрочка"', amount: -850, bank: 'Альфа-Банк' },
        { date: '16.09.2024', description: 'Супермаркет "Ашан"', amount: -950, bank: 'Альфа-Банк' },
        { date: '16.09.2024', description: 'Магазин "Перекрёсток"', amount: -1800, bank: 'Сбербанк' },
        { date: '16.09.2024', description: 'Магазин "Лента"', amount: -1100, bank: 'Тинькофф' },
        { date: '15.09.2024', description: 'Фермерский рынок', amount: -1300, bank: 'Сбербанк' }
      ],
      'Транспорт': [
        { date: '18.09.2024', description: 'Такси Яндекс', amount: -300, bank: 'Тинькофф' },
        { date: '17.09.2024', description: 'Метро (пополнение)', amount: -500, bank: 'Сбербанк' },
        { date: '16.09.2024', description: 'Заправка BP', amount: -1200, bank: 'Альфа-Банк' },
        { date: '14.09.2024', description: 'Заправка Лукойл', amount: -800, bank: 'Тинькофф' },
        { date: '13.09.2024', description: 'Парковка в центре', amount: -200, bank: 'Сбербанк' }
      ],
      'Развлечения': [
        { date: '15.09.2024', description: 'Кинотеатр "Синема Парк"', amount: -680, bank: 'Альфа-Банк' },
        { date: '15.09.2024', description: 'Боулинг "Космик"', amount: -1200, bank: 'Сбербанк' },
        { date: '15.09.2024', description: 'Кинотеатр IMAX', amount: -900, bank: 'Тинькофф' },
        { date: '14.09.2024', description: 'Театр им. Пушкина', amount: -1500, bank: 'Сбербанк' },
        { date: '13.09.2024', description: 'Концерт в Крокус Сити', amount: -2200, bank: 'Альфа-Банк' }
      ],
      'Кафе': [
        { date: '18.09.2024', description: 'Ресторан "Тануки"', amount: -1200, bank: 'Сбербанк' },
        { date: '17.09.2024', description: 'Кафе "Шоколадница"', amount: -420, bank: 'Альфа-Банк' },
        { date: '17.09.2024', description: 'Доставка еды', amount: -800, bank: 'Тинькофф' },
        { date: '14.09.2024', description: 'Кофейня "Старбакс"', amount: -500, bank: 'Сбербанк' },
        { date: '13.09.2024', description: 'Ресторан "Белуга"', amount: -2800, bank: 'Альфа-Банк' }
      ]
    }
  };

  const showHistory = (type, item = null) => {
    let data = [];
    let title = '';
    
    if (type === 'pie') {
      // Показать общую историю
      data = Object.values(expenseHistory[activeTab === 'banks' ? 'banks' : 'categories']).flat();
      title = activeTab === 'banks' ? 'История по всем банкам' : 'История по всем категориям';
    } else if (type === 'bank') {
      data = expenseHistory.banks[item] || [];
      title = `История расходов: ${item}`;
    } else if (type === 'category') {
      data = expenseHistory.categories[item] || [];
      title = `История расходов: ${item}`;
    }
    
    // Сортируем по дате (новые сверху)
    data.sort((a, b) => new Date(b.date.split('.').reverse().join('-')) - new Date(a.date.split('.').reverse().join('-')));
    
    setHistoryData(data);
    setHistoryType(title);
    setShowHistoryModal(true);
  };

  const totalExpenses = activeTab === 'banks' 
    ? expenses.byBanks.reduce((sum, item) => sum + item.amount, 0)
    : expenses.byCategories.reduce((sum, item) => sum + item.amount, 0);

  const maxAmount = activeTab === 'banks'
    ? Math.max(...expenses.byBanks.map(item => item.amount))
    : Math.max(...expenses.byCategories.map(item => item.amount));

  return (
    <div className="expenses-page">
      <div className="expenses-header">
        <button className="back-button" onClick={onBack}>←</button>
        <h1>Мои расходы</h1>
      </div>

      <div className="tab-switcher">
        <button
          className={`tab-button ${activeTab === 'banks' ? 'active' : ''}`}
          onClick={() => setActiveTab('banks')}
        >
          по банкам
        </button>
        <span className="tab-separator">|</span>
        <button
          className={`tab-button ${activeTab === 'categories' ? 'active' : ''}`}
          onClick={() => setActiveTab('categories')}
        >
          по категориям
        </button>
      </div>

      <div className="total-amount">
        {totalExpenses.toLocaleString('ru-RU')} ₽
      </div>

      <div className="chart-container">
        {activeTab === 'banks' ? (
          <div className="bar-chart">
            {expenses.byBanks.map((item, index) => {
              const bankInfo = getBankIcon(item.bank);
              return (
                <div key={index} className="bar-item clickable" onClick={() => showHistory('bank', item.bank)}>
                  <div className="bar-amount">{item.amount.toLocaleString('ru-RU')} ₽</div>
                  <div
                    className="bar"
                    style={{
                      height: `${(item.amount / maxAmount) * 200}px`,
                      backgroundColor: item.color
                    }}
                  />
                  <div className="bar-logo" style={{ backgroundColor: item.color }}>
                    {bankInfo.isImage ? (
                      <img src={bankInfo.icon} alt={item.bank} className="bank-icon-img" />
                    ) : (
                      bankInfo.icon
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        ) : (
          <div className="bar-chart">
            {expenses.byCategories.map((item, index) => (
              <div key={index} className="bar-item clickable" onClick={() => showHistory('category', item.category)}>
                <div className="bar-amount">{item.amount.toLocaleString('ru-RU')} ₽</div>
                <div
                  className="bar"
                  style={{
                    height: `${(item.amount / maxAmount) * 200}px`,
                    backgroundColor: item.color
                  }}
                />
                <div className="category-icon">{item.icon}</div>
              </div>
            ))}
          </div>
        )}
      </div>

      <div className="pie-chart-section">
        <div className="pie-chart clickable" onClick={() => showHistory('pie')}>
          <div className="pie-chart-inner">
            <span className="pie-total">{totalExpenses.toLocaleString('ru-RU')} ₽</span>
          </div>
        </div>
      </div>

      <div className="limits-section">
        <h3>Лимиты</h3>
        <div className="limits-grid">
          {limits.map((limit, index) => (
            <div key={index} className="limit-item">
              <div className="limit-name">{limit.name}</div>
              <div className="limit-bar">
                <div 
                  className="limit-progress" 
                  style={{ 
                    width: `${Math.min((limit.current / limit.limit) * 100, 100)}%`,
                    backgroundColor: limit.current > limit.limit ? '#ef4444' : limit.color
                  }}
                />
              </div>
              <div className="limit-text">
                {limit.current.toLocaleString('ru-RU')} / {limit.limit.toLocaleString('ru-RU')} ₽
              </div>
            </div>
          ))}
        </div>
      </div>

      <div className="loyalty-section">
        <h3>Моя лояльность</h3>
        <div className="loyalty-cards">
          <div className="loyalty-card">
            <div className="loyalty-icon">🍽️</div>
            <div className="loyalty-content">
              <span className="loyalty-category">Кэшбек рестораны</span>
              <div className="loyalty-banks">
                <div className="loyalty-bank">
                  <img src={alfaIcon} alt="Альфа-Банк" className="bank-icon-img" />
                  <span>3%</span>
                </div>
              </div>
            </div>
          </div>
          
          <div className="loyalty-card">
            <div className="loyalty-icon">👗</div>
            <div className="loyalty-content">
              <span className="loyalty-category">Кэшбек одежда</span>
              <div className="loyalty-banks">
                <div className="loyalty-bank">
                  <img src={sberIcon} alt="Сбербанк" className="bank-icon-img" />
                  <span>7%</span>
                </div>
                <div className="loyalty-bank">
                  <img src={tbankIcon} alt="Тинькофф" className="bank-icon-img" />
                  <span>2%</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="offers-section">
        <h3>Мои предложения</h3>
        <div className="offers-cards">
          <div className="offer-card">
            <div className="offer-icon">📊</div>
            <div className="offer-content">
              <span className="offer-title">Накопительный счет</span>
              <div className="offer-banks">
                <img src={alfaIcon} alt="Альфа-Банк" className="bank-icon-img" />
                <img src={sberIcon} alt="Сбербанк" className="bank-icon-img" />
              </div>
            </div>
          </div>
          
          <div className="offer-card">
            <div className="offer-icon">🎬</div>
            <div className="offer-content">
              <span className="offer-title">Летняя афиша</span>
              <div className="offer-banks">
                <img src={sberIcon} alt="Сбербанк" className="bank-icon-img" />
              </div>
            </div>
          </div>
          
          <div className="offer-card">
            <div className="offer-icon">✈️</div>
            <div className="offer-content">
              <span className="offer-title">Отпуск с Альфой</span>
              <div className="offer-banks">
                <img src={tbankIcon} alt="Тинькофф" className="bank-icon-img" />
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Модальное окно с историей расходов */}
      {showHistoryModal && (
        <div className="modal-overlay" onClick={() => setShowHistoryModal(false)}>
          <div className="modal-content" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>{historyType}</h3>
              <button className="modal-close" onClick={() => setShowHistoryModal(false)}>×</button>
            </div>
            <div className="modal-body">
              <div className="history-list">
                {historyData.map((transaction, index) => (
                  <div key={index} className="history-item">
                    <div className="history-date">{transaction.date}</div>
                    <div className="history-details">
                      <div className="history-description">{transaction.description}</div>
                      <div className="history-meta">
                        {transaction.category && <span className="history-category">{transaction.category}</span>}
                        {transaction.bank && <span className="history-bank">{transaction.bank}</span>}
                      </div>
                    </div>
                    <div className="history-amount">{transaction.amount.toLocaleString('ru-RU')} ₽</div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ExpensesPage;