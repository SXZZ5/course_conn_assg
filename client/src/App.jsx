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

  const handleLogout = async () => {
    const response = await fetch("http://localhost:8080/logout", {
        method: 'GET',
        credentials: 'include'
    })
    if(response.ok) {
        alert("Logged Out succesfully");
    } else {
        alert("Could Not Log Out");
    }
  }

  return (
    <div>
      <button onClick={handleSignup}>Sign Up</button>
      <button onClick={handleLogin}>Login</button>
      <button onClick={handleOldCookies}>Old Cookies</button>
      <button onClick={handleLogout}>Log Out.</button>
    </div>
  );

}

export default App
