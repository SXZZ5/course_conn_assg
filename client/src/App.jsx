import './App.css'
import { useNavigate } from 'react-router';

function App() {
  const navigate = useNavigate();

  const handleSignup = () => {
    navigate('/signup');
  };

  const handleLogin = () => {
    navigate('/login');
  };

  const handleOldCookies = () => {
    navigate('/oldcookies');
  };

  return (
    <div>
      <button onClick={handleSignup}>Sign Up</button>
      <button onClick={handleLogin}>Login</button>
      <button onClick={handleOldCookies}>Old Cookies</button>
    </div>
  );

}

export default App
