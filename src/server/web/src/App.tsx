import './css/App.css'
import { Provider } from 'react-redux';
import { Layaut } from "./features/layout/app"
import { store } from './app/store';
import LoginPage from './features/security/sign-up/app';
import ForgotPage from './features/security/forgot/app';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

export const App = () => (
  <Provider store={store}>  
  <Layaut>        
      <Router>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/forgot" element={<ForgotPage />} />
          <Route path="/" element={<LoginPage />} /> // Default to login page for now
        </Routes>
      </Router>       
  </Layaut>
  </Provider>
)
