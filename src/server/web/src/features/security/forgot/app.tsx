import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { Button, Input, Typography, Sheet, CssVarsProvider, GlobalStyles} from '@mui/joy';
import { useLocation } from 'react-router-dom'; // Added import
//import { useLoginMutation } from './forgotSlice'; 
// Placeholder for RootState type, adjust as per your actual store configuration
interface RootState {
  auth: {
    isAuthenticated: boolean;
    user: any; // Replace 'any' with your user type
    // Add other auth state properties here
  };
  // Add other slices of your state here
}

const Fortgo: React.FC = () => {
  const [usernameOrEmail, setUsernameOrEmail] = useState('');
  const location = useLocation(); // Added to get location object

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);
    const usernameFromQuery = queryParams.get('username');
    if (usernameFromQuery) {
      setUsernameOrEmail(usernameFromQuery);
    }
  }, [location.search]); // Added useEffect to read query param
  
  // Example: Selecting data from the auth slice
  // Replace with actual selectors from your authSlice.ts
  const isAuthenticated = useSelector((state: RootState) => state.auth.isAuthenticated);
  const user = useSelector((state: RootState) => state.auth.user);

  // Example: Dispatching an action
   const handleLogin = () => {
     //dispatch(loginAction({ username: 'testuser', password: 'password' }));
   };

   const handleLogout = () => {
     //dispatch(logoutAction());
   };

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
      <Typography level="h4" component="h2" textAlign="center">
        Fortgo Component
      </Typography>
      <Input
        placeholder="Enter Username or Email"
        value={usernameOrEmail}
        onChange={(e) => setUsernameOrEmail(e.target.value)}
      />
      {isAuthenticated ? (
        <Sheet sx={{ display: 'flex', flexDirection: 'column', gap: 1, alignItems: 'center' }}>
          <Typography level="body-md">Welcome, {user?.name || 'User'}!</Typography>
            <Button onClick={handleLogout} variant="soft">Logout</Button>
        </Sheet>
      ) : (
        <Sheet sx={{ display: 'flex', flexDirection: 'column', gap: 1, alignItems: 'center' }}>
          <Typography level="body-md">You are not logged in.</Typography>
          <Button onClick={handleLogin} variant="solid">Recover Password</Button>
        </Sheet>
      )}
      {/* Add more UI elements for password recovery or other security features */}
    </Sheet>
    </CssVarsProvider>
    
  );
};

export default Fortgo;
