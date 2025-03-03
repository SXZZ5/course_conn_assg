
import { createRoot } from 'react-dom/client'
import './index.css'
import { BrowserRouter, Routes, Route } from "react-router";
import App from './App';
import Login from './login';
import Signup from './signup';
import Results from './results';
createRoot(document.getElementById('root')).render(
    <BrowserRouter>
        <Routes>
            <Route path="/oldcookies" element={<Results />} />
            <Route path="/login" element={<Login/>}/>
            <Route path="/signup" element={<Signup/>}/>
            <Route path="/" element={<App />}/>
        </Routes>
    </BrowserRouter>
)
