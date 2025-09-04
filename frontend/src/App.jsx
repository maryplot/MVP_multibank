import { useState } from 'react';
import Login from './components/Login';
import Register from './components/Register';
import Dashboard from './components/Dashboard';
import './App.css'

function App() {
  const [token, setToken] = useState(localStorage.getItem('token'));
  const [showRegister, setShowRegister] = useState(false);

  const handleLogin = (newToken) => {
    setToken(newToken);
    setShowRegister(false);
  };

  const handleRegister = (userData) => {
    alert(`Пользователь ${userData.username} создан! Теперь войдите в систему.`);
    setShowRegister(false);
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user_id');
    setToken(null);
  };

  if (!token) {
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

  return (
    <div>
      <header style={{ 
        padding: '20px', 
        background: '#f5f5f5', 
        display: 'flex', 
        justifyContent: 'space-between',
        alignItems: 'center'
      }}>
        <h1>Мультибанк</h1>
        <button onClick={handleLogout}>Выйти</button>
      </header>
      
      <main>
        <Dashboard />
      </main>
    </div>
  );
}

export default App;