import React from 'react';
import './BottomNavigation.css';

const BottomNavigation = ({ activeTab, onTabChange }) => {
  return (
    <div className="bottom-navigation">
      <button 
        className={`nav-item ${activeTab === 'dashboard' ? 'active' : ''}`}
        onClick={() => onTabChange('dashboard')}
      >
        <div className="nav-icon">🏠</div>
        <span className="nav-label">Главная</span>
      </button>
      
      <button 
        className={`nav-item ${activeTab === 'expenses' ? 'active' : ''}`}
        onClick={() => onTabChange('expenses')}
      >
        <div className="nav-icon">📊</div>
        <span className="nav-label">Анализ</span>
      </button>
    </div>
  );
};

export default BottomNavigation;