import type { JSX } from "react"
import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Box, Button, Input, Typography, Sheet, CssVarsProvider, GlobalStyles, Link } from '@mui/joy';
import GoogleIcon from '@mui/icons-material/Google';
import InstagramIcon from '@mui/icons-material/Instagram';
import { loginStart, loginSuccess, loginFailure, logout, selectAuthLoading, selectAuthError, selectIsAuthenticated, selectUser, } from './authSlice'; // Import actions and selectors
import { useLoginMutation } from './authSlice'; 
import { AppDispatch } from '../../../app/store'; // Import AppDispatch, RootState will be used by selectors in authSlice


const LoginPage = (): JSX.Element | null => {
    const [login, { isLoading, error, data }] = useLoginMutation();
    
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [showPassword, setShowPassword] = useState(false); // Add new state variable
    const dispatch: AppDispatch = useDispatch(); // Typed dispatch
    
    const loading = useSelector(selectAuthLoading);
    const errorAuth = useSelector(selectAuthError);
    const isAuthenticated = useSelector(selectIsAuthenticated);
    const user = useSelector(selectUser);
  
    const handleLogin = async (event: React.FormEvent) => {
      event.preventDefault();
      dispatch(loginStart());

      const result = login({ email: username, password: password, language: "en" }).unwrap();
        
      result.then((data) => {
          dispatch(loginSuccess({ user: { id: '1', username: username }, accessToken: data.accessToken, refreshToken: data.refreshToken }));
          setUsername('');
          setPassword('');
      }
      ).catch((error) => {
          dispatch(loginFailure(error.data.message || "Failed to login1"));
      });
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
          <Button fullWidth sx={{ mt: 1 }} variant="outlined" startDecorator={<GoogleIcon />}>
            Login with Google
          </Button>
          <Button fullWidth sx={{ mt: 1 }} variant="outlined" startDecorator={<InstagramIcon />}>
            Login with Instagram
          </Button>
          <Typography textAlign="center" sx={{ my: 1, mb: 1 }}>
            Or with email and password
          </Typography>
          <form onSubmit={handleLogin}>
            <Input
              // htmlAttributes={{ autoComplete: 'username' }} // MUI Joy might use inputProps or similar
              name="username"
              type="text"
              placeholder="Username (admin)"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              onBlur={() => { if (username) setShowPassword(true); }} // Show password field on blur if username is not empty
              sx={{ mb: 1 }}
              required
              disabled={loading === 'pending'}
            />
            {showPassword && ( // Conditionally render password field
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
            )}
            {errorAuth && <Typography color="danger" sx={{ mb: 2, fontSize: 'sm' }}>{errorAuth}</Typography>}
            <Button type="submit" fullWidth loading={isLoading}>
              Log in
            </Button>
            <Box sx={{ display: 'flex', justifyContent: 'flex-end', width: '100%', mt: 1 }}>
              <Button component={Link} href={`/forgot?username=${username}`} variant="plain">
                Forgot Password?
              </Button>
            </Box>
            {error && <p>Error: {JSON.stringify(error)}</p>}
            {data && <p>Login successful! Token: {data.accessToken}</p>} 
            {data && <p>Login successful! Token: {data.refreshToken}</p>} 
          </form>
        </Sheet>
      </CssVarsProvider>
    );
  };
  
  export default LoginPage;
