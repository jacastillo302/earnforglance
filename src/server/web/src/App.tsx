import './css/App.css'
import { Provider } from 'react-redux';
import { Layaut } from "./features/layout/layout"
import { store } from './app/store';
import LoginPage from './features/security/sign-up/app';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

export const App = () => (
  <Provider store={store}>  
  <Layaut>        
      <Router>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/" element={<LoginPage />} /> // Default to login page for now
        </Routes>
      </Router>       
      {/* <LoginPage /> */}
  </Layaut>
  </Provider>
)
