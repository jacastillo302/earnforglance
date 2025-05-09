import './css/App.css'
import { Provider } from 'react-redux';
import { Layaut } from "./features/layout/app"
import { store } from './app/store';
import HomePage from './features/home/app';
import LoginPage from './features/security/sign-up/app';
import ForgotPage from './features/security/forgot/app';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

export const App = () => (
  <Provider store={store}>  
  <Layaut>        
      <Router>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/forgot" element={<ForgotPage />} />          
        </Routes>
      </Router>       
  </Layaut>
  </Provider>
)
