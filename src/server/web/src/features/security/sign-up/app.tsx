import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Button, Input, Typography, Sheet, CssVarsProvider, GlobalStyles } from '@mui/joy';
import { loginStart, loginSuccess, loginFailure, logout, selectAuthLoading, selectAuthError, selectIsAuthenticated, selectUser } from './authSlice'; // Import actions and selectors
import { AppDispatch } from '../../../app/store'; // Import AppDispatch, RootState will be used by selectors in authSlice


const LoginPage: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const dispatch: AppDispatch = useDispatch(); // Typed dispatch
    
    const loading = useSelector(selectAuthLoading);
    const error = useSelector(selectAuthError);
    const isAuthenticated = useSelector(selectIsAuthenticated);
    const user = useSelector(selectUser);
  
    const handleLogin = async (event: React.FormEvent) => {
      event.preventDefault();
      dispatch(loginStart());
      // Simulate API Call
      try {
        await new Promise(resolve => setTimeout(resolve, 1000)); // Simulate network delay
        if (username === "admin" && password === "password") {
          // In a real app, the token and user details would come from the server
          dispatch(loginSuccess({ user: { id: '1', username: username }, token: 'fake-jwt-token' }));
        } else {
          dispatch(loginFailure("Invalid username or password. Try admin/password."));
        }
      } catch (e: any) {
        dispatch(loginFailure(e.message || "Failed to login"));
      }
    };
  
    const handleLogout = () => {
      dispatch(logout());
    };
  
    if (isAuthenticated) {
      return (
        <CssVarsProvider>
          <GlobalStyles styles={{ body: { margin: 0, display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '100vh', backgroundColor: 'var(--joy-palette-background-surface)' } }} />
          <Sheet sx={{ width: 300, mx: 'auto', my: 4, py: 3, px: 2, display: 'flex', flexDirection: 'column', gap: 2, borderRadius: 'sm', boxShadow: 'md' }} variant="outlined">
            <Typography level="h4" component="h1"><b>Welcome, {user?.username}!</b></Typography>
            <Typography level="body-sm">You are logged in.</Typography>
            <Button onClick={handleLogout} fullWidth>
              Log out
            </Button>
          </Sheet>
        </CssVarsProvider>
      );
    }
  
    return (
      <CssVarsProvider>
        <GlobalStyles
          styles={{
            body: {
              margin: 0,
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              minHeight: '100vh',
              backgroundColor: 'var(--joy-palette-background-surface)',
            },
          }}
        />
        <Sheet
          sx={{
            width: 300,
            mx: 'auto', // margin left & right
            my: 4, // margin top & bottom
            py: 3, // padding top & bottom
            px: 2, // padding left & right
            display: 'flex',
            flexDirection: 'column',
            gap: 2,
            borderRadius: 'sm',
            boxShadow: 'md',
          }}
          variant="outlined"
        >
          <div>
            <Typography level="h4" component="h1">
              <b>Welcome!</b>
            </Typography>
            <Typography level="body-sm">Sign in to continue.</Typography>
          </div>
          <form onSubmit={handleLogin}>
            <Input
              // htmlAttributes={{ autoComplete: 'username' }} // MUI Joy might use inputProps or similar
              name="username"
              type="text"
              placeholder="Username (admin)"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              sx={{ mb: 1 }}
              required
              disabled={loading === 'pending'}
            />
            <Input
              // htmlAttributes={{ autoComplete: 'current-password' }} // MUI Joy might use inputProps or similar
              name="password"
              type="password"
              placeholder="Password (password)"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              sx={{ mb: 2 }}
              required
              disabled={loading === 'pending'}
            />
            {error && <Typography color="danger" sx={{ mb: 2, fontSize: 'sm' }}>{error}</Typography>}
            <Button type="submit" fullWidth loading={loading === 'pending'}>
              Log in
            </Button>
          </form>
        </Sheet>
      </CssVarsProvider>
    );
  };
  
  export default LoginPage;
  