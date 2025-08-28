import { useState } from 'react';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import './App.css'

function App() {
  const [token, setToken] = useState(localStorage.getItem('token'));

  const handleLogin = (newToken) => {
    setToken(newToken);
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user_id');
    setToken(null);
  };

  if (!token) {
    return <Login onLogin={handleLogin} />;
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
        <div>
          <span style={{ marginRight: '15px' }}>
            User ID: {localStorage.getItem('user_id')}
          </span>
          <button onClick={handleLogout}>Выйти</button>
        </div>
      </header>
      
      <main>
        <Dashboard />
      </main>
    </div>
  );
}

export default App;