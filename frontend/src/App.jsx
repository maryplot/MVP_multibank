import { useState } from 'react';
import Login from './components/Login';
import Register from './components/Register';
import Dashboard from './components/Dashboard';
import ExpensesPage from './components/ExpensesPage';
import BottomNavigation from './components/BottomNavigation';
import './App.css'

function App() {
  const [token, setToken] = useState(localStorage.getItem('token'));
  const [showRegister, setShowRegister] = useState(false);
  const [activeTab, setActiveTab] = useState('dashboard');

  console.log('App rendered with token:', token);
  console.log('localStorage token:', localStorage.getItem('token'));

  const handleLogin = (newToken) => {
    console.log('handleLogin called with token:', newToken);
    setToken(newToken);
    setShowRegister(false);
  };

  const handleRegister = (userData) => {
    alert(`Пользователь ${userData.username} создан! Теперь войдите в систему.`);
    setShowRegister(false);
  };

  const handleLogout = () => {
    console.log('handleLogout called');
    localStorage.removeItem('token');
    localStorage.removeItem('user_id');
    setToken(null);
  };

  const handleTabChange = (tab) => {
    console.log('handleTabChange called with tab:', tab);
    setActiveTab(tab);
  };

  if (!token) {
    console.log('Rendering login/register screen');
    return showRegister ? (
      <Register
        onRegister={handleRegister}
        onSwitchToLogin={() => setShowRegister(false)}
      />
    ) : (
      <Login
        onLogin={handleLogin}
        onSwitchToRegister={() => setShowRegister(true)}
      />
    );
  }

  console.log('Rendering dashboard/expenses screen with activeTab:', activeTab);

  return (
    <div className="app">
      <div className="app-content">
        {activeTab === 'dashboard' && (
          <Dashboard onLogout={handleLogout} />
        )}
        {activeTab === 'expenses' && (
          <ExpensesPage onBack={() => setActiveTab('dashboard')} />
        )}
      </div>
      <BottomNavigation
        activeTab={activeTab}
        onTabChange={handleTabChange}
      />
    </div>
  );
}

export default App;